-- Table for likes/dislikes
CREATE TABLE IF NOT EXISTS votes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    type TEXT NOT NULL CHECK(type IN ('like', 'dislike')),
    userID INTEGER NOT NULL,
    postID INTEGER,
    commentID INTEGER,
    FOREIGN KEY (commentID) REFERENCES comments(id) ON DELETE CASCADE,
    FOREIGN KEY (userID) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (postID) REFERENCES posts(id) ON DELETE CASCADE
);