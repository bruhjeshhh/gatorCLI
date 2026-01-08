-- +goose Up
CREATE TABLE users(
	id uuid primary key,
	created_at timestamp,
	updates_at timestamp,
	name text unique not null
);

-- +goose Down
DROP TABLE users;