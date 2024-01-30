package api

import (
	"fmt"
	"net/http"
	"backend/pkg/handler"
	"github.com/gorilla/mux"
)

// API layer, handlers, and routing
func Router(mux *mux.Router) {
	// User registration requires input in the form like RegistrationData struct at /pkg/model/stucts.go
	mux.HandleFunc("/api/users/register", handler.UserRegisterHandler).Methods("POST")
	
	// User login and logout
	mux.HandleFunc("/api/users/logout", handler.LogoutHandler).Methods("POST")
	mux.HandleFunc("/api/users/login", handler.LoginHandler).Methods("POST")
	mux.HandleFunc("/api/users/check-auth", handler.CheckAuth)

	// Posts

	mux.HandleFunc("/post", handler.CreatePostHandler).Methods("POST")

	// Comments

	// Groups

	// Invitations

	// Events

	// Notifications

	// Catch-all route to serve index.html for all other routes
	mux.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/public/index.html")
		fmt.Println("route called successfully")
	})
	http.Handle("/", mux)
}