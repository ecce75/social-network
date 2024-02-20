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
func (r *EventRepository) GetAllEvents() ([]model.Event, error) {
	query := `SELECT * FROM events`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []model.Event
	for rows.Next() {
		var event model.Event
		if err := rows.Scan(&event.Id, &event.Title, &event.Description, &event.Location, &event.StartTime, &event.EndTime, &event.CreatedAt); err != nil {
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
	query := `INSERT INTO events (creator_id, title, description, location, start_time, end_time) VALUES (?, ?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query, event.CreatorId, event.Title, event.Description, event.Location)
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
	err := row.Scan(&event.Id, &event.CreatorId, &event.Title, &event.Description, &event.CreatedAt)
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
func (r *EventRepository) AddOrUpdateAttendance(eventID, userID int, status string) (int64, error) {
	// Check if the attendance record already exists
	query := `
		SELECT id FROM event_attendance
		WHERE event_id = ? AND user_id = ?
	`

	row := r.db.QueryRow(query, eventID, userID)
	var attendanceID int64
	err := row.Scan(&attendanceID)

	if err == nil {
		// Attendance record exists, update the status
		updateQuery := `
			UPDATE event_attendance
			SET status = ?
			WHERE id = ?
		`

		_, err := r.db.Exec(updateQuery, status, attendanceID)
		return attendanceID, err
	}

	// Attendance record doesn't exist, insert a new record
	insertQuery := `
		INSERT INTO event_attendance (event_id, user_id, status)
		VALUES (?, ?, ?)
	`

	result, err := r.db.Exec(insertQuery, eventID, userID, status)
	if err != nil {
		return 0, err
	}

	attendanceID, err = result.LastInsertId()
	return attendanceID, err
}

// GetAttendanceByEventID retrieves attendance records for a specific event from the database.
func (r *EventRepository) GetAttendanceByEventID(eventID int) ([]model.EventAttendance, error) {
	query := `
		SELECT * FROM event_attendance
		WHERE event_id = ?
	`

	rows, err := r.db.Query(query, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attendanceList []model.EventAttendance
	for rows.Next() {
		var attendance model.EventAttendance
		if err := rows.Scan(&attendance.Id, &attendance.EventId, &attendance.UserId, &attendance.Status, &attendance.CreatedAt); err != nil {
			return nil, err
		}
		attendanceList = append(attendanceList, attendance)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return attendanceList, nil
}
