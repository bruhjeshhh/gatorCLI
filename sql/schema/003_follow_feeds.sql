-- +goose Up
CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feed_id uuid not null REFERENCES feeds(id) on delete CASCADE,
    unique(user_id,feed_id)
);

-- +goose Down
DROP TABLE feed_follows;