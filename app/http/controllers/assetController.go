package controllers

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"
)

type AssetController struct {
	DB *sql.DB
}
type CreateAssetReq struct {
	ID           uint      `json:"id"`
	Name         string    `json:"assigned_to"`
	Title        string    `json:"title"`
	SerialNumber string    `json:"serialnumber"`
	Description  string    `json:"description"`
	Price        string    `json:"price"`
	PurchaseDate time.Time `json:"date_purchased"`
	DateAssigned time.Time `json:"date_assigned"`
	CategorieID  uint      `json:"categorie_id"`
}

func (c *AssetController) Index(ctx *fiber.Ctx) error {

	var assets []models.Asset
	database.DB.Preload("Categories").Order("id asc").Find(&assets)

	return ctx.JSON(assets)
	// page, _ := strconv.Atoi(ctx.Query("page", "1"))

	// return ctx.JSON(models.Paginate(database.DB, &models.Asset{}, page))

}

func (c *AssetController) CreateAsset(ctx *fiber.Ctx) error {

	var assetReq CreateAssetReq

	if err := ctx.BodyParser(&assetReq); err != nil {
		return err
	}
	categorie := models.Categorie{
		ID: assetReq.CategorieID,
	}
	database.DB.Find(&categorie)

	asset := models.Asset{
		ID:           assetReq.ID,
		Title:        assetReq.Title,
		Description:  assetReq.Description,
		SerialNumber: assetReq.SerialNumber,
		PurchaseDate: assetReq.PurchaseDate,
		DateAssigned: assetReq.DateAssigned,
		Price:        assetReq.Price,

		// ImageType:    assetReq.ImageType,
		// ImageUrl:     assetReq.ImageUrl,
	}
	// asset.ImageID = image.ID

	database.DB.Create(&asset)

	database.DB.Model(&asset).Association("Categories").Append(&categorie)

	return ctx.JSON(asset)
}

func (c *AssetController) GetAsset(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	asset := models.Asset{
		ID: uint(id),
	}

	database.DB.Preload("Categories").Find(&asset)

	return ctx.JSON(asset)
}

func (c *AssetController) UpdateAsset(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("assetId"))

	asset := models.Asset{
		ID: uint(id),
	}
	var user models.User

	if err := ctx.BodyParser(&asset); err != nil {
		return err
	}
	if asset.IsClearedOf {
		asset.IsClearedOf = true
		asset.IsAssigned = false
		database.DB.Model(&asset).Update("is_assigned", false)
		database.DB.Model(&asset).Update("assigned_to", "")

	}
	if asset.IsDamaged {
		asset.IsDamaged = true
		database.DB.Model(&asset).Update("is_assigned", false)
		database.DB.Model(&asset).Update("assigned_to", "")
	}
	if user.NotActive {
		asset.AssignedTo = ""
		asset.IsAssigned = false

	}

	database.DB.Model(&asset).Updates(asset)

	return ctx.JSON(asset)

}

func (c *AssetController) DeleteAsset(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	asset := models.Asset{
		ID: uint(id),
	}

	database.DB.Delete(&asset)

	return nil
}
