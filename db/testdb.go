package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewTestDB() (*gorm.DB, error) {
	database, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = database.AutoMigrate(&User{}, &Session{}, &Website{}, &Region{}, &WebsiteTick{})
	if err != nil {
		return nil, err
	}

	return database, nil
}
