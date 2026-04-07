-- +goose Up
CREATE TABLE labels(
    id UUID NOT NULL PRIMARY KEY,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS labels;