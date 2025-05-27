package model

import "time"

type Device struct {
	ID         string    `gorm:"primaryKey"`
	UserID     string    `gorm:"index;not null"`
	Name       string    `gorm:"not null"`
	IP         string    `gorm:"not null"`
	OS         string
	Status     string    // online / offline
	LastSeen   time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
