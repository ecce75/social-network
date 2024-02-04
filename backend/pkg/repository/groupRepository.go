// Package repository provides the implementation of the repository layer for the social network backend.
// It includes functions to interact with the database for managing groups and invitations.
package repository

import (
	"backend/pkg/model"
	"database/sql"
	"time"
)

// GroupRepository is a repository for managing groups in the database.
type GroupRepository struct {
	db *sql.DB
}

// InvitationRepository is a repository for managing invitations in the database.
type InvitationRepository struct {
	db *sql.DB
}

// NewInvitationRepository creates a new instance of InvitationRepository.
func NewInvitationRepository(db *sql.DB) *InvitationRepository {
	return &InvitationRepository{db: db}
}

// NewGroupRepository creates a new instance of GroupRepository.
func NewGroupRepository(db *sql.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

// GetAllGroups retrieves all groups from the database.
// It returns a slice of Group objects and an error if any.
func (r *GroupRepository) GetAllGroups() ([]model.Group, error) {
	// SQL query to select all groups
	query := `SELECT * FROM groups`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []model.Group
	for rows.Next() {
		var group model.Group
		if err := rows.Scan(&group.Id, &group.Title, &group.Description, &group.CreatedAt); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return groups, nil
}

// CreateGroup creates a new group in the database.
// It returns the ID of the newly created group and an error if any.
// TODO: review this function - it may also return the new group instead of just the id
func (r *GroupRepository) CreateGroup(group model.Group) (int64, error) {
	query := `INSERT INTO groups (creator_id, title, description) VALUES (?, ?, ?)`
	result, err := r.db.Exec(query, group.CreatorId, group.Title, group.Description)
	if err != nil {
		return 0, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertID, nil
}

// GetGroupByID retrieves a group by ID from the database.
// It returns the Group object and an error if any.
func (r *GroupRepository) GetGroupByID(id int) (model.Group, error) {
	query := `SELECT * FROM groups WHERE id = ?`
	row := r.db.QueryRow(query, id)
	var group model.Group
	err := row.Scan(&group.Id, &group.CreatorId, &group.Title, &group.Description, &group.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Group{}, nil
		}
		return model.Group{}, err
	}
	return group, nil
}

// UpdateGroup updates a group in the database.
// It returns an error if any.
func (r *GroupRepository) UpdateGroup(group model.Group) error {
	query := `UPDATE groups SET creator_id = ?, title = ?, description = ?, updated_at = ? WHERE id = ?`
	_, err := r.db.Exec(query, group.CreatorId, group.Title, group.Description, time.Now(), group.Id)
	return err
}

// DeleteGroup deletes a group from the database.
// It returns an error if any.
func (r *GroupRepository) DeleteGroup(id int) error {
	query := `DELETE FROM groups WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

// AddMemberToGroup adds a member to a group in the database.
// It returns an error if any.
func (r *GroupRepository) AddMemberToGroup(groupId, userId int) error {
	query := `INSERT INTO group_members (group_id, user_id) VALUES (?, ?)`
	_, err := r.db.Exec(query, groupId, userId)
	return err
}

// RemoveMemberFromGroup removes a member from a group in the database.
// It returns an error if any.
func (r *GroupRepository) RemoveMemberFromGroup(groupId, userId int) error {
	query := `DELETE FROM group_members WHERE group_id = ? AND user_id = ?`
	_, err := r.db.Exec(query, groupId, userId)
	return err
}

// CreateGroupInvitation creates a new invitation in the database.
// It returns an error if any.
func (r *InvitationRepository) CreateGroupInvitation(invitation model.GroupInvitation) error {
	query := `INSERT INTO group_invitations (group_id, user_id, status) VALUES (?, ?, ?)`
	_, err := r.db.Exec(query, invitation.GroupId, invitation.UserId, "pending")
	return err
}

// UpdateGroupInvitation updates the status of an invitation in the database.
// It returns an error if any.
func (r *InvitationRepository) UpdateGroupInvitation(invitation model.GroupInvitation) error {
	query := `UPDATE group_invitations SET status = ? WHERE id = ?`
	_, err := r.db.Exec(query, invitation.Status, invitation.Id)
	return err
}

// DeleteGroupInvitation deletes an invitation from the database.
// It returns an error if any.
func (r *InvitationRepository) DeleteGroupInvitation(id int) error {
	query := `DELETE FROM group_invitations WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

// GetAllGroupInvitations retrieves all group invitations from the database.
// It returns a slice of GroupInvitation objects and an error if any.
func (r *InvitationRepository) GetAllGroupInvitations() ([]model.GroupInvitation, error) {
	// SQL query to select all group invitations
	query := `SELECT * FROM group_invitations`
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

// GetGroupInvitationByID retrieves an invitation by ID from the database.
// It returns the GroupInvitation object and an error if any.
func (r *InvitationRepository) GetGroupInvitationByID(id string) (model.GroupInvitation, error) {
	query := `SELECT * FROM group_invitations WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var invitation model.GroupInvitation
	if err := row.Scan(&invitation.Id, &invitation.GroupId, &invitation.UserId, &invitation.Status); err != nil {
		return model.GroupInvitation{}, err
	}
	return invitation, nil
}

// AcceptGroupInvitation updates the status of an invitation to "accepted" in the database.
// It returns an error if any.
func (r *InvitationRepository) AcceptGroupInvitation(id string) error {
	query := `UPDATE group_invitations SET status = 'accepted' WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

// DeclineGroupInvitation updates the status of an invitation to "declined" in the database.
// It returns an error if any.
func (r *InvitationRepository) DeclineGroupInvitation(id string) error {
	query := `UPDATE group_invitations SET status = 'declined' WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}