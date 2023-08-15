CREATE TABLE users (
	ID SERIAL PRIMARY KEY,
	FirstName VARCHAR(50) NOT NULL CHECK (FirstName <> ''),
	LastName VARCHAR(50) NOT NULL CHECK (LastName <> ''),
	Email VARCHAR(100) NOT NULL UNIQUE CHECK (Email <> ''),
	Password_hash CHAR(60) NOT NULL CHECK (Password_hash <> '')
);

CREATE TABLE todos (
	ID SERIAL PRIMARY KEY,
	USER_ID INTEGER NOT NULL REFERENCES users (ID) ON DELETE CASCADE,
	Task VARCHAR(255) NOT NULL,
	Status BOOLEAN NOT NULL
);

INSERT INTO users (FirstName, LastName, Email, Password_hash) VALUES ('First', 'Last', 'first@last.com', '$2a$10$KbD6B.lSAx6gW0IEnRjmVu3XqA8/5nwUhpmSl0HmPlW2I7F/sME7G');

INSERT INTO todos (USER_ID, Task, Status) VALUES (1, 'Complete the Docker setup', false);
INSERT INTO todos (USER_ID, Task, Status) VALUES (1, 'Learn more about PostgreSQL', false);
INSERT INTO todos (USER_ID, Task, Status) VALUES (1, 'Integrate database with Go application', true);
