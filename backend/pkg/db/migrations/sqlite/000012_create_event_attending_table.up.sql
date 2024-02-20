CREATE TABLE IF NOT EXISTS event_attending (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    event_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    status TEXT CHECK( status IN ('going', 'not going', 'maybe')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);