package controllers

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"
)

type AccesorieCreateReq struct {
	ID           uint      `json:"id"`
	AssignedTo   string    `json:"name"`
	Title        string    `json:"title"`
	SerialNumber string    `json:"serial_number"`
	Description  string    `json:"decripion"`
	Price        string    `json:"price"`
	DateAssigned time.Time `json:"date_assigned"`
	DateReturned time.Time `json:"date_returned"`
	PurchaseDate time.Time `json:"purchase_date"`
	IsAssigned   bool      `json:"is_assigned" gorm:"default:false"`
	IsClearedOf  bool      `json:"is_cleared_of" gorm:"default:false"`
	IsDamaged    bool      `json:"is_damaged" gorm:"default:false"`
	Reason       string    `json:"reason"`
	CategorieID  uint      `json:"categorie_id"`
}

type AccesorieController struct {
	DB *sql.DB
}

// 0722157344

func (c *AccesorieController) Index(ctx *fiber.Ctx) error {

	var accessories []models.Accessorie
	database.DB.Preload("Categories").Order("id asc").Find(&accessories)

	return ctx.JSON(accessories)
}

func (c *AccesorieController) CreateAccesorie(ctx *fiber.Ctx) error {
	var accessorieReq AccesorieCreateReq

	if err := ctx.BodyParser(&accessorieReq); err != nil {
		return err
	}

	//categorie
	categorie := models.Categorie{
		ID: accessorieReq.CategorieID,
	}
	database.DB.Find(&categorie)

	accessorie := models.Accessorie{
		ID:           accessorieReq.ID,
		Title:        accessorieReq.Title,
		SerialNumber: accessorieReq.SerialNumber,
		Description:  accessorieReq.Description,
		DateAssigned: accessorieReq.DateAssigned,
		DateReturned: accessorieReq.DateReturned,
		Price:        accessorieReq.Price,
		PurchaseDate: accessorieReq.PurchaseDate,
	}

	database.DB.Create(&accessorie)

	database.DB.Model(&accessorie).Association("Categories").Append(&categorie)

	return ctx.JSON(accessorie)
}

func (c *AccesorieController) GetAccesorie(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	accessorie := models.Accessorie{
		ID: uint(id),
	}

	database.DB.Preload("Categories").Find(&accessorie)

	return ctx.JSON(accessorie)
}

func (c *AccesorieController) UpdateAccesorie(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("accessorieId"))

	accessorie := models.Accessorie{
		ID: uint(id),
	}

	if err := ctx.BodyParser(&accessorie); err != nil {
		return err
	}

	if accessorie.IsDamaged {
		accessorie.IsDamaged = true
		database.DB.Model(&accessorie).Update("is_assigned", false)
		database.DB.Model(&accessorie).Update("assigned_to", "")
	}
	if accessorie.IsClearedOf {
		accessorie.IsClearedOf = true
		accessorie.IsAssigned = false
		database.DB.Model(&accessorie).Update("is_assigned", false)
		database.DB.Model(&accessorie).Update("assigned_to", "")
	}

	database.DB.Model(&accessorie).Updates(accessorie)

	return ctx.JSON(accessorie)

}

func (c *AccesorieController) DeleteAccesorie(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	acccesorie := models.Accessorie{
		ID: uint(id),
	}

	database.DB.Delete(&acccesorie)

	return nil
}
