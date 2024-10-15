-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS meal_choices (
    user_id VARCHAR PRIMARY KEY,
    serving_size INT,
    shopping BOOLEAN,
    time_available INT,
    innovative BOOLEAN,
    nutrition_tag JSONB,
    meal_type JSONB,
    fine_tuned BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS meal_choices;
-- +goose StatementEnd
