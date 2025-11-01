package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string    `gorm:"uniqueIndex;not null"`
	Password string    `gorm:"not null"`
	Websites []Website `gorm:"constraint:OnDelete:CASCADE;"`
	Session  Session   `gorm:"constraint:OnDelete:CASCADE;"`
}

type Session struct {
	gorm.Model
	SessionID  uuid.UUID `gorm:"uniqueIndex"`
	SecretHash string
	UserID     uint `gorm:"uniqueIndex"`
}

type Website struct {
	gorm.Model
	Url           string `gorm:"not null"`
	UserID        uint
	Website_Ticks []WebsiteTick `gorm:"constraint:OnDelete:CASCADE;"`
}

type Region struct {
	gorm.Model
	Name         string        `gorm:"not null"`
	Region_Ticks []WebsiteTick `gorm:"constraint:OnDelete:CASCADE;"`
}

type WebsiteTick struct {
	gorm.Model
	ResponseTime int
	Status       string

	WebsiteID uint
	RegionID  uint
}
