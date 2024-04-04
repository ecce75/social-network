package repository

import (
	"backend/pkg/model"
	"database/sql"
)

// NotificationRepository handles database operations related to notifications.
type NotificationRepository struct {
	db *sql.DB
}

// NewNotificationRepository creates a new instance of NotificationRepository.
func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (r *NotificationRepository) GetNotificationsByUserId(id int) ([]model.Notification, error) {
	query := `SELECT * FROM notifications WHERE user_id = ?`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []model.Notification
	for rows.Next() {
		var notification model.Notification
		var groupId, senderId sql.NullInt64 // Use sql.NullInt64 for nullable integers

		// Scan the row with sql.NullInt64 variables for nullable columns
		if err := rows.Scan(&notification.Id, &notification.UserId, &groupId, &senderId, &notification.Type, &notification.Message, &notification.IsRead, &notification.CreatedAt); err != nil {
			return nil, err
		}

		// Check if groupId is valid, then assign its value to the Notification struct
		if groupId.Valid {
			val := int(groupId.Int64) // Convert sql.NullInt64 to int
			notification.GroupId = val
		}

		// Check if senderId is valid, then assign its value to the Notification struct
		if senderId.Valid {
			val := int(senderId.Int64) // Convert sql.NullInt64 to int
			notification.SenderId = val
		}

		notifications = append(notifications, notification)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return notifications, nil
}

// CreateNotification adds a new notification to the database.
func (r *NotificationRepository) CreateNotification(notification model.Notification) (int64, error) {
	baseQuery := `INSERT ` + `INTO notifications (user_id, type, message`
	valuesQuery := `VALUES (?, ?, ?, ?)`
	args := []interface{}{notification.UserId, notification.Type, notification.Message}

	if notification.GroupId != 0 {
		baseQuery += `, group_id) `
		args = append(args, notification.GroupId)
	} else if notification.SenderId != 0 {
		baseQuery += `, sender_id) `
		args = append(args, notification.SenderId)
	}

	query := baseQuery + valuesQuery

	result, err := r.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertID, nil
}

func (r *NotificationRepository) EditFriendNotificationMessage(userID, senderID int, message string) error {
	query := `UPDATE notifications SET message = ?, is_read = false WHERE user_id = ? AND sender_id = ? AND type = 'friend'`
	_, err := r.db.Exec(query, message, userID, senderID)
	return err
}

func (r *NotificationRepository) EditGroupNotificationMessage(userID, groupID int, message string) error {
	query := `UPDATE notifications SET message = ?, is_read = false WHERE user_id = ? AND group_id = ? AND type = 'group'`
	_, err := r.db.Exec(query, message, userID, groupID)
	return err
}

// GetNotificationByID retrieves a specific notification by its ID from the database.
func (r *NotificationRepository) GetNotificationByID(id int) (model.Notification, error) {
	query := `SELECT * FROM notifications WHERE id = ?`
	row := r.db.QueryRow(query, id)
	var notification model.Notification
	err := row.Scan(&notification.Id, &notification.Type, &notification.Message, &notification.IsRead)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Notification{}, nil
		}
		return model.Notification{}, err
	}
	return notification, nil
}

// MarkNotificationAsRead updates the read status of a notification in the database.
func (r *NotificationRepository) MarkNotificationAsRead(id int) error {
	query := `UPDATE notifications SET is_read = true WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *NotificationRepository) DeleteNotification(id int) error {
	query := `DELETE FROM notifications WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
