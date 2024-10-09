package chat

import (
	"errors"

	"github.com/shrikanthcodes/butler-ai/backend/pkg/models"

	"gorm.io/gorm"
)

type DBConversationStore struct {
	db *gorm.DB
}

func NewDBConversationStore(db *gorm.DB) *DBConversationStore {
	return &DBConversationStore{
		db: db,
	}
}

// GetConversation retrieves a conversation from the database.
func (store *DBConversationStore) GetConversation(convID string) (models.Conversation, error) {
	var conv models.Conversation
	if err := store.db.Where("conv_id = ?", convID).First(&conv).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Return a new empty conversation
			return models.Conversation{
				ConvID:          convID,
				Conversation:    models.DialogueArray{},
				RecentDialogues: models.DialogueArray{},
				Summary:         nil,
				IsActive:        false,
			}, nil
		}
		return models.Conversation{}, err
	}
	return models.Conversation{
		ConvID:          conv.ConvID,
		Conversation:    conv.Conversation,
		RecentDialogues: conv.RecentDialogues,
		Summary:         conv.Summary,
		IsActive:        conv.IsActive,
	}, nil
}

// SaveConversation saves a conversation to the database.
func (store *DBConversationStore) SaveConversation(conv *models.Conversation) error {

	dbConv := models.Conversation{
		ConvID:       conv.ConvID,
		UserID:       conv.UserID,
		Conversation: conv.Conversation,
		Summary:      conv.Summary,
		IsActive:     conv.IsActive,
	}
	return store.db.Save(&dbConv).Error
}

// DataBuilder returns an empty map, as the template should handle missing data.
func (store *DBConversationStore) DataBuilder(chatType string, outputLength int32) map[string]interface{} {
	return map[string]interface{}{}
}
