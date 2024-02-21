package handler

import (
	"backend/pkg/model"
	"backend/util"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// LoginHandler handles the login request.
// It decodes the login data from the request body and validates the user's credentials.
// If the credentials are valid, it generates a session token, stores it in the database, and sets a cookie with the session token.
// Finally, it sends a success response indicating that the login was successful.
func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var logData model.LoginData

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&logData)
	if err != nil {
		http.Error(w, "Error parsing login JSON data: "+err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.userRepo.GetUserByEmailOrNickname(logData.Username)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	// Compare hashed password with provided password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logData.Password))
	if err != nil { // Wrong password
		fmt.Println("Error comparing password: ", err)
		http.Error(w, "Couldn't compare password with hashed variant: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Generate a session token and store it in database
	sessionToken := util.GenerateSessionToken()
	h.sessionRepo.StoreSessionInDB(sessionToken, user.Id)

	// Set a cookie with the session
	http.SetCookie(w, &http.Cookie{
		Name:   "session_token",
		Value:  sessionToken,
		MaxAge: 60 * 15, // 15 minutes
		Path:   "/",     // Make cookie available for all paths
	})

	// Send a success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
	})
}

// LogoutHandler handles the logout functionality by deleting the session-token cookie and sending a success response.
// If the session-token cookie is not found or there is an error, it returns a bad request error.
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session_token")
	if err != nil && err != http.ErrNoCookie {
		http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Delete the session-token cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		MaxAge:  -1, // Setting MaxAge to -1 immediately expires the cookie
		Expires: time.Unix(0, 0),
	})

	// Send a success reponse
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logout successful",
	})
}

// CheckAuth checks the authentication status of the user.
// It retrieves the session token from the request cookie and verifies it against the session stored in the database.
// If the session token is valid and not expired, the user is considered authenticated.
// The authentication status is then returned as a JSON response.
func (h *UserHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
	isAuthenticated := true
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the session cookie doesn't exist, set isAuthenticated to false
			isAuthenticated = false
		} else {
			http.Error(w, "Error checking session token: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if isAuthenticated {
		sessionToken := cookie.Value

		// Get the session from database by the session token
		session, err := h.sessionRepo.GetSessionBySessionToken(sessionToken)
		if err != nil {
			if err == sql.ErrNoRows {
				isAuthenticated = false
			} else {
				fmt.Println("Error getting session token from database: " + err.Error())
				return
			}
		}
		if time.Now().After(session.ExpiresAt) {
			isAuthenticated = false
		}
	}
	response := model.AuthResponse{
		IsAuthenticated: isAuthenticated,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error marshalling response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
