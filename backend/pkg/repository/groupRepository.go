// Package repository provides the implementation of the repository layer for the social network backend.
// It includes functions to interact with the database for managing groups.
package repository

import (
	"backend/pkg/model"
	"database/sql"
	"fmt"
	"os"
	"time"
)

// GroupRepository is a repository for managing groups in the database.
type GroupRepository struct {
	db *sql.DB
}

// NewGroupRepository creates a new instance of GroupRepository.
func NewGroupRepository(db *sql.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

// GetAllGroups retrieves all groups from the database.
// It returns a slice of Group objects and an error if any.
func (r *GroupRepository) GetAllGroups() ([]model.Group, error) {
	// SQL query to select all groups
	query := `SELECT id, creator_id, title, description, image_url, created_at, updated_at FROM groups`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []model.Group
	for rows.Next() {
		var group model.Group
		err := rows.Scan(&group.Id, &group.CreatorId, &group.Title, &group.Description, &group.Image, &group.CreatedAt, &group.UpdatedAt)
		if err != nil && err != sql.ErrNoRows {
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
func (r *GroupRepository) CreateGroup(group model.Group) (int64, error) {
	query := `INSERT INTO groups (creator_id, title, description) VALUES (?, ?, ?)`
	result, err := r.db.Exec(query, group.CreatorId, group.Title, group.Description)
	if err != nil {
		return 0, err
	}
	groupID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error getting last inserted post id")
	}
	// add group owner as a member
	query = `INSERT INTO group_members (group_id, user_id) VALUES (?, ?)`
	_, err = r.db.Exec(query, groupID, group.CreatorId)
	if err != nil {
		return 0, err
	}

	// set group image URL
	var ImageURL = os.Getenv("NEXT_PUBLIC_URL") + ":" + os.Getenv("NEXT_PUBLIC_BACKEND_PORT") + "/images/groups/" + fmt.Sprint(groupID) + ".jpg"
	query = `UPDATE groups SET image_url = ? WHERE id = ?`
	_, err = r.db.Exec(query, ImageURL, groupID)
	if err != nil {
		return 0, err
	}

	return groupID, nil
}

// GetGroupByID retrieves a group by ID from the database.
// It returns the Group object and an error if any.
func (r *GroupRepository) GetGroupByID(id int) (model.Group, error) {
	query := `SELECT * FROM groups WHERE id = ?`
	row := r.db.QueryRow(query, id)
	var group model.Group
	err := row.Scan(&group.Id, &group.CreatorId, &group.Title, &group.Description, &group.Image, &group.CreatedAt, &group.UpdatedAt)
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

// LogGroupDeletion logs the deletion of a group.
func (r *GroupRepository) LogGroupDeletion(groupID int) error {
	query := `UPDATE groups SET deleted = true WHERE id = ?`
	_, err := r.db.Exec(query, groupID)
	return err
}

func (r *GroupRepository) GetGroupTitleByID(id int) (string, error) {
	query := `SELECT title FROM groups WHERE id = ?`
	row := r.db.QueryRow(query, id)
	var title string
	err := row.Scan(&title)
	if err != nil {
		return "", err
	}
	return title, nil
}
