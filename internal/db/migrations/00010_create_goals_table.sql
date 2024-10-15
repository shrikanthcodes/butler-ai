-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS goals (
    goal_id VARCHAR PRIMARY KEY,
    user_id VARCHAR NOT NULL,
    goal VARCHAR,
    target TEXT,
    deadline TEXT,
    preference TEXT,
    plan JSONB,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS goals;
-- +goose StatementEnd
