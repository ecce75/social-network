CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    group_id INTEGER,
    title TEXT NOT NULL,
    content TEXT,
    image_url TEXT,
    privacy_setting TEXT NOT NULL CHECK(privacy_setting IN ('public', 'private', 'custom')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
);
