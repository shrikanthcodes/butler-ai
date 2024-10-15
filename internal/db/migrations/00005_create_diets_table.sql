-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS diets (
    user_id VARCHAR PRIMARY KEY,
    preferred_units JSONB,
    favorite_recipes JSONB,
    disliked_recipes JSONB,
    favorite_items JSONB,
    disliked_items JSONB,
    favorite_cuisine JSONB,
    disliked_cuisine JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS diets;
-- +goose StatementEnd
