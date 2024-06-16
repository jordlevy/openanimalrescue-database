package models

import "time"

type User struct {
	ID               uint   `gorm:"primaryKey"`
	CognitoID        string `gorm:"unique;not null"`
	Name             string `gorm:"not null"`
	Surname          string `gorm:"not null"`
	Address          string
	ContactInfo      string
	RegistrationDate time.Time `gorm:"not null"`
	Roles            []Role    `gorm:"many2many:user_roles"`
}
