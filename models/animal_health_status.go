package models

type AnimalHealthStatus struct {
	ID     uint   `gorm:"primaryKey"`
	Status string `gorm:"unique;not null"`
}
