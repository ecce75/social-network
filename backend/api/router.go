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
    mux.HandleFunc("/post", handler.GetAllPostsHandler).Methods("GET")
    mux.HandleFunc("/post", handler.CreatePostHandler).Methods("POST")
    //mux.HandleFunc("/post/{id}", handler.GetPostByIDHandler).Methods("GET")
    mux.HandleFunc("/post/{id}", handler.EditPostHandler).Methods("PUT")    // Edit a post
    mux.HandleFunc("/post/{id}", handler.DeletePostHandler).Methods("DELETE") // Delete a post

    // Comments
    mux.HandleFunc("/post/{id}/comments", handler.GetCommentByUserIDorPostID).Methods("GET")
    mux.HandleFunc("/comment", handler.CreateCommentHandler).Methods("POST")
    mux.HandleFunc("/comment/{id}", handler.DeleteCommentHandler).Methods("DELETE")

    // TODO: Groups
    mux.HandleFunc("/groups", handler.GetAllGroupsHandler).Methods("GET")
    mux.HandleFunc("/groups", handler.CreateGroupHandler).Methods("POST")
    mux.HandleFunc("/groups/{id}", handler.GetGroupByIDHandler).Methods("GET")
    mux.HandleFunc("/groups/{id}", handler.EditGroupHandler).Methods("PUT")
    mux.HandleFunc("/groups/{id}", handler.DeleteGroupHandler).Methods("DELETE")

    // TODO: Invitations
    mux.HandleFunc("/invitations", handler.GetAllInvitationsHandler).Methods("GET")
    mux.HandleFunc("/invitations", handler.CreateInvitationHandler).Methods("POST")
    mux.HandleFunc("/invitations/{id}", handler.GetInvitationByIDHandler).Methods("GET")
    mux.HandleFunc("/invitations/{id}", handler.AcceptInvitationHandler).Methods("PUT")
    mux.HandleFunc("/invitations/{id}", handler.DeclineInvitationHandler).Methods("PUT")

    // TODO: Events
    mux.HandleFunc("/events", handler.GetAllEventsHandler).Methods("GET")
    mux.HandleFunc("/events", handler.CreateEventHandler).Methods("POST")
    mux.HandleFunc("/events/{id}", handler.GetEventByIDHandler).Methods("GET")
    mux.HandleFunc("/events/{id}", handler.EditEventHandler).Methods("PUT")
    mux.HandleFunc("/events/{id}", handler.DeleteEventHandler).Methods("DELETE")

    // TODO: Notifications
    mux.HandleFunc("/notifications", handler.GetAllNotificationsHandler).Methods("GET")
    mux.HandleFunc("/notifications", handler.CreateNotificationHandler).Methods("POST")
    mux.HandleFunc("/notifications/{id}", handler.GetNotificationByIDHandler).Methods("GET")
    mux.HandleFunc("/notifications/{id}", handler.MarkNotificationAsReadHandler).Methods("PUT")

	// TODO: Friends

	mux.HandleFunc("/friends/request", handler.SendFriendRequestHandler).Methods("POST")
	mux.HandleFunc("/friends/accept", handler.AcceptFriendRequestHandler).Methods("POST")
	mux.HandleFunc("/friends/decline", handler.DeclineFriendRequestHandler).Methods("POST")
	mux.HandleFunc("/friends/block", handler.BlockUserHandler).Methods("POST")
	mux.HandleFunc("/friends/unblock", handler.UnblockUserHandler).Methods("POST")
	mux.HandleFunc("/friends", handler.GetFriendsHandler).Methods("GET")

    // Catch-all route to serve index.html for all other routes
	// TODO: remove
    mux.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "../frontend/public/index.html")
        fmt.Println("route called successfully")
    })
    http.Handle("/", mux)
}