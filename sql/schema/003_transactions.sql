-- +goose Up
CREATE TABLE transactions(
    id UUID PRIMARY KEY,
    object TEXT NOT NULL,
    cost REAL NOT NULL,
    user_id UUID REFERENCES users(id),
    label_id UUID REFERENCES labels(id)
);

-- +goose Down
DROP TABLE IF EXISTS transactions;