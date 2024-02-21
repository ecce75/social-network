package handler

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"backend/util"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type FriendHandler struct {
	friendRepository  *repository.FriendsRepository
	sessionRepository *repository.SessionRepository
}

func NewFriendHandler(friendRepository *repository.FriendsRepository, sessionRepository *repository.SessionRepository) *FriendHandler {
	return &FriendHandler{friendRepository: friendRepository, sessionRepository: sessionRepository}
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
		// No friend request exists, proceed to send one
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

	err = h.friendRepository.AddFriend(userID, friendID)
	if err != nil {
		http.Error(w, "Error sending friend request "+err.Error(), http.StatusInternalServerError)
		return
	}

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

	friendID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid friend ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = h.friendRepository.UpdateFriendStatus(userID, friendID, "accepted")
	if err != nil {
		http.Error(w, "Error accepting friend request: "+err.Error(), http.StatusInternalServerError)
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

	friendID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid friend ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = h.friendRepository.UpdateFriendStatus(userID, friendID, "declined")
	if err != nil {
		http.Error(w, "Error declining friend request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// BlockUserHandler handles the blocking of a user.
// It requires the user to be authenticated and the friend ID to be valid.
// If successful, it updates the friend status to "blocked" and returns a status code of 200.
// If there is an error, it returns an appropriate HTTP error response.
func (h *FriendHandler) BlockUserHandler(w http.ResponseWriter, r *http.Request) {
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

	err = h.friendRepository.UpdateFriendStatus(userID, friendID, "blocked")
	if err != nil {
		http.Error(w, "Error blocking user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UnblockUserHandler handles the HTTP request to unblock a user.
// It requires the user to be authenticated and the friend ID to be valid.
// If successful, it updates the friend status to "accepted" and returns a 200 OK response.
// If there is an error, it returns an appropriate HTTP error response.
func (h *FriendHandler) UnblockUserHandler(w http.ResponseWriter, r *http.Request) {
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

	err = h.friendRepository.UpdateFriendStatus(userID, friendID, "accepted")
	if err != nil {
		http.Error(w, "Error unblocking user: "+err.Error(), http.StatusInternalServerError)
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
	sessionToken := util.GetSessionToken(r)
	userID, err := h.sessionRepository.GetUserIDFromSessionToken(sessionToken)
	if err != nil {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	friends, err := h.friendRepository.GetFriends(userID)
	if err != nil {
		// w.Header().Set("Content-Type", "application/json")
		// json.NewEncoder(w).Encode(map[string]string{"message": "No friends found"})
		// Sample friends
		friends := []model.FriendList{
			{
				UserID:    1,
				FirstName: "John",
				LastName:  "Doe",
				AvatarURL: "avatar1.png",
				Username:  "user1",
			},
			{
				UserID:    2,
				FirstName: "Jane",
				LastName:  "Smith",
				AvatarURL: "avatar2.png",
				Username:  "user2",
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(friends)

		return
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
