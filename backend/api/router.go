package api

import (
	"backend/pkg/handler"
	"backend/pkg/repository"
	"database/sql"
	"net/http"
	"github.com/gorilla/mux"
)

// API layer, handlers, and routing
func Router(mux *mux.Router, db *sql.DB) {
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
    groupRepo := repository.NewGroupRepository(db)
    groupHandler := handler.NewGroupHandler(groupRepo)
    mux.HandleFunc("/groups", groupHandler.GetAllGroupsHandler).Methods("GET")
    mux.HandleFunc("/groups", groupHandler.CreateGroupHandler).Methods("POST")
    mux.HandleFunc("/groups/{id}", groupHandler.GetGroupByIDHandler).Methods("GET")
    mux.HandleFunc("/groups/{id}", groupHandler.EditGroupHandler).Methods("PUT")
    mux.HandleFunc("/groups/{id}", groupHandler.DeleteGroupHandler).Methods("DELETE")

    // TODO: Group invitations & requests
    invitationRepo := repository.NewInvitationRepository(db)
    invitationHandler := handler.NewInvitationHandler(invitationRepo)
    mux.HandleFunc("/invitations", invitationHandler.GetAllGroupInvitationsHandler).Methods("GET")
    mux.HandleFunc("/invitations", invitationHandler.CreateGroupInvitationHandler).Methods("POST")
    mux.HandleFunc("/invitations/{id}", invitationHandler.GetGroupInvitationByIDHandler).Methods("GET")
    mux.HandleFunc("/invitations/{id}", invitationHandler.AcceptGroupInvitationHandler).Methods("PUT")
    mux.HandleFunc("/invitations/{id}", invitationHandler.DeclineGroupInvitationHandler).Methods("PUT")

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

    // ----
    http.Handle("/", mux)
}