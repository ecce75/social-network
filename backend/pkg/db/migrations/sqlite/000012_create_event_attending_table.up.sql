CREATE TABLE IF NOT EXISTS event_attending (
    id INT AUTO_INCREMENT PRIMARY KEY,
    event_id INT NOT NULL,
    user_id INT NOT NULL,
    status TEXT CHECK( status IN ('going', 'not going', 'maybe')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);