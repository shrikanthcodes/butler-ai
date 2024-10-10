package database

import (
	"github.com/shrikanthcodes/butler-ai/internal/entity"
)

// CreateUser inserts a new user into the database
func (DBC *DBConnector) CreateUser(user *entity.User) error {
	return DBC.db.Create(user).Error
}

// UpdateUser updates an existing user in the database
func (DBC *DBConnector) UpdateUser(user *entity.User) error {
	return DBC.db.Save(user).Error
}

// GetUser retrieves a user from the database
func (DBC *DBConnector) GetUser(userID string) (entity.User, error) {
	var user entity.User
	if err := DBC.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// DeleteUser deletes a user from the database
func (DBC *DBConnector) DeleteUser(userID string) error {
	return DBC.db.Where("user_id = ?", userID).Delete(&entity.User{}).Error
}

// CheckUserExists checks if a user exists in the database for a given email, returns userID if exists
func (DBC *DBConnector) CheckUserExists(email string) (string, error) {
	var user entity.User
	if err := DBC.db.Where("email = ?", email).First(&user).Error; err != nil {
		return "", err
	}
	return user.UserID, nil
}

// CheckPassword checks if a password matches the hashed password in the database for a given userID
func (DBC *DBConnector) CheckPassword(userID, password string) error {
	var auth entity.Authentication
	if err := DBC.db.Where("user_id = ? AND password = ?", userID, password).First(&auth).Error; err != nil {
		return err
	}
	return nil
}

// GetUserIDByEmail retrieves a user's ID from the database for a given email
func (DBC *DBConnector) GetUserIDByEmail(email string) (string, error) {
	var user entity.User
	if err := DBC.db.Where("email = ?", email).First(&user).Error; err != nil {
		return "", err
	}
	return user.UserID, nil
}

// Register a new user with the given email, password, and role (create entry in users and authentications table)

// RegisterUser inserts a new user into the database
func (DBC *DBConnector) RegisterUser(user *entity.User, auth *entity.Authentication) error {
	tx := DBC.db.Begin()
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
	}
	if err := tx.Create(auth).Error; err != nil {
		tx.Rollback()
	}
	return nil
}

// UpdatePassword updates a user's password in the database
func (DBC *DBConnector) UpdatePassword(userID, password string) error {
	return DBC.db.Model(&entity.Authentication{}).Where("user_id = ?", userID).Update("password", password).Error
}

// UpdateEmail updates a user's email in the database
func (DBC *DBConnector) UpdateEmail(userID, email string) error {
	return DBC.db.Model(&entity.User{}).Where("user_id = ?", userID).Update("email", email).Error
}

// UpdatePhone updates a user's phone number in the database
func (DBC *DBConnector) UpdatePhone(userID, phone string) error {
	return DBC.db.Model(&entity.User{}).Where("user_id = ?", userID).Update("phone", phone).Error
}

// ADMIN (Needs guardrails)

// GetUsers retrieves all users from the database
func (DBC *DBConnector) GetUsers() ([]entity.User, error) {
	var users []entity.User
	if err := DBC.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// DeleteAllUsers deletes all users from the database
func (DBC *DBConnector) DeleteAllUsers() error {
	return DBC.db.Delete(&entity.User{}).Error
}

// GetUsersByRole retrieves all users from the database for a given role
func (DBC *DBConnector) GetUsersByRole(role string) ([]entity.User, error) {
	var users []entity.User
	if err := DBC.db.Joins("JOIN authentications ON authentications.user_id = users.user_id").
		Where("authentications.role = ?", role).
		Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
