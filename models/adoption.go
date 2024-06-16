package models

import "time"

type Adoption struct {
	ID           uint      `gorm:"primaryKey"`
	AnimalID     uint      `gorm:"not null"`
	AdopterID    uint      `gorm:"not null"`
	AdoptionDate time.Time `gorm:"not null"`
	StatusID     uint      `gorm:"not null"`
	Outcome      string
	Animal       Animal         `gorm:"foreignKey:AnimalID"`
	Adopter      User           `gorm:"foreignKey:AdopterID"`
	Status       AdoptionStatus `gorm:"foreignKey:StatusID"`
}
