-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    userName TEXT NOT NULL,
    userNumber TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE IF EXISTS users;