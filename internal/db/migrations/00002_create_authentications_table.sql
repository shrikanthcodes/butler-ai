-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS authentications (
    user_id VARCHAR PRIMARY KEY,
    role VARCHAR,
    passwd TEXT,
    last_updated TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS authentications;
-- +goose StatementEnd
