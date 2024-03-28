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

type EventHandler struct {
	eventRepo           *repository.EventRepository
	groupMemberRepo     *repository.GroupMemberRepository
	sessionRepo         *repository.SessionRepository
	userRepo            *repository.UserRepository
	notificationHandler *NotificationHandler
}

func NewEventHandler(eventRepo *repository.EventRepository, sessionRepo *repository.SessionRepository, groupMemberRepo *repository.GroupMemberRepository, userRepo *repository.UserRepository, notificationHandler *NotificationHandler) *EventHandler {
	return &EventHandler{eventRepo: eventRepo, sessionRepo: sessionRepo, groupMemberRepo: groupMemberRepo, userRepo: userRepo, notificationHandler: notificationHandler}
}

// Event Handlers
func (h *EventHandler) GetAllGroupEventsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID, _ := strconv.Atoi(vars["groupId"])
	events, err := h.eventRepo.GetAllGroupEvents(groupID)
	if err != nil {
		http.Error(w, "Failed to get events: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func (h *EventHandler) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var newEvent model.Event
	err := json.NewDecoder(r.Body).Decode(&newEvent)
	if err != nil {
		http.Error(w, "Failed to decode request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	// check if event with title already exists IN FRONTEND
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error confirming authentication: "+err.Error(), http.StatusInternalServerError)
		return
	}
	isMember, err := h.groupMemberRepo.IsUserGroupMember(userID, newEvent.GroupId)
	if !isMember {
		http.Error(w, "User not authorized to create event in this group", http.StatusUnauthorized)
		return
	}
	if err != nil {
		http.Error(w, "Failed to check if user is group member: "+err.Error(), http.StatusInternalServerError)
		return
	}
	newEvent.CreatorId = userID
	// creating the event in db
	eventID, err := h.eventRepo.CreateEvent(newEvent)
	if err != nil {
		http.Error(w, "Failed to create event: "+err.Error(), http.StatusInternalServerError)
		return
	}
	newEvent.Id = int(eventID)
	// notify group members of new event
	err = h.notificationHandler.NotifyGroupOfEvent(newEvent.GroupId, int(eventID))
	if err != nil {
		http.Error(w, "Failed to notify group members of new event: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newEvent)
}

func (h *EventHandler) GetEventByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	event, err := h.eventRepo.GetEventByID(id)
	if err != nil {
		http.Error(w, "Failed to get event: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

func (h *EventHandler) EditEventHandler(w http.ResponseWriter, r *http.Request) {
	var updatedEvent model.Event
	err := json.NewDecoder(r.Body).Decode(&updatedEvent)
	if err != nil {
		http.Error(w, "Failed to decode request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	event, err := h.eventRepo.GetEventByID(updatedEvent.Id)
	if err != nil {
		http.Error(w, "Failed to get group: "+err.Error(), http.StatusInternalServerError)
		return
	}
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error confirming authentication: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if event.CreatorId != userID {
		http.Error(w, "User not authorized to edit this event", http.StatusUnauthorized)
		return
	}
	err = h.eventRepo.EditEvent(updatedEvent)
	if err != nil {
		http.Error(w, "Failed to update event: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedEvent)
}

func (h *EventHandler) DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error confirming authentication: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// logic for deleting an event
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}

	event, err := h.eventRepo.GetEventByID(id)
	if err != nil {
		http.Error(w, "Failed to get group: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if event.CreatorId != userID {
		http.Error(w, "User not authorized to delete this group", http.StatusUnauthorized)
		return
	}
	err = h.eventRepo.DeleteEvent(id)
	if err != nil {
		http.Error(w, "Failed to delete event: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// Successful response
	response := map[string]string{
		"message": "Event deleted successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *EventHandler) GetEventsByGroupIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupIDStr, ok := vars["id"]
	intGroupID, err := strconv.Atoi(groupIDStr)
	if err != nil {
		http.Error(w, "Failed to parse group ID: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "Group ID is missing in parameters", http.StatusBadRequest)
		return
	}

	events, err := h.eventRepo.GetEventsByGroupID(intGroupID)
	if err != nil {
		http.Error(w, "Failed to retrieve events: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

// AddOrUpdateAttendanceHandler handles the HTTP POST request to add or update attendance status for a specific event and user.
func (h *EventHandler) AddOrUpdateAttendanceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventID, _ := strconv.Atoi(vars["eventId"])
	statusInt, _ := strconv.Atoi(vars["status"])

	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error confirming authentication: "+err.Error(), http.StatusInternalServerError)
		return
	}
	status := map[bool]string{true: "going", false: "not going"}[statusInt == 1]
	err = h.eventRepo.AddOrUpdateAttendance(eventID, userID, status)
	if err != nil {
		http.Error(w, "Failed to add or update attendance: "+err.Error(), http.StatusInternalServerError)
		return
	}
	username, avatar, _ := h.userRepo.GetUsernameAndAvatarByID(userID)
	response := map[string]any{
		"id":         userID,
		"username":   username,
		"avatar_url": avatar}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetAttendanceByEventIDHandler handles the HTTP GET request to retrieve attendance records for a specific event.
func (h *EventHandler) GetAttendanceByEventIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventIDStr, ok := vars["eventId"]
	if !ok {
		http.Error(w, "Event ID is missing in parameters", http.StatusBadRequest)
		return
	}

	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error confirming authentication: "+err.Error(), http.StatusInternalServerError)
		return
	}

	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		http.Error(w, "Failed to parse event ID: "+err.Error(), http.StatusInternalServerError)
		return
	}

	attendanceList, err := h.eventRepo.GetAttendanceByEventID(eventID)
	if err != nil {
		http.Error(w, "Failed to retrieve attendance: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for i := range attendanceList {
		if attendanceList[i].Id == userID {
			attendanceList[i].Status = "current_user"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(attendanceList)
}
