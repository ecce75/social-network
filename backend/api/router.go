package api

import (
	"backend/pkg/handler"
	"backend/pkg/middleware"
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

    // Groups
    groupRepo := repository.NewGroupRepository(db)
    groupHandler := handler.NewGroupHandler(groupRepo)
    mux.Handle("/groups", middleware.CheckAuthMiddleware(http.HandlerFunc(groupHandler.GetAllGroupsHandler))).Methods("GET")
    mux.Handle("/groups", middleware.CheckAuthMiddleware(http.HandlerFunc(groupHandler.CreateGroupHandler))).Methods("POST")
    mux.Handle("/groups/{id}", middleware.CheckAuthMiddleware(http.HandlerFunc(groupHandler.GetGroupByIDHandler))).Methods("GET")
    mux.Handle("/groups/{id}", middleware.CheckAuthMiddleware(http.HandlerFunc(groupHandler.EditGroupHandler))).Methods("PUT")
    mux.Handle("/groups/{id}", middleware.CheckAuthMiddleware(http.HandlerFunc(groupHandler.DeleteGroupHandler))).Methods("DELETE")

    // Group invitations & requests
    invitationRepo := repository.NewInvitationRepository(db)
    invitationHandler := handler.NewInvitationHandler(invitationRepo)
    groupMemberRepo := repository.NewGroupMemberRepository(db)
    groupMemberHandler := handler.NewGroupMemberHandler(groupMemberRepo)
    mux.Handle("/invitations", middleware.CheckAuthMiddleware(http.HandlerFunc(invitationHandler.GetAllGroupInvitationsHandler))).Methods("GET")
    mux.Handle("/invitations", middleware.CheckAuthMiddleware(http.HandlerFunc(invitationHandler.InviteGroupMemberHandler))).Methods("POST")
    mux.Handle("/invitations/{id}", middleware.CheckAuthMiddleware(http.HandlerFunc(invitationHandler.GetGroupInvitationByIDHandler))).Methods("GET")
    mux.Handle("/invitations/{id}", middleware.CheckAuthMiddleware(http.HandlerFunc(invitationHandler.DeclineGroupInvitationHandler))).Methods("PUT")
    mux.Handle("/invitations/{id}", middleware.CheckAuthMiddleware(http.HandlerFunc(groupMemberHandler.AcceptGroupInvitationHandler))).Methods("PUT")
    mux.Handle("/invitations/request/{id}", middleware.CheckAuthMiddleware(http.HandlerFunc(groupMemberHandler.RequestGroupMembershipHandler))).Methods("POST")
    mux.Handle("/groups/{groupId}/members/{userId}", middleware.CheckAuthMiddleware(http.HandlerFunc(groupMemberHandler.RemoveMemberHandler))).Methods("DELETE")
    mux.Handle("/invitations/approve/{id}", middleware.CheckAuthMiddleware(http.HandlerFunc(groupMemberHandler.ApproveGroupMembershipHandler))).Methods("PUT")

    // TODO: Group posts & comments

    // TODO: Group events
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