package models

import (
	"time"
)

type Migration struct {
	ID          uint      `gorm:"primaryKey"`
	MigrationID string    `gorm:"unique;not null"`
	AppliedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
}
