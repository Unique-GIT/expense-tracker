-- +goose Up
CREATE TABLE transactions(
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    label_id UUID NOT NULL,
    object_name TEXT NOT NULL,
    cost REAL NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_label FOREIGN KEY (label_id) REFERENCES labels(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS transactions;