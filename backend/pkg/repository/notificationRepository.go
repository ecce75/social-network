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

// GetAllNotifications retrieves all notifications from the database.
func (r *NotificationRepository) GetAllNotifications() ([]model.Notification, error) {
	query := `SELECT * FROM notifications`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []model.Notification
	for rows.Next() {
		var notification model.Notification
		if err := rows.Scan(&notification.Id, &notification.Type, &notification.Message); err != nil {
			return nil, err
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
	query := `INSERT INTO notifications (user_id, type, message, is_read) VALUES (?, ?, ?, ?)`
	result, err := r.db.Exec(query, notification.UserId, notification.Type, notification.Message, notification.IsRead)
	if err != nil {
		return 0, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertID, nil
}

// GetNotificationByID retrieves a specific notification by its ID from the database.
func (r *NotificationRepository) GetNotificationByID(id int) (model.Notification, error) {
	query := `SELECT * FROM notification WHERE id = ?`
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
