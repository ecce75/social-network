package handler

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type InvitationHandler struct {
    repo *repository.InvitationRepository
}

func NewInvitationHandler(repo *repository.InvitationRepository) *InvitationHandler {
    return &InvitationHandler{repo: repo}
}

// GetAllInvitationsHandler gets all invitations
func (h *InvitationHandler) GetAllInvitationsHandler(w http.ResponseWriter, r *http.Request) {
    invitations, err := h.repo.GetAllInvitations()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(invitations)
}

// CreateInvitationHandler creates a new invitation
func (h *InvitationHandler) CreateInvitationHandler(w http.ResponseWriter, r *http.Request) {
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
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newInvitation)
}

// GetInvitationByIDHandler gets an invitation by ID
func (h *InvitationHandler) GetInvitationByIDHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    invitation, err := h.repo.GetInvitationByID(id)
    if err != nil {
        http.Error(w, "Failed to get invitation: "+err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(invitation)
}

// AcceptGroupInvitationHandler allows a user to accept an invitation to join a group.
func (h *InvitationHandler) AcceptGroupInvitationHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    err := h.repo.AcceptGroupInvitation(id)
    if err != nil {
        http.Error(w, "Failed to accept invitation: "+err.Error(), http.StatusInternalServerError)
        return
    }
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
    w.WriteHeader(http.StatusOK)
}