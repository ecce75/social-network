package handler

import (
	"backend/pkg/db/sqlite"
	"backend/pkg/model"
	"backend/pkg/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: id may not come from request and will cause error
	var newComment model.Comment 
    err := json.NewDecoder(r.Body).Decode(&newComment)
    if err != nil {
        http.Error(w, "Error decoding request body: "+err.Error(), http.StatusBadRequest)
        return
    }

    // Authentication and authorization logic goes here
    // Assuming you have a function to check the user's session and get the user ID
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the session cookie doesn't exist, set isAuthenticated to false
			http.Error(w, "User not authenticated", http.StatusUnauthorized)
			return
		} else {
			http.Error(w, "Error checking session token: " + err.Error(), http.StatusInternalServerError)
			return
		}
	}
    userID, err := confirmAuthentication(cookie)
    if err != nil {
        http.Error(w, "User not authenticated: "+err.Error(), http.StatusUnauthorized)
        return
    }
    newComment.UserID = userID

    // Insert the comment into the database
    createdCommentId, err := repository.CreateComment(sqlite.Dbase, newComment)
    if err != nil {
        http.Error(w, "Failed to create comment: "+err.Error(), http.StatusInternalServerError)
        return
    }

	// Successful response
	response := map[string]interface{}{
		"message": "Comment created successfully",
		"data": createdCommentId,
	}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func GetCommentByUserIDorPostID(w http.ResponseWriter, r *http.Request) {
	var id string

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&id)
	if err != nil {
		http.Error(w, "Error decoding id for comment request: "+err.Error(), http.StatusBadRequest)
		return
	}
	intid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error converting id to int: "+err.Error(), http.StatusInternalServerError)
		return
	}
	comments, err := repository.GetCommentsByID(sqlite.Dbase, intid)
	if err != nil {
		http.Error(w, "Error retrieving comments: "+ err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}

func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the post ID from the URL
	vars := mux.Vars(r)
	commentID, ok := vars["id"]
	intcommentID, err := strconv.Atoi(commentID)
	if err != nil {
		http.Error(w, "Failed to parse comment ID: " +err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "Comment ID is missing in parameters", http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "Error authenticating user: " +err.Error(), http.StatusUnauthorized)
		return
	}

	// Confirm user auth and get userid
	userID, err := confirmAuthentication(cookie)
	if err != nil {
		http.Error(w, "Error confirming user authentication: " + err.Error(), http.StatusUnauthorized)
		return
	}

	// Delete the comment from the database
	err = repository.DeleteComment(sqlite.Dbase, intcommentID, userID)
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

func EditCommentHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the post ID from the URL
	vars := mux.Vars(r)
	commentID, ok := vars["id"]
	intcommentID, err := strconv.Atoi(commentID)
	if err != nil {
		http.Error(w, "Failed to parse comment ID: " +err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "Comment ID is missing in parameters", http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "Error authenticating user: " +err.Error(), http.StatusUnauthorized)
		return
	}

	// Confirm user auth and get userid
	userID, err := confirmAuthentication(cookie)
	if err != nil {
		http.Error(w, "Error confirming user authentication: " + err.Error(), http.StatusUnauthorized)
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
	err = repository.UpdateComment(sqlite.Dbase, intcommentID, userID, commentData)
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