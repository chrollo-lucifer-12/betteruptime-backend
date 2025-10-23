package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm struct {
	Db *gorm.DB
}

func NewGorm() (*Gorm, error) {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate()
	if err != nil {
		return nil, err
	}
	g := &Gorm{
		Db: db,
	}
	return g, nil
}
