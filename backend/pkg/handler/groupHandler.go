package handler

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"backend/util"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type GroupHandler struct {
	groupRepo        *repository.GroupRepository
	groupMemberRepo  *repository.GroupMemberRepository
	sessionRepo      *repository.SessionRepository
	notificationRepo *repository.NotificationRepository
}

func NewGroupHandler(groupRepo *repository.GroupRepository, sessionRepo *repository.SessionRepository, groupMemberRepo *repository.GroupMemberRepository, notificationRepo *repository.NotificationRepository) *GroupHandler {
	return &GroupHandler{groupRepo: groupRepo, sessionRepo: sessionRepo, groupMemberRepo: groupMemberRepo, notificationRepo: notificationRepo}
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
	var newGroup model.Group
	err := json.NewDecoder(r.Body).Decode(&newGroup)
	if err != nil {
		http.Error(w, "Failed to decode request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	// TODO: check if group with title already exists IN FRONTEND
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error confirming authentication: "+err.Error(), http.StatusInternalServerError)
		return
	}
	newGroup.CreatorId = userID
	// creating the group in db
	_, err = h.groupRepo.CreateGroup(newGroup)
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
	err = notifyGroupDeletion(h.groupMemberRepo, h.notificationRepo, id)
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

// -------- Notification Function -------- //

// notifyGroupDeletion notifies all group members about the group deletion.
func notifyGroupDeletion(groupMemberRepo *repository.GroupMemberRepository, notificationRepo *repository.NotificationRepository, groupID int) error {
	// Get the list of group members.
	members, err := groupMemberRepo.GetGroupMembers(groupID)
	if err != nil {
		return err
	}

	// Construct a notification message.
	message := "The group has been deleted."

	// Create notifications for each group member.
	for _, member := range members {
		// Create a new notification.
		newNotification := model.Notification{
			UserId:    member.UserId,
			Type:      "group_deletion",
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
