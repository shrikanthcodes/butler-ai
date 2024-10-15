package database

import (
	"errors"
	"github.com/shrikanthcodes/butler-ai/internal/db"
	"github.com/shrikanthcodes/butler-ai/internal/entity"
	"github.com/shrikanthcodes/butler-ai/pkg/logger"
	"github.com/shrikanthcodes/butler-ai/pkg/postgres"
	"gorm.io/gorm"
)

// DatabaseService is a service that interacts with the db.
type DatabaseService struct {
	databaseConnector *postgres.Postgres
}

func NewDatabaseService(url string, log *logger.Logger) (*DatabaseService, error) {
	conn, err := postgres.New(url, log)
	if err != nil {
		log.Fatal("Failed to initialize Postgres", err)
	}

	return &DatabaseService{conn}, nil
}

// Close closes the db connection.
func (service *DatabaseService) Close() {
	service.databaseConnector.Close()
}

// GetConversation retrieves a conversation from the db.
func (store *DBConversationStore) GetConversation(convID string) (entity.Conversation, error) {
	var conv entity.Conversation
	if err := store.db.Where("conv_id = ?", convID).First(&conv).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Return a new empty conversation
			return entity.Conversation{
				ConvID:          convID,
				Conversation:    entity.DialogueArray{},
				RecentDialogues: entity.DialogueArray{},
				Summary:         nil,
				IsActive:        false,
			}, nil
		}
		return entity.Conversation{}, err
	}
	return entity.Conversation{
		ConvID:          conv.ConvID,
		Conversation:    conv.Conversation,
		RecentDialogues: conv.RecentDialogues,
		Summary:         conv.Summary,
		IsActive:        conv.IsActive,
	}, nil
}

// SaveConversation saves a conversation to the db.
func (store *DBConversationStore) SaveConversation(conv *entity.Conversation) error {

	dbConv := entity.Conversation{
		ConvID:       conv.ConvID,
		UserID:       conv.UserID,
		Conversation: conv.Conversation,
		Summary:      conv.Summary,
		IsActive:     conv.IsActive,
	}
	return db.UpdateConversation(&dbConv).Error
}

// DataBuilder returns an empty map, as the template should handle missing data.
func (store *DBConversationStore) DataBuilder(chatType string, outputLength int32) map[string]interface{} {
	return map[string]interface{}{}
}
