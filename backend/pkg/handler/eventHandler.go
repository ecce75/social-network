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
	eventRepo   *repository.EventRepository
	sessionRepo *repository.SessionRepository
}

func NewEventHandler(eventRepo *repository.EventRepository, sessionRepo *repository.SessionRepository) *EventHandler {
	return &EventHandler{eventRepo: eventRepo, sessionRepo: sessionRepo}
}

// Event Handlers
func (h *EventHandler) GetAllEventsHandler(w http.ResponseWriter, r *http.Request) {
	events, err := h.eventRepo.GetAllEvents()
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
	newEvent.CreatorId = userID
	// creating the event in db
	_, err = h.eventRepo.CreateEvent(newEvent)
	if err != nil {
		http.Error(w, "Failed to create group: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
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
