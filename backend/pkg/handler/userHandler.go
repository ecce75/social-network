package handler

import (
	"backend/pkg/db/sqlite"
	"backend/pkg/model"
	"backend/pkg/repository"
	"backend/util"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	var regData model.RegistrationData // variable based on struct fields for storing registration data

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&regData) // decodes json data from request to the variable
	if err != nil {
		http.Error(w, "Error parsing JSON: "+ err.Error(), http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regData.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	// Change variable password data to hashed variant
	regData.Password = string(hashedPassword)

	// Store user in database
	userID, err := repository.RegisterUser(sqlite.Dbase, regData)
	if err != nil {
		http.Error(w, "Error registering user: "+ err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate a session token and store it in database with expiration time
	sessionToken := util.GenerateSessionToken()
	repository.StoreSessionInDB(sqlite.Dbase, sessionToken, int(userID))

	// Set a cookie with the session token
	http.SetCookie(w, &http.Cookie{
		Name: "session_token",
		Value: sessionToken,
		MaxAge: 60*15, // 15 minutes
	})

	// Send a success response
	response := map[string]interface{}{
		"message": "User registration successful",
		"userID": userID, // TODO: remove in production
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}