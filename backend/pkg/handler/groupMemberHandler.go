package handler

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"backend/util"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type GroupMemberHandler struct {
	groupMemberRepo     *repository.GroupMemberRepository
	invitationRepo      *repository.InvitationRepository
	sessionRepo         *repository.SessionRepository
	notificationHandler *NotificationHandler
	groupRepo           *repository.GroupRepository
	userRepo            *repository.UserRepository
}

func NewGroupMemberHandler(groupMemberRepo *repository.GroupMemberRepository, invitationRepo *repository.InvitationRepository, sessionRepo *repository.SessionRepository, notificationHandler *NotificationHandler, groupRepo *repository.GroupRepository, userRepo *repository.UserRepository) *GroupMemberHandler {
	return &GroupMemberHandler{groupMemberRepo: groupMemberRepo, invitationRepo: invitationRepo, sessionRepo: sessionRepo, notificationHandler: notificationHandler, groupRepo: groupRepo, userRepo: userRepo}
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
	var request model.GroupInvitation
	vars := mux.Vars(r)
	groupID, err := strconv.Atoi(vars["groupId"])
	if err != nil {
		http.Error(w, "Missing group ID", http.StatusBadRequest)
		return
	}

	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Failed to get user id from session token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	request.JoinUserId = userID
	request.GroupId = groupID

	// Create the membership request in the database
	err = h.groupMemberRepo.CreateGroupRequest(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Notify the group admin
	err = h.notificationHandler.NotifyGroupAdmin(request.GroupId, userID)
	if err != nil {
		http.Error(w, "Failed to notify group admin: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// ----------------------------------------------------------------------------------------------------------

// logic for setting the player as group member
func (h *GroupMemberHandler) ApproveGroupMembershipHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request URL to get the invitation ID
	vars := mux.Vars(r)
	groupID, _ := strconv.Atoi(vars["groupId"])
	// Retrieve the user ID from the session token
	userID, _ := strconv.Atoi(vars["userId"])
	// Update the status of the membership request to "approved"
	err := h.groupMemberRepo.AcceptGroupInvitationAndRequest(userID, groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// get groupId from the request invitation that was accepted
	groupInvitation, err := h.invitationRepo.GetGroupInvitationByID(userID, groupID)
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
	err = h.notificationHandler.NotifyUserRequestApproved(groupInvitation.JoinUserId, groupInvitation.GroupId)
	if err != nil {
		http.Error(w, "Failed to notify user about request approval: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// ----------------------------------------------------------------------------------------------------------

// Allows user to decline membership.
func (h *GroupMemberHandler) DeclineGroupMembershipHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request URL to get the invitation ID
	vars := mux.Vars(r)
	groupID, _ := strconv.Atoi(vars["groupId"])

	// Retrieve the user ID from URL
	userID, _ := strconv.Atoi(vars["userId"])
	// Update the status of the membership request to "declined"
	err := h.invitationRepo.DeclineGroupInvitation(userID, groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get groupId from the request invitation that was accepted
	groupInvitation, err := h.invitationRepo.GetGroupInvitationByID(userID, groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Notify the user that their request was declined
	err = h.notificationHandler.NotifyUserDecline(groupInvitation.JoinUserId, groupInvitation.GroupId)
	if err != nil {
		http.Error(w, "Failed to notify user about request decline: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// ----------------------------------------------------------------------------------------------------------

// InviteMemberHandler sends an invitation to a user to join a group.
// It creates an invitation in the database that can be accepted or declined by the user.
func (h *GroupMemberHandler) InviteGroupMemberHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID, _ := strconv.Atoi(vars["groupId"])

	// Retrieve the user ID from URL
	userID, _ := strconv.Atoi(vars["userId"])

	inviteUserID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Failed to get user id from session token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//
	newInvitation := model.GroupInvitation{
		GroupId:      groupID,
		JoinUserId:   userID,
		InviteUserId: inviteUserID,
		Status:       "pending",
	}
	err = h.invitationRepo.CreateGroupInvitation(newInvitation)
	if err != nil {
		http.Error(w, "Failed to create invitation: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// Notify the user that they have been invited to join a group
	err = h.notificationHandler.NotifyUserInvitation(userID, groupID)
	if err != nil {
		http.Error(w, "Failed to notify user about the invitation: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// ----------------------------------------------------------------------------------------------------------

// AcceptGroupInvitationHandler allows a user to accept an invitation to join a group.
func (h *GroupMemberHandler) AcceptGroupInvitationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID, _ := strconv.Atoi(vars["groupId"])

	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Failed to get user id from session token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.groupMemberRepo.AcceptGroupInvitationAndRequest(userID, groupID)
	if err != nil {
		http.Error(w, "Failed to accept invitation: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// get groupId from the request invitation that was accepted
	groupInvitation, err := h.invitationRepo.GetGroupInvitationByID(userID, groupID)
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
	groupName, _ := h.groupRepo.GetGroupTitleByID(groupID)

	message := "You joined the group " + groupName
	err = h.notificationHandler.EditGroupRequestNotification(userID, groupID, message)
	if err != nil {
		http.Error(w, "Error sending notification "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Notify the group list of the new member.
	err = h.notificationHandler.NotifyGroupOfNewMember(groupInvitation.GroupId, groupInvitation.JoinUserId)
	if err != nil {
		// Handle the error if notifying the group fails.
		http.Error(w, "Failed to notify group about new member: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// ----------------------------------------------------------------------------------------------------------

// DeclineGroupInvitationHandler allows a user to decline an invitation to join a group.
func (h *GroupMemberHandler) DeclineGroupInvitationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID, _ := strconv.Atoi(vars["groupId"])

	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Failed to get user id from session token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.invitationRepo.DeclineGroupInvitation(userID, groupID)
	if err != nil {
		http.Error(w, "Failed to decline invitation: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// You can call a notification function here to notify the sender about the decline.
	err = h.notificationHandler.NotifyInvitationDecline(userID, groupID)
	if err != nil {
		// Handle the error if notifying the sender fails.
		http.Error(w, "Failed to notify invitation sender about decline: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
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
	groupID, _ := strconv.Atoi(vars["groupId"])

	// Check if the user has the permission to view the invitation.
	invitation, err := h.invitationRepo.GetGroupInvitationByID(userID, groupID)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"message": "No group invitation found"})
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

func (h *GroupMemberHandler) GetAllGroupRequestsHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["groupId"]
	groupID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid group ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Extract the user ID from the cookie.
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error extracting user ID from session token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// check if user group owner
	isAuthorized, err := h.groupMemberRepo.IsUserGroupOwner(userID, groupID)
	if err != nil {
		http.Error(w, "Failed to check if user is group owner: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if isAuthorized {
		// Get all pending group requests for the user.
		requests, err := h.invitationRepo.GetPendingGroupRequestsForOwner(groupID, userID)
		for i, request := range requests {
			requests[i].Username, requests[i].ImageURL, _ = h.userRepo.GetUsernameAndAvatarByID(request.JoinUserId)

		}
		if err != nil {
			http.Error(w, "Failed to get group requests: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(requests)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "No pending group requests found"})
	}
}

func (h *GroupMemberHandler) GetAllNonMembersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID, err := strconv.Atoi(vars["groupId"])
	if err != nil {
		http.Error(w, "Invalid group ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Extract the user ID from the cookie.
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error extracting user ID from session token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	nonMembers, err := h.groupMemberRepo.GetNonMembers(groupID, userID)
	if err != nil {
		http.Error(w, "Failed to get non members: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nonMembers)
}

func (h *GroupMemberHandler) GetAllMembersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID, err := strconv.Atoi(vars["groupId"])
	if err != nil {
		http.Error(w, "Invalid group ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	members, err := h.groupMemberRepo.GetMembers(groupID)
	if err != nil {
		http.Error(w, "Failed to get members: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}

// ----------------------------------------------------------------------------------------------
