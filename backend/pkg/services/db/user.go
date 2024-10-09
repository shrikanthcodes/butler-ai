package db

import (
	"github.com/shrikanthcodes/butler-ai/backend/pkg/models"

	"gorm.io/gorm"
)

// CreateUser inserts a new user into the database
func CreateUser(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}

// UpdateUser updates an existing user in the database
func UpdateUser(db *gorm.DB, user *models.User) error {
	return db.Save(user).Error
}

// CreateConversation inserts a new conversation into the database
func CreateConversation(db *gorm.DB, conversation *models.Conversation) error {
	return db.Create(conversation).Error
}

// UpdateConversation updates an existing conversation in the database
func UpdateConversation(db *gorm.DB, conversation *models.Conversation) error {
	return db.Save(conversation).Error
}
