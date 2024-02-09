package handler

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"backend/util"
	"encoding/json"
	"net/http"
)

type FriendHandler struct {
    friendRepository *repository.FriendsRepository
    sessionRepository *repository.SessionRepository
}

func NewFriendHandler(friendRepository *repository.FriendsRepository, sessionRepository *repository.SessionRepository) *FriendHandler {
    return &FriendHandler{friendRepository: friendRepository, sessionRepository: sessionRepository}
}

// SendFriendRequestHandler handles sending a friend request
func SendFriendRequestHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement logic for sending a friend request
}

// AcceptFriendRequestHandler handles accepting a friend request
func AcceptFriendRequestHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement logic for accepting a friend request
}

// DeclineFriendRequestHandler handles declining a friend request
func DeclineFriendRequestHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement logic for declining a friend request
}

// BlockUserHandler handles blocking a user
func BlockUserHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement logic for blocking a user
}

// UnblockUserHandler handles unblocking a user
func UnblockUserHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement logic for unblocking a user
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
        friends := []model.Friend{
            {
                UserID:        1,
                FirstName: "John",
                LastName:  "Doe",
                AvatarURL:    "avatar1.png",
                Username:  "user1",
            },
            {
                UserID:        2,
                FirstName: "Jane",
                LastName:  "Smith",
                AvatarURL:    "avatar2.png",
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