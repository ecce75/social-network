package handler

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"backend/util"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PostHandler struct {
	postRepo *repository.PostRepository
	sessionRepo *repository.SessionRepository
	friendsRepo *repository.FriendsRepository
	groupMemberRepo *repository.GroupMemberRepository
}

func NewPostHandler(postRepo *repository.PostRepository, sessionRepo *repository.SessionRepository, friendsRepo *repository.FriendsRepository, groupMemberRepo *repository.GroupMemberRepository) *PostHandler {
	return &PostHandler{postRepo: postRepo, sessionRepo: sessionRepo, friendsRepo: friendsRepo, groupMemberRepo: groupMemberRepo}
}

func (h *PostHandler) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var request model.CreatePostRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to decode request data", http.StatusBadRequest)
		return
	}

	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error confirming authentication: " + err.Error(), http.StatusInternalServerError)
		return
	}

	// Creates the post in database
	postID, err := h.postRepo.CreatePost(request, userID)
	if err != nil {
		http.Error(w, "Failed to create the post: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Successful response
	response := map[string]interface{}{
		"message": "Post created successfully",
		"data": postID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *PostHandler) EditPostHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the request body for updating the post
	var request model.UpdatePostRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to decode request data", http.StatusBadRequest)
		return
	}

	// Confirm user auth and get userid
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error confirming user authentication: " + err.Error(), http.StatusUnauthorized)
		return
	}

	// Update the post in the database
	err = h.postRepo.UpdatePost(request.Id, userID, request)
	if err != nil {
		http.Error(w, "Failed to update the post: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Successful response
	response := map[string]string{
		"message": "Post updated successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


func (h *PostHandler) DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the post ID from the URL
	vars := mux.Vars(r)
	postID, ok := vars["id"]
	intpostID, err := strconv.Atoi(postID)
	if err != nil {
		http.Error(w, "Failed to parse post ID: " +err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "Post ID is missing in parameters", http.StatusBadRequest)
		return
	}

	userId, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Delete the post from the database
	err = h.postRepo.DeletePost(intpostID, userId)
	if err != nil {
		http.Error(w, "Failed to delete the post: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Successful response
	response := map[string]string{
		"message": "Post deleted successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


func (h *PostHandler) GetAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error confirming user authentication: " + err.Error(), http.StatusUnauthorized)
		return
	}

	userGroupsPosts, err := h.postRepo.GetPostsByUserGroups(userID)
	if err != nil {
		http.Error(w, "Failed to retrieve posts by user groups: " + err.Error(), http.StatusInternalServerError)
		return
	}
	posts, err := h.postRepo.GetAllPostsWithUserIDAccess(userID)
	if err != nil {
		http.Error(w, "Failed to retrieve posts: " + err.Error(), http.StatusInternalServerError)
		return
	}
	posts = append(posts, userGroupsPosts...)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// GetAllUserPosts retrieves all posts for a specific user.
// It takes the user ID from the request parameters and checks the user's authentication.
// If the requesting user is the same as the user ID in the parameters, it retrieves all posts for that user.
// If the requesting user is not the same, it checks the friend status between the requesting user and the user ID in the parameters.
// If they are friends, it retrieves all posts for that user.
// If they are not friends, it retrieves only the public posts for that user.
// The retrieved posts are encoded as JSON and sent in the response.
func (h *PostHandler) GetAllUserPostsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, ok := vars["id"]
	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Failed to parse user ID: " +err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "User ID is missing in parameters", http.StatusBadRequest)
		return
	}
	requestingUserID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error confirming user authentication: " + err.Error(), http.StatusUnauthorized)
		return
	}
	var posts []model.Post
	if requestingUserID == intUserID {
		posts, err = h.postRepo.GetAllUserPosts(requestingUserID)
		if err != nil {
			http.Error(w, "Failed to retrieve posts: " + err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// check if the users are friends
		status, err := h.friendsRepo.GetFriendStatus(requestingUserID, intUserID)
		if err != nil {
			http.Error(w, "Failed to retrieve friend status: " + err.Error(), http.StatusInternalServerError)
			return
		}
		// retrieve users public posts, and private posts if they are friends
		if status == "accepted" {
			posts, err = h.postRepo.GetAllUserPosts(intUserID)
			if err != nil {
				http.Error(w, "Failed to retrieve posts: " + err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			posts, err = h.postRepo.GetAllUserPublicPosts(intUserID)
			if err != nil {
				http.Error(w, "Failed to retrieve posts: " + err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// ---------------------------------------------- //
// ------------ Group Posts Handlers ------------ //
// ---------------------------------------------- //

// TODO: check if user requesting is in group
func (h *PostHandler) GetPostsByGroupIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID, ok := vars["id"]
	intGroupID, err := strconv.Atoi(groupID)
	if err != nil {
		http.Error(w, "Failed to parse group ID: " +err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "Group ID is missing in parameters", http.StatusBadRequest)
		return
	}
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error confirming user authentication: " + err.Error(), http.StatusUnauthorized)
		return
	}
	// check if user is in group
	isMember, err := h.groupMemberRepo.IsUserGroupMember(userID, intGroupID)
	if err != nil {
		http.Error(w, "Failed to check if user is in group: " + err.Error(), http.StatusInternalServerError)
		return
	}
	if !isMember {
		http.Error(w, "User is not a member of the group", http.StatusUnauthorized)
		return
	}

	posts, err := h.postRepo.GetPostsByGroupID(intGroupID)
	if err != nil {
		http.Error(w, "Failed to retrieve posts: " + err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}