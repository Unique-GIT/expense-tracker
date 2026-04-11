-- +goose Up
ALTER TABLE transactions
ADD COLUMN created_at TIMESTAMP DEFAULT NOW();

-- +goose Down
ALTER TABLE transactions
DROP COLUMN created_at;