package repository

import (
	"backend/pkg/model"
	"database/sql"
)

type GroupMemberRepository struct {
	db *sql.DB
}

func NewGroupMemberRepository(db *sql.DB) *GroupMemberRepository {
	return &GroupMemberRepository{db: db}
}

// InvitationRepository is a repository for managing invitations in the database.
type InvitationRepository struct {
	db *sql.DB
}

// NewInvitationRepository creates a new instance of InvitationRepository.
func NewInvitationRepository(db *sql.DB) *InvitationRepository {
	return &InvitationRepository{db: db}
}

// AddMemberToGroup adds a member to a group in the database.
// It returns an error if any.
func (r *GroupMemberRepository) AddMemberToGroup(groupId, userId int) error {
	query := `INSERT INTO group_members (group_id, user_id) VALUES (?, ?)`
	_, err := r.db.Exec(query, groupId, userId)
	return err
}

// RemoveMemberFromGroup removes a member from a group in the database.
// It returns an error if any.
func (r *GroupMemberRepository) RemoveMemberFromGroup(groupId, userId int) error {
	query := `DELETE FROM group_members WHERE group_id = ? AND user_id = ?`
	_, err := r.db.Exec(query, groupId, userId)
	return err
}

// CreateGroupInvitation creates a new invitation in the database.
// It returns an error if any.
// TODO: implement notifications for group invitations
func (r *InvitationRepository) CreateGroupInvitation(invitation model.GroupInvitation) error {
	query := `INSERT INTO group_invitations (group_id, user_id, status) VALUES (?, ?, ?)`
	_, err := r.db.Exec(query, invitation.GroupId, invitation.JoinUserId, &invitation.InviteUserId, "pending")
	return err
}

// DeleteGroupInvitation deletes an invitation from the database.
// It returns an error if any.
func (r *InvitationRepository) DeleteGroupInvitation(id int) error {
	query := `DELETE FROM group_invitations WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

// AcceptGroupInvitation updates the status of an invitation to "accepted" in the database.
// It returns an error if any.
func (r *GroupMemberRepository) AcceptGroupInvitationAndRequest(id string) error {
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

func (r *GroupMemberRepository) CreateGroupRequest(request model.GroupInvitation) error {
	query := `INSERT INTO group_invitations (group_id, join_user_id, invite_user_id, status) VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(query, request.GroupId, request.JoinUserId, request.InviteUserId, "pending")
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
		if err := rows.Scan(&invitation.Id, &invitation.GroupId, &invitation.JoinUserId, &invitation.InviteUserId, &invitation.Status); err != nil {
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
	if err := row.Scan(&invitation.Id, &invitation.GroupId, &invitation.JoinUserId, &invitation.InviteUserId, &invitation.Status); err != nil {
		return model.GroupInvitation{}, err
	}
	return invitation, nil
}

// GetGroupInvitationByID retrieves an invitation by ID from the database.
// It returns the GroupInvitation object and an error if any.
func (r *GroupMemberRepository) GetGroupInvitationByID(id string) (model.GroupInvitation, error) {
	query := `SELECT * FROM group_invitations WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var invitation model.GroupInvitation
	if err := row.Scan(&invitation.Id, &invitation.GroupId, &invitation.JoinUserId, &invitation.InviteUserId, &invitation.Status); err != nil {
		return model.GroupInvitation{}, err
	}
	return invitation, nil
}

func (r *GroupMemberRepository) IsUserGroupOwner(userId, groupId int) (bool, error) {
	query := `SELECT creator_id FROM groups WHERE id = ?`
	row := r.db.QueryRow(query, groupId)

	var creatorId int
	err := row.Scan(&creatorId)
	if err != nil {
		return false, err
	}
	return creatorId == userId, nil
}

func (r *GroupMemberRepository) IsUserGroupMember(userId, groupId int) (bool, error) {
	query := `SELECT user_id FROM group_members WHERE group_id = ? AND user_id = ?`
	row := r.db.QueryRow(query, groupId, userId)

	var memberId int
	err := row.Scan(&memberId)
	if err != nil {
		return false, err
	}
	return memberId == userId, nil
}