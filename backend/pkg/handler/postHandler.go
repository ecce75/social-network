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
}

func NewPostHandler(postRepo *repository.PostRepository, sessionRepo *repository.SessionRepository) *PostHandler {
	return &PostHandler{postRepo: postRepo, sessionRepo: sessionRepo}
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

// ---------------------------------------------- //
// ------------ Group Posts Handlers ------------ //
// ---------------------------------------------- //


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

	posts, err := h.postRepo.GetPostsByGroupID(intGroupID)
	if err != nil {
		http.Error(w, "Failed to retrieve posts: " + err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}