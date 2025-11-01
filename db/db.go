package db

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Gorm struct {
	Db *gorm.DB
}

func NewGorm() (*Gorm, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		),
	})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&Session{}, &User{}, &Region{}, &Website{}, &WebsiteTick{})
	if err != nil {
		return nil, err
	}
	g := &Gorm{
		Db: db,
	}
	return g, nil
}
