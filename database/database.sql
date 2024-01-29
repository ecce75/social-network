-- Users Table
CREATE TABLE IF NOT EXISTS
    `Users` (
        `userID` INTEGER PRIMARY KEY AUTOINCREMENT,
        `username` TEXT NOT NULL UNIQUE,
        `password` TEXT NOT NULL, -- Store hashed password
        `email` TEXT NOT NULL UNIQUE,
        `firstName` TEXT NOT NULL,
        `lastName` TEXT NOT NULL,
        `age` INTEGER NOT NULL,
        `gender` TEXT NOT NULL,
        `avatar` TEXT, -- Optional field for user's profile picture
        `bio` TEXT, -- Optional field for user's biography
        `joinDate` DATETIME DEFAULT CURRENT_TIMESTAMP
    );

-- Categories Table
CREATE TABLE IF NOT EXISTS
    `Categories` (
        `categoryID` INTEGER PRIMARY KEY AUTOINCREMENT,
        `title` TEXT NOT NULL UNIQUE,
        `description` TEXT NOT NULL,
        `createdAt` DATETIME DEFAULT CURRENT_TIMESTAMP
    );

-- Comments Table
CREATE TABLE IF NOT EXISTS
    `Comments` (
        `commentID` INTEGER PRIMARY KEY AUTOINCREMENT,
        `userID` INTEGER NOT NULL,
        `postID` INTEGER NOT NULL,
        `content` TEXT NOT NULL,
        `createdAt` DATETIME DEFAULT CURRENT_TIMESTAMP,
        `updatedAt` DATETIME,
        FOREIGN KEY (`userID`) REFERENCES `Users` (`userID`) ON DELETE CASCADE,
        FOREIGN KEY (`postID`) REFERENCES `Posts` (`postID`) ON DELETE CASCADE
    );

-- ChatMessages Table
CREATE TABLE IF NOT EXISTS
    `ChatMessages` (
        `messageID` INTEGER PRIMARY KEY AUTOINCREMENT,
        `senderID` INTEGER NOT NULL,
        `receiverID` INTEGER NOT NULL,
        `message` TEXT NOT NULL,
        `createdAt` DATETIME DEFAULT CURRENT_TIMESTAMP,
        `updatedAt` DATETIME,
        `status` TEXT, -- 'sent', 'delivered', 'read'
        FOREIGN KEY (`senderID`) REFERENCES `Users` (`userID`),
        FOREIGN KEY (`receiverID`) REFERENCES `Users` (`userID`)
    );

-- Sessions Table
CREATE TABLE IF NOT EXISTS
    `Sessions` (
        `sessionID` TEXT PRIMARY KEY,
        `userID` INTEGER UNIQUE NOT NULL,
        `expiresAt` TIMESTAMP,
        FOREIGN KEY (`userID`) REFERENCES `Users` (`userID`) ON DELETE CASCADE
    );

-- Posts Table
CREATE TABLE IF NOT EXISTS
    `Posts` (
        `postID` INTEGER PRIMARY KEY AUTOINCREMENT,
        `title` TEXT NOT NULL,
        `content` TEXT NOT NULL,
        `createdAt` DATETIME DEFAULT CURRENT_TIMESTAMP,
        `updatedAt` DATETIME,
        `categoryID` INTEGER NOT NULL,
        `userID` INTEGER NOT NULL,
        FOREIGN KEY (`categoryID`) REFERENCES `Categories` (`categoryID`) ON DELETE CASCADE,
        FOREIGN KEY (`userID`) REFERENCES `Users` (`userID`) ON DELETE CASCADE
    );

-- Ratings Table
CREATE TABLE IF NOT EXISTS
    `Ratings` (
        `ratingID` INTEGER PRIMARY KEY AUTOINCREMENT,
        `type` TEXT CHECK (`type` IN ('like', 'dislike')) NOT NULL, -- Using Text with a CHECK constraint
        `userID` INTEGER NOT NULL,
        `postID` INTEGER,
        `commentID` INTEGER,
        FOREIGN KEY (`commentID`) REFERENCES `Comments` (`commentID`) ON DELETE CASCADE,
        FOREIGN KEY (`userID`) REFERENCES `Users` (`userID`) ON DELETE CASCADE,
        FOREIGN KEY (`postID`) REFERENCES `Posts` (`postID`) ON DELETE CASCADE
    );

-- Additional Table for Followers/Following Mechanism
CREATE TABLE IF NOT EXISTS
    `Followers` (
        `followerID` INTEGER NOT NULL,
        `followingID` INTEGER NOT NULL,
        PRIMARY KEY (`followerID`, `followingID`),
        FOREIGN KEY (`followerID`) REFERENCES `Users` (`userID`) ON DELETE CASCADE,
        FOREIGN KEY (`followingID`) REFERENCES `Users` (`userID`) ON DELETE CASCADE
    );

-- Additional Table for Group Functionality (if applicable)
CREATE TABLE IF NOT EXISTS
    `Groups` (
        `groupID` INTEGER PRIMARY KEY AUTOINCREMENT,
        `name` TEXT NOT NULL,
        `description` TEXT,
        `createdAt` DATETIME DEFAULT CURRENT_TIMESTAMP,
        `ownerID` INTEGER NOT NULL,
        FOREIGN KEY (`ownerID`) REFERENCES `Users` (`userID`) ON DELETE CASCADE
    );

-- Table for Group Memberships
CREATE TABLE IF NOT EXISTS
    `GroupMembers` (
        `groupID` INTEGER NOT NULL,
        `userID` INTEGER NOT NULL,
        PRIMARY KEY (`groupID`, `userID`),
        FOREIGN KEY (`groupID`) REFERENCES `Groups` (`groupID`) ON DELETE CASCADE,
        FOREIGN KEY (`userID`) REFERENCES `Users` (`userID`) ON DELETE CASCADE
    );

-- Event Table for Groups (if applicable)
CREATE TABLE IF NOT EXISTS
    `GroupEvents` (
        `eventID` INTEGER PRIMARY KEY AUTOINCREMENT,
        `groupID` INTEGER NOT NULL,
        `title` TEXT NOT NULL,
        `description` TEXT,
        `eventDate` DATETIME NOT NULL,
        `createdAt` DATETIME DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (`groupID`) REFERENCES `Groups` (`groupID`) ON DELETE CASCADE
    );