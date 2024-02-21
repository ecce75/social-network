package api

import (
	"backend/pkg/handler"
	"backend/pkg/repository"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

// API layer, handlers, and routing
func Router(mux *mux.Router, db *sql.DB) {
	// User registration requires input in the form like RegistrationData struct at /pkg/model/stucts.go
	userRepository := repository.NewUserRepository(db)
	postRepository := repository.NewPostRepository(db)
	commentRepository := repository.NewCommentRepository(db)
	groupRepository := repository.NewGroupRepository(db)
	groupMemberRepository := repository.NewGroupMemberRepository(db)
	invitationRepository := repository.NewInvitationRepository(db)
	notificationRepository := repository.NewNotificationRepository(db)
	eventRepository := repository.NewEventRepository(db)
	sessionRepository := repository.NewSessionRepository(db)
	friendsRepository := repository.NewFriendsRepository(db)

	userHandler := handler.NewUserHandler(userRepository, sessionRepository, friendsRepository)
	mux.HandleFunc("/api/users/register", userHandler.UserRegisterHandler).Methods("POST")
	// User login and logout
	mux.HandleFunc("/api/users/logout", handler.LogoutHandler).Methods("POST")
	mux.HandleFunc("/api/users/login", userHandler.LoginHandler).Methods("POST")
	mux.HandleFunc("/api/users/check-auth", userHandler.CheckAuth)
	mux.HandleFunc("/api/users/list", userHandler.ListUsersHandler).Methods("GET")

	// Posts
	postHandler := handler.NewPostHandler(postRepository, sessionRepository, friendsRepository, groupMemberRepository)
	mux.HandleFunc("/post", postHandler.GetAllPostsHandler).Methods("GET") // Main feed, all public posts + user groups posts
	mux.HandleFunc("/post", postHandler.CreatePostHandler).Methods("POST")
	// mux.HandleFunc("/post/{id}", handler.GetPostByIDHandler).Methods("GET")
	mux.HandleFunc("/post/{id}", postHandler.EditPostHandler).Methods("PUT")      // Edit a post
	mux.HandleFunc("/post/{id}", postHandler.DeletePostHandler).Methods("DELETE") // Delete a post
	mux.HandleFunc("/groups/posts/{id}", postHandler.GetPostsByGroupIDHandler).Methods("GET")

	// Profile
	mux.HandleFunc("/profile/users/{id}", userHandler.GetUserProfileByIDHandler).Methods("GET")
	mux.HandleFunc("/profile/users/{id}", userHandler.EditUserProfileHandler).Methods("PUT")
	// Profile feed, all posts by user
	mux.HandleFunc("/profile/posts/{id}", postHandler.GetAllUserPostsHandler).Methods("GET")

	// Comments
	commentHandler := handler.NewCommentHandler(commentRepository, sessionRepository)
	mux.HandleFunc("/post/{id}/comments", commentHandler.GetCommentsByUserIDorPostID).Methods("GET")
	mux.HandleFunc("/post/comment", commentHandler.CreateCommentHandler).Methods("POST")
	mux.HandleFunc("/post/comment/{id}", commentHandler.DeleteCommentHandler).Methods("DELETE")

	// Groups
	groupHandler := handler.NewGroupHandler(groupRepository, sessionRepository)
	mux.HandleFunc("/groups", groupHandler.GetAllGroupsHandler).Methods("GET")
	mux.HandleFunc("/groups", groupHandler.CreateGroupHandler).Methods("POST")
	mux.HandleFunc("/groups/{id}", groupHandler.GetGroupByIDHandler).Methods("GET")
	mux.HandleFunc("/groups/{id}", groupHandler.EditGroupHandler).Methods("PUT")
	mux.HandleFunc("/groups/{id}", groupHandler.DeleteGroupHandler).Methods("DELETE")

	// Group invitations & requests
	groupMemberHandler := handler.NewGroupMemberHandler(groupMemberRepository, invitationRepository, sessionRepository)
	mux.HandleFunc("/invitations", groupMemberHandler.GetAllGroupInvitationsHandler).Methods("GET")
	mux.HandleFunc("/invitations", groupMemberHandler.InviteGroupMemberHandler).Methods("POST")
	mux.HandleFunc("/invitations/{id}", groupMemberHandler.GetGroupInvitationByIDHandler).Methods("GET")
	mux.HandleFunc("/invitations/{id}", groupMemberHandler.DeclineGroupInvitationHandler).Methods("PUT")
	mux.HandleFunc("/invitations/{id}", groupMemberHandler.AcceptGroupInvitationHandler).Methods("PUT")
	mux.HandleFunc("/invitations/request/{id}", groupMemberHandler.RequestGroupMembershipHandler).Methods("POST")
	mux.HandleFunc("/groups/{groupId}/members/{userId}", groupMemberHandler.RemoveMemberHandler).Methods("DELETE")
	mux.HandleFunc("/invitations/approve/{id}", groupMemberHandler.ApproveGroupMembershipHandler).Methods("PUT")

	// Events
	eventHandler := handler.NewEventHandler(eventRepository, sessionRepository, groupMemberRepository)
	mux.HandleFunc("/events", eventHandler.GetAllEventsHandler).Methods("GET")
	mux.HandleFunc("/events", eventHandler.CreateEventHandler).Methods("POST")
	mux.HandleFunc("/events/{id}", eventHandler.GetEventByIDHandler).Methods("GET")
	mux.HandleFunc("/events/{id}", eventHandler.EditEventHandler).Methods("PUT")
	mux.HandleFunc("/events/{id}", eventHandler.DeleteEventHandler).Methods("DELETE")
	mux.HandleFunc("/events/{id}", eventHandler.GetEventsByGroupIDHandler).Methods("GET")
	mux.HandleFunc("/events/{id}", eventHandler.AddOrUpdateAttendanceHandler).Methods("PUT")
	mux.HandleFunc("/events/{id}", eventHandler.GetAttendanceByEventIDHandler).Methods("GET")

	// Notifications
	notificationHandler := handler.NewNotificationHandler(notificationRepository, sessionRepository)
	mux.HandleFunc("/notifications", notificationHandler.GetAllNotificationsHandler).Methods("GET")
	mux.HandleFunc("/notifications", notificationHandler.CreateNotificationHandler).Methods("POST")
	mux.HandleFunc("/notifications/{id}", notificationHandler.GetNotificationByIDHandler).Methods("GET")
	mux.HandleFunc("/notifications/{id}", notificationHandler.MarkNotificationAsReadHandler).Methods("PUT")

	// Friends
	friendHandler := handler.NewFriendHandler(friendsRepository, sessionRepository)
	mux.HandleFunc("/friends/request/{id}", friendHandler.SendFriendRequestHandler).Methods("POST")
	mux.HandleFunc("/friends/accept/{id}", friendHandler.AcceptFriendRequestHandler).Methods("POST")
	mux.HandleFunc("/friends/decline", friendHandler.DeclineFriendRequestHandler).Methods("POST")
	mux.HandleFunc("/friends/block", friendHandler.BlockUserHandler).Methods("POST")
	mux.HandleFunc("/friends/unblock", friendHandler.UnblockUserHandler).Methods("POST")
	mux.HandleFunc("/friends/check/{id}", friendHandler.CheckFriendStatusHandler).Methods("GET")

	mux.HandleFunc("/friends", friendHandler.GetFriendsHandler).Methods("GET")

	// route to serve images
	http.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/images/", http.FileServer(http.Dir("./pkg/db/images"))).ServeHTTP(w, r)
	})
	// CORS
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},                   // Replace with your frontend's origin
		AllowCredentials: true,                                                // Important for cookies, authorization headers with HTTPS
		AllowedHeaders:   []string{"Authorization", "Content-Type"},           // You can adjust this based on your needs
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Adjust the methods based on your requirements
		// You can include other settings like ExposedHeaders, MaxAge, etc., according to your needs
	})
	mux_cors := corsOptions.Handler(mux)
	http.Handle("/", mux_cors)
}
