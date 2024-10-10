package database

import (
	"github.com/shrikanthcodes/butler-ai/internal/entity"
)

// CreateConversation inserts a new conversation into the database
func (DBC *DBConnector) CreateConversation(conversation *entity.Conversation) error {
	return DBC.db.Create(conversation).Error
}

// UpdateConversationBatch updates an existing conversation in the database in table conversations given a convID
func (DBC *DBConnector) UpdateConversationBatch(conversation *entity.Conversation) error {
	return DBC.db.Save(conversation).Error
}

// UpdateConversation updates an existing conversation with low overhead in the database in table conversations given a convID

// GetConversation retrieves a conversation from the database for a given convID from table conversations
func (DBC *DBConnector) GetConversation(convID string) (entity.Conversation, error) {
	var conversation entity.Conversation
	if err := DBC.db.Where("conv_id = ?", convID).First(&conversation).Error; err != nil {
		return entity.Conversation{}, err
	}
	return conversation, nil
}

// GetConversationsByUserID retrieves all conversations from the database for a given userID from table conversations
func (DBC *DBConnector) GetConversationsByUserID(userID string) ([]entity.Conversation, error) {
	var conversations []entity.Conversation
	if err := DBC.db.Where("user_id = ?", userID).Find(&conversations).Error; err != nil {
		return nil, err
	}
	return conversations, nil
}

// DeleteConversation deletes a conversation from the database for a given convID from table conversations
func (DBC *DBConnector) DeleteConversation(convID string) error {
	return DBC.db.Where("conv_id = ?", convID).Delete(&entity.Conversation{}).Error
}

//ADMIN (Needs guardrails)

// DeleteConversations deletes all conversations from the database for a given userID from table conversations
func (DBC *DBConnector) DeleteConversations(userID string) error {
	return DBC.db.Where("user_id = ?", userID).Delete(&entity.Conversation{}).Error
}

// GetAllConversations retrieves all conversations from the database from table conversations
func (DBC *DBConnector) GetAllConversations() ([]entity.Conversation, error) {
	var conversations []entity.Conversation
	if err := DBC.db.Find(&conversations).Error; err != nil {
		return nil, err
	}
	return conversations, nil
}

// DeleteAllConversations deletes all conversations from the database from table conversations;
func (DBC *DBConnector) DeleteAllConversations() error {
	return DBC.db.Delete(&entity.Conversation{}).Error
}

// GetActiveConversations retrieves all active conversations from the database for a given userID from table conversations
func (DBC *DBConnector) GetActiveConversations(userID string) ([]entity.Conversation, error) {
	var conversations []entity.Conversation
	if err := DBC.db.Where("user_id = ? AND is_active = true", userID).Find(&conversations).Error; err != nil {
		return nil, err
	}
	return conversations, nil
}
