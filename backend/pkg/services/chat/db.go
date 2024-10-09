package chat

import (
	"errors"
	"github.com/shrikanthcodes/butler-ai/backend/internal/entity"

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

// SaveConversation saves a conversation to the database.
func (store *DBConversationStore) SaveConversation(conv *entity.Conversation) error {

	dbConv := entity.Conversation{
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
