package models

import "time"

type Animal struct {
	ID                    uint   `gorm:"primaryKey"`
	Name                  string `gorm:"not null"`
	Species               string `gorm:"not null"`
	Breed                 string
	Age                   int
	Sex                   string
	Description           string
	ArrivalDate           time.Time `gorm:"not null"`
	HealthStatusID        uint
	SterilisationStatusID uint
	ChipNumber            string
	InternalNotes         string
	ReasonOnboarded       string
	LatestVaccinationDate time.Time
	CurrentLocationID     uint
	Status                string                    `gorm:"not null"`
	HealthStatus          AnimalHealthStatus        `gorm:"foreignKey:HealthStatusID"`
	SterilisationStatus   AnimalSterilisationStatus `gorm:"foreignKey:SterilisationStatusID"`
	CurrentLocation       AnimalLocation            `gorm:"foreignKey:CurrentLocationID"`
}
