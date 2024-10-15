-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS profiles (
    user_id VARCHAR PRIMARY KEY,
    age INT,
    height INT,
    weight INT,
    height_unit VARCHAR,
    weight_unit VARCHAR,
    gender VARCHAR,
    lifestyle VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS profiles;
-- +goose StatementEnd
