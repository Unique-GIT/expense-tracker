-- +goose Up
CREATE TABLE transactions(
    id UUID PRIMARY KEY,
    object TEXT NOT NULL,
    cost REAL NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id),
    label_id UUID NOT NULL REFERENCES labels(id)
);

-- +goose Down
DROP TABLE IF EXISTS transactions;