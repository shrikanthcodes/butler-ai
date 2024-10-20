package database

import (
	"context"
	"fmt"
	"github.com/shrikanthcodes/butler-ai/internal/db"
	"github.com/shrikanthcodes/butler-ai/internal/entity"
	"github.com/shrikanthcodes/butler-ai/pkg/logger"
	"github.com/shrikanthcodes/butler-ai/pkg/postgres"
)

// DbService is a service that interacts with the db.
type DbService struct {
	conn *postgres.ConnPool
	log  *logger.Logger
}

func NewDatabaseService(url string, log *logger.Logger) (*DbService, error) {
	pool, err := postgres.New(url, log)
	if err != nil {
		log.Fatal("Failed to initialize Postgres", err)
	}

	return &DbService{pool,
		log}, nil
}

// Close closes the db connection.
func (s *DbService) Close() error {
	s.conn.Close()
	return nil
}

// GetConversation retrieves a conversation by ID or returns a new empty conversation.
func (s *DbService) GetConversation(ctx context.Context, convID, chatType string) (entity.Conversation, error) {
	conv, err := db.GetConversationByID(ctx, s.conn, convID)
	if err != nil {
		return entity.Conversation{}, fmt.Errorf("failed to get conversation: %w", err)
	}

	if conv == nil {
		// Return a new empty conversation if not found
		return entity.Conversation{
			ConvID:     convID,
			Transcript: entity.JSONB[[]entity.Dialogue]{Data: []entity.Dialogue{}},
			ChatType:   chatType,
			Title:      "",
			Summary:    nil,
			IsActive:   false,
		}, nil
	}
	return *conv, nil
}

// SaveConversation saves or updates a conversation.
func (s *DbService) SaveConversation(ctx context.Context, conv *entity.Conversation) error {
	return db.SaveOrUpdateConversation(ctx, s.conn, conv)
}

// GetConversationsByUser retrieves all conversations for a specific user.
func (s *DbService) GetConversationsByUser(ctx context.Context, userID string) ([]entity.Conversation, error) {
	return db.GetConversationsByUser(ctx, s.conn, userID)
}

func (s *DbService) GetRecentDialogues(ctx context.Context, convID string) []entity.Dialogue {
	var recent []entity.Dialogue
	//TODO: check if conversation exists in cache
	recent, err := db.GetNRecentDialogues(ctx, s.conn, convID, 10)
	if err != nil {
		s.log.Error("Failed to get recent dialogues", err)
		return nil
	}
	return recent
}
