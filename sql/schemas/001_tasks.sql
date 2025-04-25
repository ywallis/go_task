-- +goose Up

CREATE TABLE tasks (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL, 
	updated_at TIMESTAMP NOT NULL	
);

-- +goose Down

DROP TABLE IF EXISTS tasks;
