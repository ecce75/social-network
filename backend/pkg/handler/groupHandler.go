package handler

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"backend/util"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type GroupHandler struct {
	groupRepo           *repository.GroupRepository
	groupMemberRepo     *repository.GroupMemberRepository
	sessionRepo         *repository.SessionRepository
	notificationHandler *NotificationHandler
}

func NewGroupHandler(groupRepo *repository.GroupRepository, sessionRepo *repository.SessionRepository, groupMemberRepo *repository.GroupMemberRepository, notificationHandler *NotificationHandler) *GroupHandler {
	return &GroupHandler{groupRepo: groupRepo, sessionRepo: sessionRepo, groupMemberRepo: groupMemberRepo, notificationHandler: notificationHandler}
}

// Group Handlers
func (h *GroupHandler) GetAllGroupsHandler(w http.ResponseWriter, r *http.Request) {
	// logic for getting all groups
	groups, err := h.groupRepo.GetAllGroups()
	if err != nil {
		http.Error(w, "Failed to get groups: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
}

func (h *GroupHandler) CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	// logic for creating a group
	err1 := r.ParseMultipartForm(10 << 20) // Maximum memory 10MB, change this based on your requirements
	if err1 != nil {
		http.Error(w, "Error parsing form data: "+err1.Error(), http.StatusBadRequest)
		return
	}

	// TODO: check if group with title already exists IN FRONTEND
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error confirming authentication: "+err.Error(), http.StatusInternalServerError)
		return
	}
	var newGroup model.Group
	newGroup.Title = r.FormValue("title")
	newGroup.Description = r.FormValue("description")
	newGroup.CreatorId = userID

	// creating the group in db
	groupID, err := h.groupRepo.CreateGroup(newGroup)
	if err != nil {
		http.Error(w, "Failed to create group: "+err.Error(), http.StatusInternalServerError)
		return
	}
	util.ImageSave(w, r, strconv.Itoa(int(groupID)), "group")
	response := map[string]interface{}{
		"message": "Group created successfully",
		"id":      groupID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func (h *GroupHandler) GetGroupByIDHandler(w http.ResponseWriter, r *http.Request) {
	// logic for getting a group by ID
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}

	group, err := h.groupRepo.GetGroupByID(id)
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
	group, err := h.groupRepo.GetGroupByID(updatedGroup.Id)
	if err != nil {
		http.Error(w, "Failed to get group: "+err.Error(), http.StatusInternalServerError)
		return
	}
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error confirming authentication: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if group.CreatorId != userID {
		http.Error(w, "User not authorized to edit this group", http.StatusUnauthorized)
		return
	}
	err = h.groupRepo.UpdateGroup(updatedGroup)
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
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error confirming authentication: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// logic for deleting a group
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}

	group, err := h.groupRepo.GetGroupByID(id)
	if err != nil {
		http.Error(w, "Failed to get group: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if group.CreatorId != userID {
		http.Error(w, "User not authorized to delete this group", http.StatusUnauthorized)
		return
	}

	// Notify all group members that the group has been deleted
	err = h.notificationHandler.NotifyGroupDeletion(id)
	if err != nil {
		http.Error(w, "Failed to notify group members about group deletion: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Implement logging of the deletion or add a bool field "deleted"
	err = h.groupRepo.LogGroupDeletion(id)
	if err != nil {
		http.Error(w, "Failed to log group deletion: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.groupRepo.DeleteGroup(id)
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
