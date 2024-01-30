package handler

import (
	"backend/pkg/db/sqlite"
	"backend/pkg/model"
	"backend/pkg/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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
		fmt.Println("Error confirming authentication: " + err.Error(), http.StatusInternalServerError)
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