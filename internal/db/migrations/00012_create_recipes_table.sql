-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS recipes (
    recipe_id VARCHAR PRIMARY KEY,
    name VARCHAR,
    tags JSONB,
    cuisine VARCHAR,
    ingredients JSONB,
    instructions JSONB,
    nutritional_info JSONB,
    user_id VARCHAR,
    url TEXT,
    recipe_time JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE SET NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS recipes;
-- +goose StatementEnd
