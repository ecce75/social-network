package repository

import (
	"backend/pkg/model"
	"database/sql"
)
type GroupRepository struct {
    db *sql.DB
}

func NewGroupRepository(db *sql.DB) *GroupRepository {
    return &GroupRepository{db: db}
}

func (r *GroupRepository) GetAllGroups() ([]model.Group, error) {
	query  := `SELECT * FROM groups`
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

func (r *GroupRepository) UpdateGroup(group model.Group) error {
    query := `UPDATE groups SET creator_id = ?, title = ?, description = ? WHERE id = ?`
    _, err := r.db.Exec(query, group.CreatorId, group.Title, group.Description, group.Id)
    return err
}

func (r *GroupRepository) DeleteGroup(id int) error {
    query := `DELETE FROM groups WHERE id = ?`
    _, err := r.db.Exec(query, id)
    return err
}