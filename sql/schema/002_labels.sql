-- +goose Up
CREATE TABLE labels(
    id UUID NOT NULL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE IF EXISTS labels;