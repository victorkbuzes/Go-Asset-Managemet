package models

import (
	"time"
)

type User struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	Email      string    `json:"email" gorm:"unique"`
	IsActive   bool      `json:"is_active" gorm:"default:true"`
	NotActive  bool      `json:"not_active" gorm:"default:false"`
	Department string    `json:"department"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	// Assets     []Asset   `json:"assets" gorm:"many2many:user_assets;"`

	// Tags        []Tag        `json:"tag" gorm:"many2many:user_tags;"`
	Accessories []Accessorie `json:"accessories" gorm:"many2many:user_accessories;"`
	Departments []Department `json:"departments" gorm:"many2many:user_department;"`
}

// func (user *User) Count(db *gorm.DB) int64 {
// 	var total int64
// 	db.Model(&User{}).Count(&total)

// 	return total
// }

// //Relationship btwn a User and Asset
// func (user *User) Take(db *gorm.DB, limit int, offset int) interface{} {
// 	var users []User

// 	db.Preload("Assets").Preload("Accesories").Preload("Departments").Offset(offset).Limit(limit).Find(&users)

// 	return users
// }
