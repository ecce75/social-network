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
	// 1. Parse the multipart form data from the request
	err1 := r.ParseMultipartForm(10 << 20) // Maximum memory 10MB, change this based on your requirements
	if err1 != nil {
		http.Error(w, "Error parsing form data: "+err1.Error(), http.StatusBadRequest)
		return
	}
	var regData model.RegistrationData // variable based on struct fields for storing registration data

	regData.Username = r.FormValue("username")
	regData.Email = r.FormValue("email")
	regData.Password = r.FormValue("password")
	regData.FirstName = r.FormValue("first_name")
	regData.LastName = r.FormValue("last_name")
	regData.DOB = r.FormValue("dob")
	regData.About = r.FormValue("about")
	

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regData.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	// Change variable password data to hashed variant
	regData.Password = string(hashedPassword)

	util.ImageSave(w, r, &regData) // parses image data from request to the variable
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