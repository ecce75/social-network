package handler

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"backend/util"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userRepo    *repository.UserRepository
	sessionRepo *repository.SessionRepository
	friendsRepo *repository.FriendsRepository
}

func NewUserHandler(uRepo *repository.UserRepository, sRepo *repository.SessionRepository, fRepo *repository.FriendsRepository) *UserHandler {
	return &UserHandler{userRepo: uRepo, sessionRepo: sRepo, friendsRepo: fRepo}
}

func (h *UserHandler) UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
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
	// Change input password data to hashed variant
	regData.Password = string(hashedPassword)

	util.ImageSave(w, r, &regData) // parses image data from request to the variable
	// Store user in database
	userID, err := h.userRepo.RegisterUser(regData)
	if err != nil {
		http.Error(w, "Error registering user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate a session token and store it in database with expiration time
	sessionToken := util.GenerateSessionToken()
	h.sessionRepo.StoreSessionInDB(sessionToken, int(userID))

	// Set a cookie with the session token
	http.SetCookie(w, &http.Cookie{
		Name:   "session_token",
		Value:  sessionToken,
		MaxAge: 60 * 15, // 15 minutes
		Path:   "/",     // Make cookie available for all paths
	})

	// Send a success response
	response := map[string]interface{}{
		"message": "User registration successful",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) GetUserProfileByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL
	requestUserID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Invalid user ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	userID, ok := vars["id"]
	if !ok {
		http.Error(w, "User ID not found in URL", http.StatusBadRequest)
		return
	}
	if userID == "me" {
		// No ID in URL, use the logged-in user's ID
		userID = strconv.Itoa(requestUserID)
	}

	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID format: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Get the user profile from the database
	profile, err := h.userRepo.GetUserProfileByID(intUserID)
	if err != nil {
		http.Error(w, "Error getting user profile: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if profile.ProfileSetting == "private" {
		// check whether the user requesting is the same as the user profile
		status, err := h.friendsRepo.GetFriendStatus(requestUserID, intUserID)
		if err != nil {
			http.Error(w, "Error getting friend status: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if status == "accepted" {
			// if the user is friends with the user profile, return the profile
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(profile)
		} else {
			// TODO: return profile with restricted access
			http.Error(w, "User profile is private", http.StatusUnauthorized)
			return
		}
		// or friends with the user profile, depending on that return the profile
	} else if intUserID == requestUserID {
		// if the user is the same as the user profile, return the profile
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(profile)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func (h *UserHandler) EditUserProfileHandler(w http.ResponseWriter, r *http.Request) {
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
	regData.ProfileSetting = r.FormValue("profile_setting")

	// get userid from cookie
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error getting user id: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regData.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	// Change input password data to hashed variant
	regData.Password = string(hashedPassword)

	util.ImageSave(w, r, &regData) // parses image data from request to the variable
	// Store user in database
	err = h.userRepo.UpdateUserProfile(userID, regData)
	if err != nil {
		http.Error(w, "Error updating user profile: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a success response
	response := map[string]interface{}{
		"message": "User profile updated",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	// get userid from cookie
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error getting user id: "+err.Error(), http.StatusInternalServerError)
		return
	}

	users, err := h.userRepo.GetAllUsersExcludeRequestingUserAndFriends(userID)
	if err != nil {
		http.Error(w, "Error listing users: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
