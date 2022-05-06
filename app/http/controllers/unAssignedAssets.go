package controllers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"
)

type UnAssignedAssetController struct {
	DB *sql.DB
}

func (c *UnAssignedAssetController) Index(ctx *fiber.Ctx) error {

	var asset []models.Asset
	// Get first matched record
	database.DB.Order("id asc").Where("is_assigned = ?", false).Where("is_cleared_of = ?", false).Where("is_damaged = ?", false).Find(&asset)

	return ctx.JSON(&asset)
}
