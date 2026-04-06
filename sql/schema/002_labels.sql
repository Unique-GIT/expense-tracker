-- +goose Up
CREATE TABLE labels(
    id UUID PRIMARY KEY,
    name TEXT 
);

-- +goose Down
DROP TABLE IF EXISTS labels;