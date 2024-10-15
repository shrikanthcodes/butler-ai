-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS shopping (
    user_id VARCHAR PRIMARY KEY,
    shopping_mode VARCHAR,
    budget_currency VARCHAR,
    budget INT,
    ease_of_availability VARCHAR,
    shopping_list JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS shopping;
-- +goose StatementEnd
