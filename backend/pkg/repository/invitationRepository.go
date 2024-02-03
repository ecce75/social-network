package repository

import (
    "backend/pkg/model"
    "database/sql"
)

type InvitationRepository struct {
    db *sql.DB
}

func NewInvitationRepository(db *sql.DB) *InvitationRepository {
    return &InvitationRepository{db: db}
}

// CreateInvitation creates a new invitation in the database.
func (r *InvitationRepository) CreateGroupInvitation(invitation model.GroupInvitation) error {
    query := `INSERT INTO invitations (group_id, user_id, status) VALUES (?, ?, ?)`
    _, err := r.db.Exec(query, invitation.GroupId, invitation.UserId, "pending")
    return err
}

// UpdateInvitation updates the status of an invitation in the database.
func (r *InvitationRepository) UpdateGroupInvitation(invitation model.GroupInvitation) error {
    query := `UPDATE invitations SET status = ? WHERE id = ?`
    _, err := r.db.Exec(query, invitation.Status, invitation.Id)
    return err
}

// DeleteInvitation deletes an invitation from the database.
func (r *InvitationRepository) DeleteGroupInvitation(id int) error {
    query := `DELETE FROM invitations WHERE id = ?`
    _, err := r.db.Exec(query, id)
    return err
}

// GetAllInvitations gets all invitations from the database.
func (r *InvitationRepository) GetAllInvitations() ([]model.GroupInvitation, error) {
    query := `SELECT * FROM invitations`
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var invitations []model.GroupInvitation
    for rows.Next() {
        var invitation model.GroupInvitation
        if err := rows.Scan(&invitation.Id, &invitation.GroupId, &invitation.UserId, &invitation.Status); err != nil {
            return nil, err
        }
        invitations = append(invitations, invitation)
    }
    return invitations, nil
}

// GetInvitationByID gets an invitation by ID from the database.
func (r *InvitationRepository) GetInvitationByID(id string) (model.GroupInvitation, error) {
    query := `SELECT * FROM invitations WHERE id = ?`
    row := r.db.QueryRow(query, id)

    var invitation model.GroupInvitation
    if err := row.Scan(&invitation.Id, &invitation.GroupId, &invitation.UserId, &invitation.Status); err != nil {
        return model.GroupInvitation{}, err
    }
    return invitation, nil
}

// AcceptGroupInvitation updates the status of an invitation to "accepted" in the database.
func (r *InvitationRepository) AcceptGroupInvitation(id string) error {
    query := `UPDATE invitations SET status = 'accepted' WHERE id = ?`
    _, err := r.db.Exec(query, id)
    return err
}

// DeclineGroupInvitation updates the status of an invitation to "declined" in the database.
func (r *InvitationRepository) DeclineGroupInvitation(id string) error {
    query := `UPDATE invitations SET status = 'declined' WHERE id = ?`
    _, err := r.db.Exec(query, id)
    return err
}