package models

type Role struct {
	ID       uint   `gorm:"primaryKey"`
	RoleName string `gorm:"unique;not null"`
}
