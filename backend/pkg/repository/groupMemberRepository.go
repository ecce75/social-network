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
func (r *InvitationRepository) CreateGroupInvitation(invitation model.GroupInvitation) error {
	query := `INSERT INTO group_invitations (group_id, join_user_id, invite_user_id, status) VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(query, invitation.GroupId, invitation.JoinUserId, invitation.InviteUserId, "pending")
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
func (r *GroupMemberRepository) AcceptGroupInvitationAndRequest(userID, groupID int) error {
	query := `UPDATE group_invitations SET status = 'accepted' WHERE join_user_id = ? AND group_id = ?`
	_, err := r.db.Exec(query, userID, groupID)
	return err
}

// DeclineGroupInvitation updates the status of an invitation to "declined" in the database.
// It returns an error if any.
func (r *InvitationRepository) DeclineGroupInvitation(userID, groupID int) error {
	query := `UPDATE group_invitations SET status = 'declined' WHERE join_user_id = ? AND group_id = ?`
	_, err := r.db.Exec(query, userID, groupID)
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
func (r *InvitationRepository) GetGroupInvitationByID(userID, groupID int) (model.GroupInvitation, error) {
	query := `SELECT * FROM group_invitations WHERE join_user_id = ? AND group_id = ?`
	row := r.db.QueryRow(query, userID, groupID)
	var invitation model.GroupInvitation
	if err := row.Scan(&invitation.Id, &invitation.GroupId, &invitation.JoinUserId, &invitation.InviteUserId, &invitation.Status, &invitation.CreatedAt); err != nil {
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

func (r *GroupMemberRepository) GetGroupAdminByID(groupId int) (int, error) {
	query := `SELECT creator_id FROM groups WHERE id = ?`
	row := r.db.QueryRow(query, groupId)

	var creatorId int
	err := row.Scan(&creatorId)
	if err != nil {
		return 0, err
	}
	return creatorId, nil
}

func (r *GroupMemberRepository) IsUserGroupMember(userId, groupId int) (bool, error) {
	query := `SELECT user_id FROM group_members WHERE group_id = ? AND user_id = ?`
	row := r.db.QueryRow(query, groupId, userId)

	var memberId int
	err := row.Scan(&memberId)
	if err == sql.ErrNoRows {
		// Check if the user is the creator of the group
		query = `SELECT creator_id FROM groups WHERE id = ? AND creator_id = ?`
		row = r.db.QueryRow(query, groupId, userId)

		var creatorId int
		err = row.Scan(&creatorId)
		if err == sql.ErrNoRows {
			return false, nil
		} else if err != nil {
			return false, err
		}
	} else if err != nil {
		return false, err
	}

	return true, nil
}

// GetGroupMembers retrieves the list of members for a given group ID.
func (r *GroupMemberRepository) GetGroupMembers(groupID int) ([]model.GroupMember, error) {
	query := `SELECT group_id, user_id, joined_at FROM group_members WHERE group_id = ?`
	rows, err := r.db.Query(query, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []model.GroupMember
	for rows.Next() {
		var member model.GroupMember
		err := rows.Scan(&member.GroupID, &member.UserID, &member.JoinedAt)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return members, nil
}

// GetPendingGroupInvitationsForUser retrieves all pending group invitations for the user.
func (r *InvitationRepository) GetPendingGroupInvitationsForUser(userID int) ([]model.GroupInvitation, error) {
	query := `SELECT * FROM group_invitations WHERE join_user_id = ? AND status = 'pending'`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invitations []model.GroupInvitation
	for rows.Next() {
		var invitation model.GroupInvitation
		err := rows.Scan(&invitation.Id, &invitation.GroupId, &invitation.JoinUserId, &invitation.InviteUserId, &invitation.Status, &invitation.CreatedAt)
		if err != nil {
			return nil, err
		}
		invitations = append(invitations, invitation)
	}

	return invitations, nil
}

// GetPendingGroupInvitationsForOwner retrieves all pending group invitations for the owner.
func (r *InvitationRepository) GetPendingGroupRequestsForOwner(groupID, userID int) ([]model.GroupInvitation, error) {
	query := `SELECT * FROM group_invitations WHERE group_id = ? AND (status = 'pending' OR status = 'declined') AND invite_user_id != ?`
	rows, err := r.db.Query(query, groupID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invitations []model.GroupInvitation
	for rows.Next() {
		var invitation model.GroupInvitation
		err := rows.Scan(&invitation.Id, &invitation.GroupId, &invitation.JoinUserId, &invitation.InviteUserId, &invitation.Status, &invitation.CreatedAt)
		if err != nil {
			return nil, err
		}
		invitations = append(invitations, invitation)
	}

	return invitations, nil
}

// RemoveGroupMembers removes all group members of a specific group.
func (r *GroupMemberRepository) RemoveGroupMembers(groupID int) error {
	query := `DELETE FROM group_members WHERE group_id = ?`
	_, err := r.db.Exec(query, groupID)
	return err
}

func (r *GroupMemberRepository) GetNonMembers(groupID, userID int) ([]model.UserList, error) {
	// Define the SQL query to select all users who are not members of the specified group
	// The query excludes the current user as well
	query := `
        SELECT id, username, avatar_url
        FROM users
        WHERE id != ?
        AND id NOT IN (
            SELECT user_id
            FROM group_members
            WHERE group_id = ?
        )
        AND id NOT IN (
        	SELECT join_user_id
        	FROM group_invitations
        WHERE group_id = ? AND status = 'pending' OR status = 'declined')
    `

	// Execute the query
	rows, err := r.db.Query(query, userID, groupID, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Slice to hold the users who are not members of the group
	var nonMembers []model.UserList

	// Iterate over the query results and populate the nonMembers slice
	for rows.Next() {
		var user model.UserList
		err := rows.Scan(&user.Id, &user.Username, &user.AvatarURL)
		if err != nil {
			return nil, err
		}
		nonMembers = append(nonMembers, user)
	}

	// Check for errors encountered during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return nonMembers, nil
}

func (r *GroupMemberRepository) GetMembers(groupID int) ([]model.UserList, error) {
	query := `
		SELECT id, username, avatar_url
		FROM users
		WHERE id IN (
			SELECT user_id
			FROM group_members
			WHERE group_id = ?
		)
	`

	rows, err := r.db.Query(query, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []model.UserList
	for rows.Next() {
		var member model.UserList
		err := rows.Scan(&member.Id, &member.Username, &member.AvatarURL)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return members, nil
}
