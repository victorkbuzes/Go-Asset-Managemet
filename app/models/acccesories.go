package models

import (
	"time"
)

type Accessorie struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	AssignedTo string `json:"name"`
	// Type            string    `json:"type"`
	Title        string    `json:"title"`
	SerialNumber string    `json:"serialnumber"`
	Description  string    `json:"description"`
	Price        string    `json:"price"`
	PurchaseDate time.Time `json:"purchase_date"`
	DateAssigned time.Time `json:"date_assigned"`
	DateReturned time.Time `json:"date_returned"`
	Reason       string    `json:"reason"`
	IsAssigned   bool      `json:"is_assigned" gorm:"default:true"`
	IsClearedOf  bool      `json:"is_cleared_of" gorm:"default:false"`
	IsDamaged    bool      `json:"is_damaged" gorm:"default:false"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Categories []Categorie `json:"categories" gorm:"many2many:accessorie_categorie;"`
}
