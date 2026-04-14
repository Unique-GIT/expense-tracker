-- +goose Up
ALTER TABLE transactions
ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT NOW();

-- +goose Down
ALTER TABLE transactions
DROP COLUMN created_at;