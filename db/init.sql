CREATE TABLE todos (
	ID SERIAL PRIMARY KEY,
	Task VARCHAR(255) NOT NULL,
	Status BOOLEAN NOT NULL
);

CREATE TABLE users (
	ID SERIAL PRIMARY KEY,
	FirstName VARCHAR(50) NOT NULL CHECK (FirstName <> ''),
	LastName VARCHAR(50) NOT NULL CHECK (LastName <> ''),
	Email VARCHAR(100) NOT NULL UNIQUE CHECK (Email <> ''),
	Password_hash CHAR(60) NOT NULL CHECK (Password_hash <> '')
);

INSERT INTO todos (Task, Status) VALUES ('Complete the Docker setup', false);
INSERT INTO todos (Task, Status) VALUES ('Learn more about PostgreSQL', false);
INSERT INTO todos (Task, Status) VALUES ('Integrate database with Go application', true);
