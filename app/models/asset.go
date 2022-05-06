package models

import (
	"time"

	"gorm.io/gorm"
)

type Asset struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	AssignedTo     string    `json:"assigned_to"`
	Title          string    `json:"title"`
	SerialNumber   string    `json:"serialnumber"`
	Description    string    `json:"description"`
	Price          string    `json:"price"`
	DateAssigned   time.Time `json:"date_assigned"`
	PurchaseDate   time.Time `json:"purchase_date"`
	CategorieTitle string    `json:"categorie_tile"`
	IsAssigned     bool      `json:"is_assigned" gorm:"default:true"`
	IsClearedOf    bool      `json:"is_cleared_of" gorm:"default:false"`
	IsDamaged      bool      `json:"is_damaged" gorm:"default:false"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	Created string `json:"created"`

	Categories []Categorie `json:"categorie" gorm:"many2many:asset_categorie;"`
}

func (asset *Asset) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Asset{}).Count(&total)

	return total
}

func (asset *Asset) Take(db *gorm.DB, limit int, offset int) interface{} {
	var assets []Asset
	db.Preload("Categories").Offset(offset).Limit(limit).Find(&assets)

	return assets
}
