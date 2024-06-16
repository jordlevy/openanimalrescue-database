package models

type Venue struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"not null"`
	Address        string `gorm:"not null"`
	ContactInfo    string
	GoogleMapsLink string
	OperatingHours string
}
