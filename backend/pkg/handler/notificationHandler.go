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

// NotificationHandler handles HTTP requests related to notifications.
type NotificationHandler struct {
	notificationRepo *repository.NotificationRepository
	sessionRepo      *repository.SessionRepository
}

// NewNotificationHandler creates a new instance of NotificationHandler.
// It takes a NotificationRepository and a SessionRepository as parameters.
// Returns a pointer to the newly created NotificationHandler.
func NewNotificationHandler(notificationRepo *repository.NotificationRepository, sessionRepo *repository.SessionRepository) *NotificationHandler {
	return &NotificationHandler{notificationRepo: notificationRepo, sessionRepo: sessionRepo}
}

// GetAllNotificationsHandler retrieves all notifications and responds
func (h *NotificationHandler) GetAllNotificationsHandler(w http.ResponseWriter, r *http.Request) {
	notifications, err := h.notificationRepo.GetAllNotifications()
	if err != nil {
		http.Error(w, "Failed to get notifications: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifications)
}

// CreateNotificationHandler creates a new notification based on the request body and user session.
func (h *NotificationHandler) CreateNotificationHandler(w http.ResponseWriter, r *http.Request) {
	var newNotification model.Notification
	err := json.NewDecoder(r.Body).Decode(&newNotification)
	if err != nil {
		http.Error(w, "Failed to decode request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	// check if notification with title already exists IN FRONTEND
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error confirming authentication: "+err.Error(), http.StatusInternalServerError)
		return
	}
	newNotification.UserId = userID
	// creating the notification in db
	_, err = h.notificationRepo.CreateNotification(newNotification)
	if err != nil {
		http.Error(w, "Failed to create notification: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
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
