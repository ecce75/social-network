package handler

import (
	"backend/pkg/repository"
	"backend/util"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type FriendHandler struct {
	userRepository      *repository.UserRepository
	friendRepository    *repository.FriendsRepository
	sessionRepository   *repository.SessionRepository
	notificationHandler *NotificationHandler
}

func NewFriendHandler(friendRepository *repository.FriendsRepository, sessionRepository *repository.SessionRepository, notificationHandler *NotificationHandler, userRepository *repository.UserRepository) *FriendHandler {
	return &FriendHandler{friendRepository: friendRepository,
		sessionRepository:   sessionRepository,
		notificationHandler: notificationHandler,
		userRepository:      userRepository,
	}
}

// GetFriendRequests retrieves the friend requests for the authenticated user.
// It requires the user to be authenticated and returns the friend requests in JSON format.
// If there is an error retrieving the friend requests, it returns an HTTP 500 Internal Server Error.
func (h *FriendHandler) GetFriendRequestsHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := h.sessionRepository.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}
	requests, err := h.friendRepository.GetFriendRequests(userID)
	if err != nil {
		http.Error(w, "Error getting friend requests: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(requests)
}

// SendFriendRequestHandler handles the HTTP request for sending a friend request.
// It checks if the user is authenticated, validates the friend ID, and checks the friend status.
// If no friend request exists, it sends a friend request to the specified user.
// If a friend request is already pending or the users are already friends or one user has blocked the other, it returns an error.
// It returns http.StatusCreated if the friend request is sent successfully.
func (h *FriendHandler) SendFriendRequestHandler(w http.ResponseWriter, r *http.Request) {
	sessionToken := util.GetSessionToken(r)
	userID, err := h.sessionRepository.GetUserIDFromSessionToken(sessionToken)
	if err != nil {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	friendID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid friend ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	status, err := h.friendRepository.GetFriendStatus(userID, friendID)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "Error checking friend status: "+err.Error(), http.StatusInternalServerError)
		return
	}

	switch status {
	case "":
		err = h.friendRepository.AddFriend(userID, friendID)
		if err != nil {
			http.Error(w, "Error sending friend request "+err.Error(), http.StatusInternalServerError)
			return
		} // No friend request exists, proceed to send one
	case "declined":
		err = h.friendRepository.UpdateFriendStatus(userID, friendID, "pending")
		if err != nil {
			http.Error(w, "Error sending friend request "+err.Error(), http.StatusInternalServerError)
			return
		}
		// The user has declined a friend request from the other user, proceed to send a new friend request
	case "pending":
		http.Error(w, "A friend request is already pending between these users", http.StatusConflict)
		return
	case "accepted":
		http.Error(w, "These users are already friends", http.StatusConflict)
		return
	case "blocked":
		http.Error(w, "One of these users has blocked the other", http.StatusConflict)
		return
	default:
		http.Error(w, "Unknown friend status: "+status, http.StatusInternalServerError)
		return
	}

	// get user for notification message
	user, err := h.userRepository.GetUserProfileByID(userID)
	if err != nil {
		http.Error(w, "Error getting user data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	message := user.FirstName + " " + user.LastName + " sent you a friend request"
	err = h.notificationHandler.CreateNotification(friendID, userID, "friend", message)
	if err != nil {
		http.Error(w, "Error sending notification "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Friend request sent", userID, friendID, status, err)
	w.WriteHeader(http.StatusCreated)
}

// AcceptFriendRequestHandler handles the HTTP request for accepting a friend request.
// It requires the user to be authenticated and the friend ID to be valid.
// If successful, it updates the friend status to "accepted" and returns a 200 OK response.
// If any error occurs, it returns an appropriate HTTP error response.
func (h *FriendHandler) AcceptFriendRequestHandler(w http.ResponseWriter, r *http.Request) {
	sessionToken := util.GetSessionToken(r)
	userID, err := h.sessionRepository.GetUserIDFromSessionToken(sessionToken)
	if err != nil {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	friendRequestID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid friend ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = h.friendRepository.UpdateFriendStatus(userID, friendRequestID, "accepted")
	if err != nil {
		http.Error(w, "Error accepting friend request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	addedFriend, err := h.userRepository.GetUsernameByID(friendRequestID)
	if err != nil {
		http.Error(w, "Error getting friend data: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// Notification
	message := "You are now friends with " + addedFriend
	err = h.notificationHandler.EditFriendRequestNotification(userID, friendRequestID, message)
	if err != nil {
		http.Error(w, "Error sending notification "+err.Error(), http.StatusInternalServerError)
		return
	}
	username, err := h.userRepository.GetUsernameByID(userID)
	message = username + " accepted your friend request"
	err = h.notificationHandler.CreateNotification(friendRequestID, userID, "friend", message)
	if err != nil {
		http.Error(w, "Error changing notification message"+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeclineFriendRequestHandler handles the HTTP request for declining a friend request.
// It requires the user to be authenticated and the friend ID to be valid.
// If successful, it updates the friend status to "declined" and returns a 200 OK response.
// If there is an error, it returns an appropriate HTTP error response.
func (h *FriendHandler) DeclineFriendRequestHandler(w http.ResponseWriter, r *http.Request) {
	sessionToken := util.GetSessionToken(r)
	userID, err := h.sessionRepository.GetUserIDFromSessionToken(sessionToken)
	if err != nil {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	friendRequestID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid friend ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = h.friendRepository.UpdateFriendStatus(userID, friendRequestID, "declined")
	if err != nil {
		http.Error(w, "Error declining friend request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	username, err := h.userRepository.GetUsernameByID(friendRequestID)

	message := "You declined the friend request from " + username
	err = h.notificationHandler.EditFriendRequestNotification(userID, friendRequestID, message)
	if err != nil {
		http.Error(w, "Error changing notification message"+err.Error(), http.StatusInternalServerError)
		return
	}

	username, err = h.userRepository.GetUsernameByID(userID)
	message = username + " declined your friend request"
	err = h.notificationHandler.CreateNotification(friendRequestID, userID, "friend", message)
	if err != nil {
		http.Error(w, "Error changing notification message"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetFriendsHandler handles the HTTP request for retrieving the friends of a user.
// It requires a valid session token in the request header for authentication.
// If the user is not authenticated, it returns a 401 Unauthorized error.
// If there is an error retrieving the friends, it returns a 500 Internal Server Error.
// The response is encoded in JSON format.
func (h *FriendHandler) GetFriendsHandler(w http.ResponseWriter, r *http.Request) {

	userID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		sessionToken := util.GetSessionToken(r)
		userID, err = h.sessionRepository.GetUserIDFromSessionToken(sessionToken)
		if err != nil {
			http.Error(w, "User not authenticated", http.StatusUnauthorized)
			return
		}
	}

	friends, err := h.friendRepository.GetFriends(userID)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "Error getting friends: "+err.Error(), http.StatusInternalServerError)
		return
	} else if err == sql.ErrNoRows {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "No friends found"})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(friends)
}

func (h *FriendHandler) CheckFriendStatusHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := h.sessionRepository.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}
	friendID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid friend ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	status, err := h.friendRepository.GetFriendStatus(userID, friendID)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "Error checking friend status: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}
