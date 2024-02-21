CREATE TABLE IF NOT EXISTS group_invitations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    group_id INTEGER NOT NULL,
    join_user_id INTEGER NOT NULL,
    invite_user_id INTEGER,
    status TEXT NOT NULL CHECK( status IN ('pending', 'accepted', 'declined')) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (group_id) REFERENCES groups(id),
    FOREIGN KEY (join_user_id) REFERENCES users(id),
    FOREIGN KEY (invite_user_id) REFERENCES users(id)
    );