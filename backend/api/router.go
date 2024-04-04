package api

import (
	"backend/pkg/handler"
	"backend/pkg/repository"
	"backend/pkg/ws"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	voteRepository := repository.NewVoteRepository(db)

	notificationHandler := handler.NewNotificationHandler(notificationRepository, sessionRepository, groupMemberRepository, groupRepository, userRepository, invitationRepository, eventRepository)
	voteHandler := handler.NewVoteHandler(voteRepository, sessionRepository)
	chatRepository := ws.NewChatRepository(db)

	chatHandler := ws.NewChatHandler(chatRepository, sessionRepository)
	hub := ws.NewHub(chatHandler)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		hub.ServeWs(w, r)
	})

	userHandler := handler.NewUserHandler(userRepository, sessionRepository, friendsRepository)
	mux.HandleFunc("/api/users/register", userHandler.UserRegisterHandler).Methods("POST")
	// User login and logout
	mux.HandleFunc("/api/users/logout", handler.LogoutHandler).Methods("POST")
	mux.HandleFunc("/api/users/login", userHandler.LoginHandler).Methods("POST")
	mux.HandleFunc("/api/users/check-auth", userHandler.CheckAuth)
	mux.HandleFunc("/api/users/auth-update", userHandler.UpdateAuth).Methods("PUT")
	mux.HandleFunc("/api/users/list", userHandler.ListUsersHandler).Methods("GET")

	// Posts
	postHandler := handler.NewPostHandler(postRepository, sessionRepository, friendsRepository, groupMemberRepository, userRepository, voteHandler)
	mux.HandleFunc("/posts", postHandler.GetAllPostsHandler).Methods("GET") // Main feed, all public posts + user groups posts
	mux.HandleFunc("/post", postHandler.CreatePostHandler).Methods("POST")
	// mux.HandleFunc("/post/{id}", handler.GetPostByIDHandler).Methods("GET")
	mux.HandleFunc("/post/{id}", postHandler.EditPostHandler).Methods("PUT")      // Edit a post
	mux.HandleFunc("/post/{id}", postHandler.DeletePostHandler).Methods("DELETE") // Delete a post
	mux.HandleFunc("/groups/{groupId}/posts", postHandler.GetPostsByGroupIDHandler).Methods("GET")

	// Profile
	mux.HandleFunc("/profile/users/{id}", userHandler.GetUserProfileByIDHandler).Methods("GET")
	mux.HandleFunc("/profile/users/{id}", userHandler.EditUserProfileHandler).Methods("PUT")
	// Profile feed, all posts by user
	mux.HandleFunc("/profile/posts/{id}", postHandler.GetAllUserPostsHandler).Methods("GET")

	// Comments
	commentHandler := handler.NewCommentHandler(commentRepository, sessionRepository, notificationHandler, postRepository, userRepository, voteHandler)
	mux.HandleFunc("/post/{id}/comments", commentHandler.GetCommentsByPostID).Methods("GET")
	mux.HandleFunc("/post/{id}/comment", commentHandler.CreateCommentHandler).Methods("POST")
	mux.HandleFunc("/post/comment", commentHandler.CreateCommentHandler).Methods("POST")
	mux.HandleFunc("/post/comment/{id}", commentHandler.DeleteCommentHandler).Methods("DELETE")

	// Likes & dislikes for comments and posts ... the getPosts and getComments methods return the number of likes and dislikes with each post/comment
	mux.HandleFunc("/vote", voteHandler.VotePostOrCommentHandler).Methods("POST")

	// Groups
	groupHandler := handler.NewGroupHandler(groupRepository, sessionRepository, groupMemberRepository, notificationHandler, userRepository, friendsRepository)
	mux.HandleFunc("/groups", groupHandler.GetAllGroupsHandler).Methods("GET")
	mux.HandleFunc("/groups", groupHandler.CreateGroupHandler).Methods("POST")
	mux.HandleFunc("/groups/{id}", groupHandler.GetGroupByIDHandler).Methods("GET")
	mux.HandleFunc("/groups/{id}", groupHandler.EditGroupHandler).Methods("PUT")
	mux.HandleFunc("/groups/{id}", groupHandler.DeleteGroupHandler).Methods("DELETE")

	// Group invitations & requests
	groupMemberHandler := handler.NewGroupMemberHandler(groupMemberRepository, invitationRepository, sessionRepository, notificationHandler, groupRepository, userRepository)
	// for all members
	mux.HandleFunc("/invitations/invite/{groupId}/{userId}", groupMemberHandler.InviteGroupMemberHandler).Methods("POST")
	// for group member
	mux.HandleFunc("/invitations", groupMemberHandler.GetAllGroupInvitationsHandler).Methods("GET")
	mux.HandleFunc("/invitations/{groupId}", groupMemberHandler.GetGroupInvitationByIDHandler).Methods("GET")
	mux.HandleFunc("/invitations/decline/{groupId}", groupMemberHandler.DeclineGroupInvitationHandler).Methods("POST")
	mux.HandleFunc("/invitations/accept/{groupId}", groupMemberHandler.AcceptGroupInvitationHandler).Methods("POST")
	mux.HandleFunc("/invitations/request/{groupId}", groupMemberHandler.RequestGroupMembershipHandler).Methods("POST")
	// for group owner
	mux.HandleFunc("/groups/{groupId}/non-members", groupMemberHandler.GetAllNonMembersHandler).Methods("GET")
	mux.HandleFunc("/groups/{groupId}/members", groupMemberHandler.GetAllMembersHandler).Methods("GET")
	mux.HandleFunc("/groups/{groupId}/members/{userId}", groupMemberHandler.RemoveMemberHandler).Methods("DELETE")
	mux.HandleFunc("/invitations/approve/{groupId}/{userId}", groupMemberHandler.ApproveGroupMembershipHandler).Methods("PUT")
	mux.HandleFunc("/invitations/decline/{groupId}/{userId}", groupMemberHandler.DeclineGroupMembershipHandler).Methods("PUT")
	mux.HandleFunc("/groups/{groupId}/requests", groupMemberHandler.GetAllGroupRequestsHandler).Methods("GET")

	// Events
	eventHandler := handler.NewEventHandler(eventRepository, sessionRepository, groupMemberRepository, userRepository, notificationHandler)
	mux.HandleFunc("/events/group/{groupId}", eventHandler.GetAllGroupEventsHandler).Methods("GET")
	mux.HandleFunc("/events", eventHandler.CreateEventHandler).Methods("POST")
	mux.HandleFunc("/events/{id}", eventHandler.EditEventHandler).Methods("PUT")
	mux.HandleFunc("/events/{id}", eventHandler.DeleteEventHandler).Methods("DELETE")
	mux.HandleFunc("/events/{id}", eventHandler.GetEventsByGroupIDHandler).Methods("GET")
	mux.HandleFunc("/events/{eventId}/{status}", eventHandler.AddOrUpdateAttendanceHandler).Methods("PUT")
	mux.HandleFunc("/events/attendance/{eventId}", eventHandler.GetAttendanceByEventIDHandler).Methods("GET")

	// Notifications
	mux.HandleFunc("/notifications", notificationHandler.GetAllNotificationsForUserHandler).Methods("GET")
	mux.HandleFunc("/notifications/{id}", notificationHandler.GetNotificationByIDHandler).Methods("GET")
	mux.HandleFunc("/notifications/{id}", notificationHandler.MarkNotificationAsReadHandler).Methods("PUT")
	mux.HandleFunc("/notifications/{id}", notificationHandler.DeleteNotificationHandler).Methods("DELETE")

	// Friends
	friendHandler := handler.NewFriendHandler(friendsRepository, sessionRepository, notificationHandler, userRepository)
	mux.HandleFunc("/friends/requests", friendHandler.GetFriendRequestsHandler).Methods("GET")
	mux.HandleFunc("/friends/request/{id}", friendHandler.SendFriendRequestHandler).Methods("POST")
	mux.HandleFunc("/friends/accept/{id}", friendHandler.AcceptFriendRequestHandler).Methods("POST")
	mux.HandleFunc("/friends/decline/{id}", friendHandler.DeclineFriendRequestHandler).Methods("POST")
	mux.HandleFunc("/friends/check/{id}", friendHandler.CheckFriendStatusHandler).Methods("GET")

	mux.HandleFunc("/friends/{id}", friendHandler.GetFriendsHandler).Methods("GET")

	// route to serve images
	http.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/images/", http.FileServer(http.Dir("./pkg/db/images"))).ServeHTTP(w, r)
	})

	go hub.Run()

	address := os.Getenv("NEXT_PUBLIC_URL")
	port := os.Getenv("NEXT_PUBLIC_HTTPS_PORT")
	if address == "" {
		address = "http://localhost" // fallback address
	} else if port == "" {
		port = "3000" // fallback port
	}

	fmt.Println("Cors address: " + address + ":" + port)

	// CORS
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{address + ":" + port},                      // Replace with your frontend's origin
		AllowCredentials: true,                                                // Important for cookies, authorization headers with HTTPS
		AllowedHeaders:   []string{"Authorization", "Content-Type"},           // You can adjust this based on your needs
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Adjust the methods based on your requirements
		// You can include other settings like ExposedHeaders, MaxAge, etc., according to your needs
	})
	mux_cors := corsOptions.Handler(mux)
	http.Handle("/", mux_cors)
}
