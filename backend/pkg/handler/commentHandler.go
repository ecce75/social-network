package handler

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"backend/util"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type CommentHandler struct {
	commentRepo         *repository.CommentRepository
	sessionRepo         *repository.SessionRepository
	notificationHandler *NotificationHandler
	postRepo            *repository.PostRepository
	userRepo            *repository.UserRepository
	VoteHandler         *VoteHandler
}

func NewCommentHandler(commentRepo *repository.CommentRepository, sessionRepo *repository.SessionRepository, notificationHandler *NotificationHandler, postRepo *repository.PostRepository, userRepo *repository.UserRepository, voteHandler *VoteHandler) *CommentHandler {
	return &CommentHandler{commentRepo: commentRepo, sessionRepo: sessionRepo, notificationHandler: notificationHandler, postRepo: postRepo, userRepo: userRepo, VoteHandler: voteHandler}
}

func (h *CommentHandler) CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	err1 := r.ParseMultipartForm(10 << 20) // Maximum memory 10MB, change this based on your requirements
	if err1 != nil {
		http.Error(w, "Error parsing form data: "+err1.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	postId, ok := vars["id"]
	intPostId, err := strconv.Atoi(postId)
	if !ok {
		http.Error(w, "Post ID is missing in parameters", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "Error decoding id for comment request: "+err.Error(), http.StatusBadRequest)
		return
	}
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "User not authenticated: "+err.Error(), http.StatusUnauthorized)
		return
	}
	var newComment model.Comment
	newComment.Content = r.FormValue("content")
	newComment.PostID = intPostId
	newComment.UserID = userID

	_, _, err = r.FormFile("image")
	if err != nil {
		newComment.Image.String = ""
	} else {
		newComment.Image.String = os.Getenv("NEXT_PUBLIC_URL") + ":" + os.Getenv("NEXT_PUBLIC_BACKEND_PORT") + "/images/comments/brt"
	}

	// Insert the comment into the database
	commentID, err := h.commentRepo.CreateComment(&newComment)
	if err != nil {
		http.Error(w, "Failed to create comment: "+err.Error(), http.StatusInternalServerError)
		return
	}
	util.ImageSave(w, r, strconv.Itoa(int(commentID)), "comment")

	// NOTIFICATION
	postOwnerId, err := h.postRepo.GetPostOwnerIDByPostID(newComment.PostID)
	if err != nil {
		http.Error(w, "Failed to get post owner id: "+err.Error(), http.StatusInternalServerError)
		return
	}
	post, err := h.postRepo.GetPostByID(newComment.PostID)
	if err != nil {
		http.Error(w, "Failed to get post: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// authenticated user username
	username, err := h.userRepo.GetUsernameByID(userID)
	if err != nil {
		http.Error(w, "Failed to get username: "+err.Error(), http.StatusInternalServerError)
		return
	}
	message := username + " commented on your post: " + post.Title

	if newComment.UserID != int(postOwnerId) {
		err = h.notificationHandler.CreateNotification(int(postOwnerId), userID, "post", message)
		if err != nil {
			http.Error(w, "Failed to create notification: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	user, err := h.userRepo.GetUserProfileByID(newComment.UserID)
	if err != nil {
		http.Error(w, "Error getting user profile: "+err.Error(), http.StatusInternalServerError)
		return
	}
	commentImageUrl, err := h.commentRepo.GetCommentImageURL(int(commentID))
	if err != nil {
		http.Error(w, "Error getting comment image url: "+err.Error(), http.StatusInternalServerError)
		return

	}
	commentsResponse := model.CommentsResponse{
		Id:        int(commentID),
		PostID:    newComment.PostID,
		UserID:    newComment.UserID,
		Content:   newComment.Content,
		Image:     commentImageUrl,
		CreatedAt: time.Now(),
		Likes:     0,
		Dislikes:  0,
		Username:  username,
		ImageURL:  user.AvatarURL, // Set the avatar URL here
	}

	// Successful response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(commentsResponse)
}

func (h *CommentHandler) GetCommentsByPostID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId, ok := vars["id"]
	intPostId, err := strconv.Atoi(postId)

	if !ok {
		http.Error(w, "Post ID is missing in parameters", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "Error decoding id for comment request: "+err.Error(), http.StatusBadRequest)
		return
	}

	comments, err := h.commentRepo.GetAllPostComments(intPostId)
	if err != nil {
		http.Error(w, "Error retrieving comments: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// get votes for each comment and append to CommentResponse
	commentsWithVotes, err := h.VoteHandler.AppendVotesToCommentsResponse(comments)
	if err != nil {
		http.Error(w, "Error appending votes to comments: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// we need user information per comment aswell - username, profile picture
	for i, comment := range commentsWithVotes {
		user, err := h.userRepo.GetUserProfileByID(comment.UserID)
		if err != nil {
			http.Error(w, "Error getting user profile: "+err.Error(), http.StatusInternalServerError)
			return
		}
		commentsWithVotes[i].Username = user.Username
		commentsWithVotes[i].ImageURL = user.AvatarURL
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(commentsWithVotes)
}

func (h *CommentHandler) DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the post ID from the URL
	vars := mux.Vars(r)
	commentID, ok := vars["id"]
	intcommentID, err := strconv.Atoi(commentID)
	if err != nil {
		http.Error(w, "Failed to parse comment ID: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "Comment ID is missing in parameters", http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "Error authenticating user: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Confirm user auth and get userid
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(cookie.Value)
	if err != nil {
		http.Error(w, "Error confirming user authentication: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Delete the comment from the database
	err = h.commentRepo.DeleteComment(intcommentID, userID)
	if err != nil {
		http.Error(w, "Failed to delete the comment: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Successful response
	response := map[string]string{
		"message": "Comment deleted successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *CommentHandler) EditCommentHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the post ID from the URL
	vars := mux.Vars(r)
	commentID, ok := vars["id"]
	intcommentID, err := strconv.Atoi(commentID)
	if err != nil {
		http.Error(w, "Failed to parse comment ID: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "Comment ID is missing in parameters", http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "Error authenticating user: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Confirm user auth and get userid
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(cookie.Value)
	if err != nil {
		http.Error(w, "Error confirming user authentication: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Parse the comment data from the request body
	var commentData model.UpdateCommentRequest
	err = json.NewDecoder(r.Body).Decode(&commentData)
	if err != nil {
		http.Error(w, "Error decoding request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Update the comment in the database
	err = h.commentRepo.UpdateComment(intcommentID, userID, commentData)
	if err != nil {
		http.Error(w, "Failed to update the comment: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Successful response
	response := map[string]string{
		"message": "Comment updated successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
