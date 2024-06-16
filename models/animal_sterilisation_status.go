package models

type AnimalSterilisationStatus struct {
	ID     uint   `gorm:"primaryKey"`
	Status string `gorm:"unique;not null"`
}
