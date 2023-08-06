CREATE TABLE todos (
	ID SERIAL PRIMARY KEY,
	Task VARCHAR(255) NOT NULL,
	Status BOOLEAN NOT NULL
);

INSERT INTO todos (Task, Status) VALUES ('Complete the Docker setup', false);
INSERT INTO todos (Task, Status) VALUES ('Learn more about PostgreSQL', false);
INSERT INTO todos (Task, Status) VALUES ('Integrate database with Go application', true);
