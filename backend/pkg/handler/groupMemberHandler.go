package handler

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type GroupMemberHandler struct {
	repo *repository.GroupMemberRepository
}

func NewGroupMemberHandler(repo *repository.GroupMemberRepository) *GroupMemberHandler {
	return &GroupMemberHandler{repo: repo}
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
    intGroupId, err := strconv.Atoi(groupId); if err != nil {
        http.Error(w, "Failed to convert groupid string to int: " + err.Error(), http.StatusBadRequest)
        return
    }

    userId, ok := vars["userId"]
    if !ok {
        http.Error(w, "Missing user ID", http.StatusBadRequest)
        return
    }
    intUserId, err := strconv.Atoi(userId); if err != nil {
        http.Error(w, "Failed to convert userid string to int: " + err.Error(), http.StatusBadRequest)
        return
    }
    requestingUserId := r.Context().Value("AuthUserID").(int)
    // logic to check if the user trying to remove the member is the owner of the group
    isAuthenticated, err := h.repo.IsUserGroupOwner(requestingUserId, intGroupId)
    if err != nil {
        http.Error(w, "Failed to check if user is group owner: "+err.Error(), http.StatusInternalServerError)
        return
    }
    if !isAuthenticated {
        http.Error(w, "User requesting the removal is not the group owner", http.StatusUnauthorized)
        return
    }

    err = h.repo.RemoveMemberFromGroup(intGroupId, intUserId)
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
    userID := r.Context().Value("AuthUserID").(int)
    request.JoinUserId = userID
    // Create the membership request in the database
    err = h.repo.CreateGroupRequest(request)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    // TODO: this should notify the group admin that a user has requested to join the group
    // TODO: this should notify the user that their request was sent and is pending

    w.WriteHeader(http.StatusCreated)
}

//  logic for setting the player as group member
func (h *GroupMemberHandler) ApproveGroupMembershipHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the request URL to get the invitation ID
    vars := mux.Vars(r)
    id := vars["id"]

    //userID := r.Context().Value("AuthUserID").(int)
    // Update the status of the membership request to "approved"
    err := h.repo.AcceptGroupInvitationAndRequest(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    // get groupId from the request invitation that was accepted
    groupInvitation, err := h.repo.GetGroupInvitationByID(id); if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    // this should add player to group
    err = h.repo.AddMemberToGroup(groupInvitation.GroupId, groupInvitation.JoinUserId)
    if err != nil {
        http.Error(w, "Error adding member to the group: " +err.Error(), http.StatusInternalServerError)
        return
    }
    // TODO: this should notify the user that their request was approved
    // TODO: this should delete the invitation from the database
    w.WriteHeader(http.StatusOK)
}

func (h *InvitationHandler) DeclineGroupMembershipHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the request URL to get the invitation ID
    vars := mux.Vars(r)
    id := vars["id"]

    // Update the status of the membership request to "declined"
    err := h.repo.DeclineGroupInvitation(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // TODO: this should notify the user that their request was declined
    // TODO: this should delete the request from the database

    w.WriteHeader(http.StatusOK)
}

// InviteMemberHandler sends an invitation to a user to join a group.
// It creates an invitation in the database that can be accepted or declined by the user.
func (h *InvitationHandler) InviteGroupMemberHandler(w http.ResponseWriter, r *http.Request) {
    var newInvitation model.GroupInvitation
    err := json.NewDecoder(r.Body).Decode(&newInvitation)
    if err != nil {
        http.Error(w, "Failed to decode request body: "+err.Error(), http.StatusBadRequest)
        return
    }
    err = h.repo.CreateGroupInvitation(newInvitation)
    if err != nil {
        http.Error(w, "Failed to create invitation: "+err.Error(), http.StatusInternalServerError)
        return
    }
    // TODO: this should notify the user that they have been invited to join a group
    w.WriteHeader(http.StatusCreated)
}

// AcceptGroupInvitationHandler allows a user to accept an invitation to join a group.
func (h *GroupMemberHandler) AcceptGroupInvitationHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    err := h.repo.AcceptGroupInvitationAndRequest(id)
    if err != nil {
        http.Error(w, "Failed to accept invitation: "+err.Error(), http.StatusInternalServerError)
        return
    }
    // get groupId from the request invitation that was accepted
    groupInvitation, err := h.repo.GetGroupInvitationByID(id); if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    //  add player to group
    err = h.repo.AddMemberToGroup(groupInvitation.GroupId, groupInvitation.JoinUserId)
    if err != nil {
        http.Error(w, "Error adding member to the group: " +err.Error(), http.StatusInternalServerError)
        return
    }
    // FUTURE TODO: this notify the group list of the new member
    w.WriteHeader(http.StatusOK)
}

// DeclineGroupInvitationHandler allows a user to decline an invitation to join a group.
func (h *InvitationHandler) DeclineGroupInvitationHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    err := h.repo.DeclineGroupInvitation(id)
    if err != nil {
        http.Error(w, "Failed to decline invitation: "+err.Error(), http.StatusInternalServerError)
        return
    }
    // TODO: this should notify the invitation sender that the invitation was declined
    // it also then should delete the invitation from the database
    w.WriteHeader(http.StatusOK)
}

// GetInvitationByIDHandler gets an invitation by ID
// TODO: refactor to use userid from cookie to get all group invitations for the user (pending)
func (h *InvitationHandler) GetGroupInvitationByIDHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    invitation, err := h.repo.GetGroupInvitationByID(id)
    if err != nil {
        http.Error(w, "Failed to get invitation: "+err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(invitation)
}

// GetAllInvitationsHandler gets all invitations (ok)
func (h *InvitationHandler) GetAllGroupInvitationsHandler(w http.ResponseWriter, r *http.Request) {
    invitations, err := h.repo.GetAllGroupInvitations()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(invitations)
}
