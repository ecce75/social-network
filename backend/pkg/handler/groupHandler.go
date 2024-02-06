package handler

import (
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

type InvitationHandler struct {
    repo *repository.InvitationRepository
}

func NewGroupHandler(repo *repository.GroupRepository) *GroupHandler {
    return &GroupHandler{repo: repo}
}

func NewInvitationHandler(repo *repository.InvitationRepository) *InvitationHandler {
    return &InvitationHandler{repo: repo}
}

// Group Handlers
func (h *GroupHandler) GetAllGroupsHandler(w http.ResponseWriter, r *http.Request) {
    // logic for getting all groups
    groups, err := h.repo.GetAllGroups()
    if err != nil {
        http.Error(w, "Failed to get groups: "+err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(groups)
}

func (h *GroupHandler) CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
    // logic for creating a group
    var newGroup model.Group
    err := json.NewDecoder(r.Body).Decode(&newGroup)
    if err != nil {
        http.Error(w, "Failed to decode request body: "+err.Error(), http.StatusBadRequest)
        return
    }
    // TODO: check if group with title already exists IN FRONTEND
    userID := r.Context().Value("AuthUserID").(int)

    newGroup.CreatorId = userID
    // creating the group in db
    _, err = h.repo.CreateGroup(newGroup)
    if err != nil {
        http.Error(w, "Failed to create group: "+err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)

}

func (h *GroupHandler) GetGroupByIDHandler(w http.ResponseWriter, r *http.Request) {
    // logic for getting a group by ID
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid group ID", http.StatusBadRequest)
        return
    }
    
    group, err := h.repo.GetGroupByID(id)
    if err != nil {
        http.Error(w, "Failed to get group: "+err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(group)
}

func (h *GroupHandler) EditGroupHandler(w http.ResponseWriter, r *http.Request) {
    // logic for editing a group
    var updatedGroup model.Group
    err := json.NewDecoder(r.Body).Decode(&updatedGroup)
    if err != nil {
        http.Error(w, "Failed to decode request body: "+err.Error(), http.StatusBadRequest)
        return
    }
    
    err = h.repo.UpdateGroup(updatedGroup)
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
    userID := r.Context().Value("AuthUserID").(int)

    // logic for deleting a group
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid group ID", http.StatusBadRequest)
        return
    }
    
    group, err := h.repo.GetGroupByID(id)
    if err != nil {
        http.Error(w, "Failed to get group: "+err.Error(), http.StatusInternalServerError)
        return
    }
    if group.CreatorId != userID {
        http.Error(w, "User not authorized to delete this group", http.StatusUnauthorized)
        return
    }
    err = h.repo.DeleteGroup(id)
    if err != nil {
        http.Error(w, "Failed to delete group: "+err.Error(), http.StatusInternalServerError)
        return
    }
    // Successful response
	response := map[string]string{
		"message": "Group deleted successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}