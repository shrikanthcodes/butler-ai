package db

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/shrikanthcodes/butler-ai/internal/entity"
	"github.com/shrikanthcodes/butler-ai/pkg/postgres"
)

// GetConversationByID fetches a conversation by its ID from the database.
func GetConversationByID(ctx context.Context, pool *postgres.ConnPool, convID string) (*entity.Conversation, error) {
	query := `
        SELECT conv_id, user_id, chat_type, title, transcript, summary, is_active 
        FROM conversations 
        WHERE conv_id = $1
    `
	row := pool.QueryRow(ctx, query, convID)

	var conv entity.Conversation
	err := row.Scan(
		&conv.ConvID, &conv.UserID, &conv.ChatType, &conv.Title,
		&conv.Transcript, &conv.Summary, &conv.IsActive,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil // No conversation found, return nil
		}
		return nil, err
	}

	return &conv, nil
}

// SaveOrUpdateConversation saves a new conversation or updates an existing one.
func SaveOrUpdateConversation(ctx context.Context, pool *postgres.ConnPool, conv *entity.Conversation) error {
	query := `
        INSERT INTO conversations (conv_id, user_id, chat_type, title, transcript, summary, is_active) 
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (conv_id) DO UPDATE 
        SET transcript = EXCLUDED.transcript, 
            summary = EXCLUDED.summary, 
            is_active = EXCLUDED.is_active
    `
	_, err := pool.Exec(ctx, query,
		&conv.ConvID, &conv.UserID, &conv.ChatType, &conv.Title,
		&conv.Transcript, &conv.Summary, &conv.IsActive,
	)
	return err
}

// AppendTranscript updates dialogues to transcript in real time
func AppendTranscript(ctx context.Context, pool *postgres.ConnPool, convId string, dialogue []entity.Dialogue) error {
	query := `
        UPDATE conversations 
        SET transcript = COALESCE(transcript, '[]'::JSONB) || $1::JSONB 
        WHERE conv_id = $2
    `
	_, err := pool.Exec(ctx, query,
		dialogue, convId,
	)
	return err
}

// GetNRecentDialogues gets n most recent dialogues from a chat given convID
func GetNRecentDialogues(ctx context.Context, pool *postgres.ConnPool, convID string, n int) ([]entity.Dialogue, error) {
	query := `
		SELECT jsonb_agg(dialogue)
		FROM (
			SELECT transcript -> ord AS dialogue
			FROM generate_series(
				jsonb_array_length(transcript) - $1, 
				jsonb_array_length(transcript)
			) AS ord
			WHERE conv_id = $2
		) subquery;
	`
	row := pool.QueryRow(ctx, query, n, convID)

	var dialogue []entity.Dialogue
	err := row.Scan(&dialogue)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return dialogue, nil
}

// GetConversationsByUser fetches all conversations for a specific user.
func GetConversationsByUser(ctx context.Context, pool *postgres.ConnPool, userID string) ([]entity.Conversation, error) {
	query := `
        SELECT user_id, conv_id, is_active, chat_type, title 
        FROM conversations 
        WHERE user_id = $1
    `
	rows, err := pool.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var conversations []entity.Conversation
	for rows.Next() {
		var conv entity.Conversation
		if err := rows.Scan(
			&conv.UserID, &conv.ConvID, &conv.IsActive,
			&conv.ChatType, &conv.Title,
		); err != nil {
			return nil, err
		}
		conversations = append(conversations, conv)
	}

	return conversations, nil
}
