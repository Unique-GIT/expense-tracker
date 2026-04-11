-- +goose Up
CREATE TABLE labels ( 
    id UUID PRIMARY KEY,
    labelName TEXT NOT NULL,
    user_id UUID NOT NULL,
    UNIQUE(user_id,labelName),
    CONSTRAINT fk_key FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS labels;