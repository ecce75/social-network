// Package repository provides the implementation of the repository layer for the social network backend.
// It includes functions to interact with the database for managing events.

package repository

import (
	"backend/pkg/model"
	"database/sql"
	"time"
)

// EventRepository is a repository for managing events in the database.
type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

// GetAllEvents retrieves all events from the database.
// It returns a slice of Event objects and an error if any.
func (r *EventRepository) GetAllGroupEvents(groupID int) ([]model.Event, error) {
	query := `SELECT * FROM events WHERE group_id = ?`
	rows, err := r.db.Query(query, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []model.Event
	for rows.Next() {
		var event model.Event
		if err := rows.Scan(&event.Id, &event.CreatorId, &event.GroupId, &event.Title, &event.Description, &event.Location, &event.StartTime, &event.EndTime, &event.CreatedAt); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

// CreateEvent creates a new event in the database.
// It returns the ID of the newly created event and an error if any.
func (r *EventRepository) CreateEvent(event model.Event) (int64, error) {
	query := `INSERT INTO events (creator_id, group_id, title, description, location, start_time, end_time) VALUES (?, ?, ?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query, event.CreatorId, event.GroupId, event.Title, event.Description, event.Location, event.StartTime, event.EndTime)
	if err != nil {
		return 0, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertID, nil
}

// GetEventByID retrieves an event by ID from the database.
// It returns the Event object and an error if any.
func (r *EventRepository) GetEventByID(id int) (model.Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := r.db.QueryRow(query, id)
	var event model.Event
	err := row.Scan(&event.Id, &event.CreatorId, &event.GroupId, &event.Title, &event.Description, &event.Location, &event.StartTime, &event.EndTime, &event.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Event{}, nil
		}
		return model.Event{}, err
	}
	return event, nil
}

// EditEvent updates an event in the database.
// It returns an error if any.
func (r *EventRepository) EditEvent(event model.Event) error {
	query := `UPDATE events SET title = ?, description = ?, location = ?, start_time = ?, end_time = ? WHERE id = ?`
	_, err := r.db.Exec(query, event.Title, event.Description, time.Now(), event.Id)
	return err
}

// DeleteEvent deletes an event from the database.
// It returns an error if any.
func (r *EventRepository) DeleteEvent(id int) error {
	query := `DELETE FROM events WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

// GetEventsByGroupID retrieves events associated with a specific group ID from the database.
func (r *EventRepository) GetEventsByGroupID(groupID int) ([]model.Event, error) {
	query := `
		SELECT * FROM events WHERE group_id = ?`

	rows, err := r.db.Query(query, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []model.Event
	for rows.Next() {
		var event model.Event
		if err := rows.Scan(&event.Id, &event.CreatorId, &event.GroupId, &event.Title, &event.Description, &event.Location, &event.StartTime, &event.EndTime, &event.CreatedAt); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

// AddOrUpdateAttendance adds or updates attendance status for an event and a user.
// It returns the ID of the attendance record and an error if any.
func (r *EventRepository) AddOrUpdateAttendance(eventID, userID int, status string) error {
	// Attendance record exists, update the status
	query := `
   		INSERT OR REPLACE INTO event_attending (event_id, user_id, status)
   		VALUES (?, ?, ?)
  	`
	_, err := r.db.Exec(query, eventID, userID, status)
	if err != nil {
		return err
	}
	return nil
}

// GetAttendanceByEventID retrieves attendance records for a specific event from the database.
func (r *EventRepository) GetAttendanceByEventID(eventID int) ([]model.Attendance, error) {
	query := `
		SELECT id, username, avatar_url FROM users 
		WHERE id IN (
			SELECT user_id FROM event_attending
			WHERE event_id = ? AND status = 'going'
		)
	`

	rows, err := r.db.Query(query, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attendanceList []model.Attendance
	for rows.Next() {
		var attendance model.Attendance
		if err := rows.Scan(&attendance.Id, &attendance.Username, &attendance.AvatarURL); err != nil {
			return nil, err
		}
		attendanceList = append(attendanceList, attendance)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return attendanceList, nil
}
