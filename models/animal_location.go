package models

type AnimalLocation struct {
	ID       uint   `gorm:"primaryKey"`
	Location string `gorm:"unique;not null"`
}
