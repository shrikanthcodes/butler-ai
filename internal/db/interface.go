package db

import "gorm.io/gorm"

type DBConnector struct {
	db *gorm.DB
}

func NewDBConnector(db *gorm.DB) *DBConnector {
	return &DBConnector{
		db: db,
	}
}
