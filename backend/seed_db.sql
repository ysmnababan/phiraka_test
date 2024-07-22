-- CREATE NEW TABLE FOR USER
CREATE TABLE IF NOT EXISTS users(
	UserID INT AUTO_INCREMENT PRIMARY KEY, 
	Username VARCHAR(128) UNIQUE NOT NULL,
	Password VARCHAR(64) NOT NULL,
	CreateTime DATETIME NOT NULL
);

-- SET TIME FOR GMT+07
SET time_zone = '+07:00';

-- SEED DATA FOR USER TABLE
INSERT INTO users (Username, Password, CreateTime) VALUES 
('john_doe', 'password123', NOW()),
('jane_smith', 'password456', NOW()),
('alice_jones', 'password789', NOW());

