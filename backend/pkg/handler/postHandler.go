package handler

import (
	"backend/pkg/db/sqlite"
	"backend/pkg/model"
	"backend/pkg/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var request model.CreatePostRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to decode request data", http.StatusBadRequest)
		return
	}

	// check auth and get userid from cookie
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
		http.Error(w, "Error confirming authentication: " + err.Error(), http.StatusInternalServerError)
		return
	}

	// Creates the post in database
	postID, err := repository.CreatePost(sqlite.Dbase, request, userID)
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

func confirmAuthentication(cookie *http.Cookie) (int, error) {
	sessionToken := cookie.Value

	session, err := repository.GetSessionBySessionToken(sqlite.Dbase, sessionToken)
	if err != nil {
		fmt.Println("Could not get session token: ", err)
		return 0, err
	}
	if time.Now().After(session.ExpiresAt) {
		return 0, fmt.Errorf("Session token expired: %v", session.ExpiresAt)
	}
	return session.UserID, nil
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the request body for updating the post
	var request model.UpdatePostRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to decode request data", http.StatusBadRequest)
		return
	}

	// check auth and get userid from cookie
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

	// Update the post in the database
	err = repository.UpdatePost(sqlite.Dbase, request.Id, userID, request)
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


func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
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

	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "Error authenticating user", http.StatusBadRequest)
		return
	}
	userId, err := confirmAuthentication(cookie)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Delete the post from the database
	err = repository.DeletePost(sqlite.Dbase, intpostID, userId)
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


func GetAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	// check auth and get userid from cookie
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

	posts, err := repository.GetAllPostsWithUserIDAccess(sqlite.Dbase, userID)
	if err != nil {
		http.Error(w, "Failed to retrieve posts: " + err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}