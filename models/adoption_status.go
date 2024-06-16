package models

type AdoptionStatus struct {
	ID     uint   `gorm:"primaryKey"`
	Status string `gorm:"unique;not null"`
}
