package models

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ImageType string    `json:"image_type"`
	Image     string    `json:"image_name"`
	ImageName string    `json:"name"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ImagePost struct {
	Name string `json:"name"`
}

func (image *Image) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Image{}).Count(&total)

	return total
}

func (image *Image) Take(db *gorm.DB, limit int, offset int) interface{} {
	var images []Image

	db.Offset(offset).Limit(limit).Find(&images)

	return images
}
