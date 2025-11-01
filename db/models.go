package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string    `gorm:"uniqueIndex;not null"`
	Password string    `gorm:"not null"`
	Websites []Website `gorm:"constraint:OnDelete:CASCADE;"`
	Session  []Session
}

type Session struct {
	gorm.Model
	SecretHash string
	UserID     uint
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
