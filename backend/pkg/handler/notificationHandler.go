package handler

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"backend/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// NotificationHandler handles HTTP requests related to notifications.
type NotificationHandler struct {
	notificationRepo *repository.NotificationRepository
	sessionRepo      *repository.SessionRepository
	groupMemberRepo  *repository.GroupMemberRepository
	groupRepo        *repository.GroupRepository
	userRepo         *repository.UserRepository
	invitationRepo   *repository.InvitationRepository
	eventRepo        *repository.EventRepository
}

// NewNotificationHandler creates a new instance of NotificationHandler.
// It takes a NotificationRepository and a SessionRepository as parameters.
// Returns a pointer to the newly created NotificationHandler.
func NewNotificationHandler(notificationRepo *repository.NotificationRepository, sessionRepo *repository.SessionRepository, groupMemberRepo *repository.GroupMemberRepository, groupRepo *repository.GroupRepository, userRepo *repository.UserRepository, invitationRepo *repository.InvitationRepository, eventRepo *repository.EventRepository) *NotificationHandler {
	return &NotificationHandler{notificationRepo: notificationRepo, sessionRepo: sessionRepo, groupMemberRepo: groupMemberRepo, groupRepo: groupRepo, userRepo: userRepo, invitationRepo: invitationRepo, eventRepo: eventRepo}
}

func (h *NotificationHandler) CreateNotification(userID, senderID int, messageType, message string) error {
	notification := model.Notification{
		UserId:   userID,
		SenderId: senderID,
		Type:     messageType,
		Message:  message,
	}
	_, err := h.notificationRepo.CreateNotification(notification)
	return err
}
func (h *NotificationHandler) EditFriendRequestNotification(userID, senderID int, message string) error {
	return h.notificationRepo.EditFriendNotificationMessage(userID, senderID, message)
}

func (h *NotificationHandler) EditGroupRequestNotification(userID, groupID int, message string) error {
	return h.notificationRepo.EditGroupNotificationMessage(userID, groupID, message)
}

func (h *NotificationHandler) CreateGroupNotification(userID, groupID int, message string) error {
	notification := model.Notification{
		UserId:  userID,
		GroupId: groupID,
		Type:    "group",
		Message: message,
	}
	_, err := h.notificationRepo.CreateNotification(notification)
	return err
}

func (h *NotificationHandler) DeleteNotificationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid notification ID", http.StatusBadRequest)
	}

	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "User not authenticated: "+err.Error(), http.StatusUnauthorized)
	}
	notification, err := h.notificationRepo.GetNotificationByID(id)
	if err != nil {
		http.Error(w, "Failed to get notification: "+err.Error(), http.StatusInternalServerError)
	}
	if notification.UserId != userID {
		http.Error(w, "User not authorized to delete this notification", http.StatusUnauthorized)
		return
	}

	err = h.notificationRepo.DeleteNotification(id)
	if err != nil {
		http.Error(w, "Failed to delete notification: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *NotificationHandler) GetAllNotificationsForUserHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "User not authenticated: "+err.Error(), http.StatusUnauthorized)
	}

	notifications, err := h.notificationRepo.GetNotificationsByUserId(userID)
	if err != nil {
		http.Error(w, "Error fetching notifications: "+err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifications)
}

// GetNotificationByIDHandler retrieves a specific notification by its ID and responds with a JSON object.
func (h *NotificationHandler) GetNotificationByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid notification ID", http.StatusBadRequest)
		return
	}

	notification, err := h.notificationRepo.GetNotificationByID(id)
	if err != nil {
		http.Error(w, "Failed to get notification: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notification)
}

// MarkNotificationAsReadHandler marks a notification as read based on its ID.
func (h *NotificationHandler) MarkNotificationAsReadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid notification ID", http.StatusBadRequest)
		return
	}

	// Mark the notification as read
	err = h.notificationRepo.MarkNotificationAsRead(id)
	if err != nil {
		http.Error(w, "Failed to mark notification as read: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with a success message or appropriate response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Notification marked as read successfully!"))
}

// func (h *NotificationHandler) NotifyPostOwnerOfNewComment(postID, commentID, userID int) error {
// 	post, err := h.notificationRepo.GetPostByID(postID)
// 	if err != nil {
// 		return err
// 	}

// 	username, err := h.userRepo.GetUsernameByID(userID)
// 	if err != nil {
// 		return err
// 	}

// }

func (h *NotificationHandler) NotifyGroupDeletion(groupID int) error {
	// Get the list of group members.
	members, err := h.groupMemberRepo.GetGroupMembers(groupID)
	if err != nil {
		return err
	}
	// Get the group title.
	groupTitle, err := h.groupRepo.GetGroupTitleByID(groupID)
	if err != nil {
		return err
	}
	// Construct a notification message.
	message := fmt.Sprintf("The group '%s' has been deleted.", groupTitle)

	// Create notifications for each group member.
	for _, member := range members {
		// Create a new notification.
		err = h.CreateNotification(member.UserID, 0, "group", message)
		if err != nil {
			return err
		}
	}

	return nil
}

// notifyGroupAdmin notifies the group admin that a user has requested to join the group.
func (h *NotificationHandler) NotifyGroupAdmin(groupID, userID int) error {
	groupTitle, err := h.groupRepo.GetGroupTitleByID(groupID)
	if err != nil {
		return err
	}

	username, err := h.userRepo.GetUsernameByID(userID)
	if err != nil {
		return err
	}
	message := username + " has requested to join your group: " + groupTitle
	// Get the group admin.
	adminID, err := h.groupMemberRepo.GetGroupAdminByID(groupID)
	if err != nil {
		return err
	}
	return h.CreateNotification(adminID, userID, "group", message)
}

// notifyUserRequestApproved notifies the user that their request was approved.
func (h *NotificationHandler) NotifyUserRequestApproved(userID, groupID int) error {
	groupTitle, err := h.groupRepo.GetGroupTitleByID(groupID)
	if err != nil {
		return err
	}
	username, _ := h.userRepo.GetUsernameByID(userID)
	message := fmt.Sprintf("You approved ", username, "'s request to join ", groupTitle)
	err = h.notificationRepo.EditGroupNotificationMessage(userID, groupID, message)
	if err != nil {
		return err
	}

	message = fmt.Sprintf("Your request to join the group %s has been approved.", groupTitle)
	return h.CreateNotification(userID, 0, "group", message)
}

// Function to notify the user about the declined request.
func (h *NotificationHandler) NotifyUserDecline(userID, groupID int) error {
	// message
	groupTitle, err := h.groupRepo.GetGroupTitleByID(groupID)
	if err != nil {
		return err
	}
	username, _ := h.userRepo.GetUsernameByID(userID)
	message := fmt.Sprintf("You declined ", username, "'s request to join ", groupTitle)
	err = h.notificationRepo.EditGroupNotificationMessage(userID, groupID, message)
	if err != nil {
		return err
	}

	message = fmt.Sprintf("Your request to join the group %s has been declined.", groupTitle)
	return h.CreateNotification(userID, 0, "group", message)
}

// notifyUserInvitation notifies the user about the group invitation.
func (h *NotificationHandler) NotifyUserInvitation(userID, groupID int) error {
	groupTitle, err := h.groupRepo.GetGroupTitleByID(groupID)
	if err != nil {
		return err
	}
	message := fmt.Sprintf("You have been invited to join the group %s.", groupTitle)
	return h.CreateGroupNotification(userID, groupID, message)
}

// notifyGroupOfNewMember notifies the group members about the new member.
func (h *NotificationHandler) NotifyGroupOfNewMember(groupID, joinUserID int) error {
	group, err := h.groupRepo.GetGroupByID(groupID)
	if err != nil {
		return err
	}

	username, err := h.userRepo.GetUsernameByID(joinUserID)
	if err != nil {
		return err
	}

	message := fmt.Sprintf("'%s' has joined the group '%s'.", username, group.Title)

	members, err := h.groupMemberRepo.GetGroupMembers(groupID)
	if err != nil {
		return err
	}

	for _, member := range members {
		if member.UserID != joinUserID {
			err := h.CreateGroupNotification(member.UserID, groupID, message)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// NotifyInvitationDecline notifies the group owner about the declined invitation.
func (h *NotificationHandler) NotifyInvitationDecline(userID, groupID int) error {
	// Get the invitation details from the repository based on the ID.
	invitation, err := h.invitationRepo.GetGroupInvitationByID(userID, groupID)
	if err != nil {
		return err
	}
	username, err := h.userRepo.GetUsernameByID(invitation.JoinUserId)
	if err != nil {
		return err
	}
	groupTitle, err := h.groupRepo.GetGroupTitleByID(invitation.GroupId)
	if err != nil {
		return err
	}
	// Construct a notification message.
	message := fmt.Sprintf("The user %s has declined your invitation to join the group %s.", username, groupTitle)

	return h.CreateGroupNotification(invitation.InviteUserId, invitation.GroupId, message)
}

func (h *NotificationHandler) NotifyGroupOfEvent(groupID, eventID int) error {
	event, err := h.eventRepo.GetEventByID(eventID)
	if err != nil {
		return err
	}

	username, err := h.userRepo.GetUsernameByID(event.CreatorId)
	if err != nil {
		return err
	}

	message := fmt.Sprintf("'%s' created an event '%s'.", username, event.Title)

	members, err := h.groupMemberRepo.GetGroupMembers(groupID)
	if err != nil {
		return err
	}

	for _, member := range members {
		err := h.CreateGroupNotification(member.UserID, groupID, message)
		if err != nil {
			return err
		}
	}
	return err
}
