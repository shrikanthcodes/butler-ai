-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS writeups (
    user_id VARCHAR NOT NULL,
    recipe_writeup TEXT,
    shopping_writeup TEXT,
    calorie_tracker_writeup TEXT,
    health_writeup TEXT,
    motivation_writeup TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (user_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS writeups;
-- +goose StatementEnd
