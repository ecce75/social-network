package handler

import (
	"backend/pkg/db/sqlite"
	"backend/pkg/model"
	"backend/pkg/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type GroupHandler struct {
    repo *repository.GroupRepository
}

func NewGroupHandler(repo *repository.GroupRepository) *GroupHandler {
    return &GroupHandler{repo: repo}
}

// Group Handlers
func (h *GroupHandler) GetAllGroupsHandler(w http.ResponseWriter, r *http.Request) {
    // check auth and get userid from cookie
    cookie, err := r.Cookie("session_token")
    if err != nil {
        if err == http.ErrNoCookie {
            // If the session cookie doesn't exist, set isAuthenticated to false
            http.Error(w, "User not authenticated", http.StatusUnauthorized)
            return
        } else {
            http.Error(w, "Error checking session token: " + err.Error(), http.StatusInternalServerError)
            return
        }
    }
    _, err = confirmAuthentication(cookie) //(userid, err :=) when later would highlight groups user is a part of
    if err != nil {
        http.Error(w, "Error confirming authentication: " + err.Error(), http.StatusUnauthorized)
        return
    }
    // TODO: Implement logic for getting all groups
    groups, err := repository.NewGroupRepository(sqlite.Dbase).GetAllGroups()
    if err != nil {
        http.Error(w, "Failed to get groups: "+err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(groups)
}

func (h *GroupHandler) CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement logic for creating a group
    var newGroup model.Group
    err := json.NewDecoder(r.Body).Decode(&newGroup)
    if err != nil {
        http.Error(w, "Failed to decode request body: "+err.Error(), http.StatusBadRequest)
        return
    }
    // check auth and get userid from cookie
    cookie, err := r.Cookie("session_token")
    if err != nil {
        if err == http.ErrNoCookie {
            // If the session cookie doesn't exist, set isAuthenticated to false
            http.Error(w, "User not authenticated", http.StatusUnauthorized)
            return
        } else {
            http.Error(w, "Error checking session token: " + err.Error(), http.StatusInternalServerError)
            return
        }
    }
    userid, err := confirmAuthentication(cookie)
    if err != nil {
        http.Error(w, "Error confirming authentication: " + err.Error(), http.StatusUnauthorized)
        return
    }
    // TODO: check if group with title already exists

    newGroup.CreatorId = userid
    groupRepo := repository.NewGroupRepository(sqlite.Dbase)
    createdGroup, err := groupRepo.CreateGroup(newGroup)
    if err != nil {
        http.Error(w, "Failed to create group: "+err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(createdGroup)

}

func (h *GroupHandler) GetGroupByIDHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement logic for getting a group by ID
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid group ID", http.StatusBadRequest)
        return
    }
    groupRepo := repository.NewGroupRepository(sqlite.Dbase)
    group, err := groupRepo.GetGroupByID(id)
    if err != nil {
        http.Error(w, "Failed to get group: "+err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(group)
}

func (h *GroupHandler) EditGroupHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement logic for editing a group
    var updatedGroup model.Group
    err := json.NewDecoder(r.Body).Decode(&updatedGroup)
    if err != nil {
        http.Error(w, "Failed to decode request body: "+err.Error(), http.StatusBadRequest)
        return
    }
    groupRepo := repository.NewGroupRepository(sqlite.Dbase)
    err = groupRepo.UpdateGroup(updatedGroup)
    if err != nil {
        http.Error(w, "Failed to update group: "+err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedGroup)
}

// DeleteGroupHandler handles the HTTP request for deleting a group.
// It checks the user's authentication, verifies their authorization to delete the group,
// and deletes the group from the repository if all conditions are met.
// If any errors occur during the process, appropriate HTTP error responses are returned.
// TODO: implement notification to all group members that the group has been deleted, and remove all group members;
// implement logging of the deletion or add bool field "deleted"
func (h *GroupHandler) DeleteGroupHandler(w http.ResponseWriter, r *http.Request) {
    // check auth and get userid from cookie
    cookie, err := r.Cookie("session_token")
    if err != nil {
        if err == http.ErrNoCookie {
            // If the session cookie doesn't exist, set isAuthenticated to false
            http.Error(w, "User not authenticated", http.StatusUnauthorized)
            return
        } else {
            http.Error(w, "Error checking session token: " + err.Error(), http.StatusInternalServerError)
            return
        }
    }
    userId, err := confirmAuthentication(cookie)
    if err != nil {
        http.Error(w, "Error confirming authentication: " + err.Error(), http.StatusUnauthorized)
        return
    }

    // TODO: Implement logic for deleting a group
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid group ID", http.StatusBadRequest)
        return
    }
    groupRepo := repository.NewGroupRepository(sqlite.Dbase)
    group, err := groupRepo.GetGroupByID(id)
    if err != nil {
        http.Error(w, "Failed to get group: "+err.Error(), http.StatusInternalServerError)
        return
    }
    if group.CreatorId != userId {
        http.Error(w, "User not authorized to delete this group", http.StatusUnauthorized)
        return
    }
    err = groupRepo.DeleteGroup(id)
    if err != nil {
        http.Error(w, "Failed to delete group: "+err.Error(), http.StatusInternalServerError)
        return
    }
    // Successful response
	response := map[string]string{
		"message": "Post deleted successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// AddMemberToGroup adds a user to a group. It takes two parameters: the ID of the group
// and the ID of the user. It inserts a new row into the group_members table in the database,
// which represents the user being a member of the group. If the operation is successful,
// it returns nil. If there is an error, it returns the error.
func (h *GroupHandler) AddMemberHandler(w http.ResponseWriter, r *http.Request) {
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

    err = h.repo.AddMemberToGroup(intGroupId, intUserId); if err != nil {
        http.Error(w, "Failed to add member to group: "+err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

// RemoveMemberFromGroup removes a user from a group. It takes two parameters: the ID of the group
// and the ID of the user. It deletes the row from the group_members table in the database that
// represents the user being a member of the group. If the operation is successful, it returns nil.
// If there is an error, it returns the error.
func (h *GroupHandler) RemoveMemberHandler(w http.ResponseWriter, r *http.Request) {
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

    err = h.repo.RemoveMemberFromGroup(intGroupId, intUserId)
    if err != nil {
        http.Error(w, "Failed to remove member from group: "+err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

// RequestMembershipHandler allows a user to request membership in a group.
// It creates a membership request in the database that can be approved or denied by the group's admin.
func (h *GroupHandler) RequestMembershipHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement logic for a user requesting membership in a group
}

// ApproveMembershipHandler allows the group's admin to approve a membership request.
// It changes the status of the membership request in the database to "approved".
func (h *GroupHandler) ApproveMembershipHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement logic for approving a membership request
}

// DeclineMembershipHandler allows the group's admin to decline a membership request.
// It changes the status of the membership request in the database to "declined".
func (h *GroupHandler) DeclineMembershipHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement logic for declining a membership request
}

// InviteMemberHandler sends an invitation to a user to join a group.
// It creates an invitation in the database that can be accepted or declined by the user.
func (h *GroupHandler) InviteMemberHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement logic for inviting a member to a group
}