package handler

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"backend/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type GroupMemberHandler struct {
	groupMemberRepo  *repository.GroupMemberRepository
	invitationRepo   *repository.InvitationRepository
	sessionRepo      *repository.SessionRepository
	notificationRepo *repository.NotificationRepository
	groupRepo        *repository.GroupRepository
}

func NewGroupMemberHandler(groupMemberRepo *repository.GroupMemberRepository, invitationRepo *repository.InvitationRepository, sessionRepo *repository.SessionRepository, notificationRepo *repository.NotificationRepository, groupRepo *repository.GroupRepository) *GroupMemberHandler {
	return &GroupMemberHandler{groupMemberRepo: groupMemberRepo, invitationRepo: invitationRepo, sessionRepo: sessionRepo, notificationRepo: notificationRepo, groupRepo: groupRepo}
}

// RemoveMemberFromGroup removes a user from a group. It takes two parameters: the ID of the group
// and the ID of the user. It deletes the row from the group_members table in the database that
// represents the user being a member of the group. If the operation is successful, it returns nil.
// If there is an error, it returns the error.
func (h *GroupMemberHandler) RemoveMemberHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupId, ok := vars["groupId"]
	if !ok {
		http.Error(w, "Missing group ID", http.StatusBadRequest)
		return
	}
	intGroupId, err := strconv.Atoi(groupId)
	if err != nil {
		http.Error(w, "Failed to convert groupid string to int: "+err.Error(), http.StatusBadRequest)
		return
	}

	userId, ok := vars["userId"]
	if !ok {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}
	intUserId, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Failed to convert userid string to int: "+err.Error(), http.StatusBadRequest)
		return
	}
	requestingUserId, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Failed to get user id from session token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// logic to check if the user trying to remove the member is the owner of the group
	isAuthorized, err := h.groupMemberRepo.IsUserGroupOwner(requestingUserId, intGroupId)
	if err != nil {
		http.Error(w, "Failed to check if user is group owner: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !isAuthorized {
		http.Error(w, "User requesting the removal is not the group owner", http.StatusUnauthorized)
		return
	}

	err = h.groupMemberRepo.RemoveMemberFromGroup(intGroupId, intUserId)
	if err != nil {
		http.Error(w, "Failed to remove member from group: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// ----------------- Group Membership/Invitation/Request Handlers -----------------

// RequestGroupMembershipHandler allows a user to request membership in a group.
// It creates a membership request in the database that can be approved or denied by the group's admin.
func (h *GroupMemberHandler) RequestGroupMembershipHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the group ID and user ID
	var request model.GroupInvitation
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Failed to get user id from session token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	request.JoinUserId = userID
	// Create the membership request in the database
	err = h.groupMemberRepo.CreateGroupRequest(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Notify the group admin
	err = notifyGroupAdmin(h.notificationRepo, request.GroupId, userID)
	if err != nil {
		http.Error(w, "Failed to notify group admin: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Notify the user
	err = notifyUserRequestSent(h.notificationRepo, userID)
	if err != nil {
		http.Error(w, "Failed to notify user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// -------- Notification Functions -------- //

// notifyGroupAdmin notifies the group admin that a user has requested to join the group.
func notifyGroupAdmin(notificationRepo *repository.NotificationRepository, groupID, userID int) error {
	message := fmt.Sprintf("User %d has requested to join the group %d.", userID, groupID)
	newNotification := model.Notification{
		UserId:  userID,
		Type:    "GroupRequest",
		Message: message,
		IsRead:  false,
	}
	_, err := notificationRepo.CreateNotification(newNotification)
	return err
}

// notifyUserRequestSent notifies the user that their request was sent and is pending.
func notifyUserRequestSent(notificationRepo *repository.NotificationRepository, userID int) error {
	message := "Your request to join the group has been sent and is pending approval."
	notification := model.Notification{
		UserId:  userID,
		Type:    "RequestSent",
		Message: message,
		IsRead:  false,
	}
	_, err := notificationRepo.CreateNotification(notification)
	return err
}

// ----------------------------------------------------------------------------------------------------------

// logic for setting the player as group member
func (h *GroupMemberHandler) ApproveGroupMembershipHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request URL to get the invitation ID
	vars := mux.Vars(r)
	id := vars["id"]

	//userID := r.Context().Value("AuthUserID").(int)
	// Update the status of the membership request to "approved"
	err := h.groupMemberRepo.AcceptGroupInvitationAndRequest(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// get groupId from the request invitation that was accepted
	groupInvitation, err := h.invitationRepo.GetGroupInvitationByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// this should add player to group
	err = h.groupMemberRepo.AddMemberToGroup(groupInvitation.GroupId, groupInvitation.JoinUserId)
	if err != nil {
		http.Error(w, "Error adding member to the group: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Notify the user that their request was approved
	err = notifyUserRequestApproved(h.notificationRepo, groupInvitation.JoinUserId, groupInvitation.GroupId)
	if err != nil {
		http.Error(w, "Failed to notify user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Mark the invitation as accepted in the database
	err = markGroupInvitationAs(h.invitationRepo, id)
	if err != nil {
		http.Error(w, "Failed to delete group invitation: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// -------- Notification Functions -------- //

// notifyUserRequestApproved notifies the user that their request was approved.
func notifyUserRequestApproved(notificationRepo *repository.NotificationRepository, userID, groupID int) error {
	message := fmt.Sprintf("Your request to join the group %d has been approved.", groupID)
	newNotification := model.Notification{
		UserId:  userID,
		Type:    "RequestApproved",
		Message: message,
		IsRead:  false,
	}
	_, err := notificationRepo.CreateNotification(newNotification)
	return err
}

// deleteGroupInvitation deletes a group invitation from the database.
func markGroupInvitationAs(invitationRepo *repository.InvitationRepository, id string) error {
	invitationID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return invitationRepo.DeleteGroupInvitation(invitationID)
}

// ----------------------------------------------------------------------------------------------------------

// Allows user to decline membership.
func (h *GroupMemberHandler) DeclineGroupMembershipHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request URL to get the invitation ID
	vars := mux.Vars(r)
	id := vars["id"]

	// Retrieve the user ID from the session token
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Failed to get user id from session token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Update the status of the membership request to "declined"
	err = h.invitationRepo.DeclineGroupInvitation(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Notify the user that their request was declined
	err = notifyUserDecline(h.notificationRepo, userID, "Your group membership request was declined.")
	if err != nil {
		http.Error(w, "Failed to notify user about request decline: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Mark invitation as
	err = markGroupInvitationAs(h.invitationRepo, id)
	if err != nil {
		// Handle the error if deleting the request fails.
		http.Error(w, "Failed to delete request from the database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// -------- Notification Functions -------- //

// Function to notify the user about the declined request.
func notifyUserDecline(notificationRepo *repository.NotificationRepository, userID int, message string) error {
	// Create a Notification object
	newNotification := model.Notification{
		UserId:  userID,
		Type:    "decline", // You can customize the type based on your needs
		Message: message,
		IsRead:  false, // Assuming the notification is initially unread
	}

	// Add the notification to the database
	_, err := notificationRepo.CreateNotification(newNotification)
	return err
}

// ----------------------------------------------------------------------------------------------------------

// InviteMemberHandler sends an invitation to a user to join a group.
// It creates an invitation in the database that can be accepted or declined by the user.
func (h *GroupMemberHandler) InviteGroupMemberHandler(w http.ResponseWriter, r *http.Request) {
	var newInvitation model.GroupInvitation
	err := json.NewDecoder(r.Body).Decode(&newInvitation)
	if err != nil {
		http.Error(w, "Failed to decode request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Failed to get user id from session token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	newInvitation.InviteUserId = userID
	err = h.invitationRepo.CreateGroupInvitation(newInvitation)
	if err != nil {
		http.Error(w, "Failed to create invitation: "+err.Error(), http.StatusInternalServerError)
		return
	}

	groupID, err := h.groupRepo.GetGroupByID(newInvitation.GroupId)
	if err != nil {
		http.Error(w, "Failed to get group id from session token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	newInvitation.GroupId = groupID.Id

	// Notify the user that they have been invited to join a group
	err = notifyUserInvitation(h.notificationRepo, newInvitation.InviteUserId, newInvitation.GroupId, "You have been invited to join a group.")
	if err != nil {
		http.Error(w, "Failed to notify user about the invitation: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// -------- Notification Functions -------- //

// notifyUserInvitation notifies the user about the group invitation.
func notifyUserInvitation(notificationRepo *repository.NotificationRepository, userID, groupID int, message string) error {
	// Create a new notification for the user
	newNotification := model.Notification{
		UserId:  userID,
		GroupId: groupID,
		Type:    "group_invitation",
		Message: message,
		IsRead:  false,
	}

	// Add the notification to the database
	_, err := notificationRepo.CreateNotification(newNotification)
	return err
}

// ----------------------------------------------------------------------------------------------------------

// AcceptGroupInvitationHandler allows a user to accept an invitation to join a group.
func (h *GroupMemberHandler) AcceptGroupInvitationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := h.groupMemberRepo.AcceptGroupInvitationAndRequest(id)
	if err != nil {
		http.Error(w, "Failed to accept invitation: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// get groupId from the request invitation that was accepted
	groupInvitation, err := h.invitationRepo.GetGroupInvitationByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//  add player to group
	err = h.groupMemberRepo.AddMemberToGroup(groupInvitation.GroupId, groupInvitation.JoinUserId)
	if err != nil {
		http.Error(w, "Error adding member to the group: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// FUTURE TODO: this notify the group list of the new member

	// Notify the group list of the new member.
	err = notifyGroupOfNewMember(h.groupRepo, h.notificationRepo, h.groupMemberRepo, groupInvitation.GroupId, groupInvitation.JoinUserId)
	if err != nil {
		// Handle the error if notifying the group fails.
		http.Error(w, "Failed to notify group about new member: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// -------- Notification Functions -------- //

// notifyGroupOfNewMember notifies the group members about the new member.
func notifyGroupOfNewMember(groupRepo *repository.GroupRepository, notificationRepo *repository.NotificationRepository, groupMemberRepo *repository.GroupMemberRepository, groupID, JoinUserId int) error {
	// Get the details of the group.
	group, err := groupRepo.GetGroupByID(groupID)
	if err != nil {
		return err
	}

	// Construct a notification message.
	message := fmt.Sprintf("A new member has joined the group '%s'.", group.Title)

	// Get the list of group members.
	members, err := groupMemberRepo.GetGroupMembers(groupID)
	if err != nil {
		return err
	}

	// Create notifications for each group member.
	for _, member := range members {
		// Skip notifying the new member.
		if member.UserId == JoinUserId {
			continue
		}

		// Create a new notification.
		newNotification := model.Notification{
			UserId:    member.UserId,
			Type:      "new_group_member",
			Message:   message,
			IsRead:    false,
			CreatedAt: time.Now(),
		}

		// Add the notification to the database.
		_, err := notificationRepo.CreateNotification(newNotification)
		if err != nil {
			return err
		}
	}

	return nil
}

// ----------------------------------------------------------------------------------------------------------

// DeclineGroupInvitationHandler allows a user to decline an invitation to join a group.
func (h *GroupMemberHandler) DeclineGroupInvitationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := h.invitationRepo.DeclineGroupInvitation(id)
	if err != nil {
		http.Error(w, "Failed to decline invitation: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// You can call a notification function here to notify the sender about the decline.
	err = notifyInvitationDecline(h.invitationRepo, h.notificationRepo, id)
	if err != nil {
		// Handle the error if notifying the sender fails.
		http.Error(w, "Failed to notify invitation sender about decline: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// This deletes the invitation from the database
	err = markGroupInvitationAs(h.invitationRepo, id)
	if err != nil {
		// Handle the error if deleting the invitation fails.
		http.Error(w, "Failed to delete invitation from the database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// -------- Notification Functions -------- //

// NotifyInvitationDecline notifies the sender about the declined invitation.
func notifyInvitationDecline(invitationRepo *repository.InvitationRepository, notificationRepo *repository.NotificationRepository, id string) error {
	// Get the invitation details from the repository based on the ID.
	invitation, err := invitationRepo.GetGroupInvitationByID(id)
	if err != nil {
		return err
	}

	// Construct a notification message.
	message := fmt.Sprintf("Your invitation to join the group %d has been declined by the user %d.", invitation.GroupId, invitation.InviteUserId)

	// Create a new notification.
	newNotification := model.Notification{
		UserId:    invitation.InviteUserId, // Group owner's user ID
		GroupId:   invitation.GroupId,
		Type:      "invitation_declined",
		Message:   message,
		IsRead:    false,
		CreatedAt: time.Now(),
	}

	// Add the notification to the database.
	_, err = notificationRepo.CreateNotification(newNotification)
	if err != nil {
		return err
	}

	return nil
}

// ----------------------------------------------------------------------------------------------------------

// GetGroupInvitationByIDHandler gets an invitation by ID for the user.
func (h *GroupMemberHandler) GetGroupInvitationByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the cookie.
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error extracting user ID from session token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	// Check if the user has the permission to view the invitation.
	invitation, err := h.invitationRepo.GetGroupInvitationByID(id)
	if err != nil {
		http.Error(w, "Failed to get invitation: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if invitation.JoinUserId != userID {
		http.Error(w, "User does not have permission to view this invitation", http.StatusUnauthorized)
		return
	}

	// Encode the invitation in the response.
	json.NewEncoder(w).Encode(invitation)
}

// GetAllGroupInvitationsHandler gets all pending invitations for the user.
func (h *GroupMemberHandler) GetAllGroupInvitationsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the cookie.
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error extracting user ID from session token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Get all pending group invitations for the user.
	invitations, err := h.invitationRepo.GetPendingGroupInvitationsForUser(userID)
	if err != nil {
		http.Error(w, "Failed to get group invitations: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the invitations in the response.
	json.NewEncoder(w).Encode(invitations)
}

// ----------------------------------------------------------------------------------------------
