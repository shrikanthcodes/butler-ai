-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS medical (
    user_id VARCHAR PRIMARY KEY,
    health_conditions JSONB,
    medications JSONB,
    allergies JSONB,
    dietary_restrictions JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS medical;
-- +goose StatementEnd
