package models

import "time"

type PopUpDay struct {
	ID         uint      `gorm:"primaryKey"`
	VenueID    uint      `gorm:"not null"`
	Date       time.Time `gorm:"not null"`
	Volunteers string    `gorm:"type:json;not null"`
	Dogs       string    `gorm:"type:json;not null"`
	Venue      Venue     `gorm:"foreignKey:VenueID"`
}
