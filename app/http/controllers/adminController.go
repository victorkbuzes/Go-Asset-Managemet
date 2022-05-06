package controllers

import (
	"strconv"

	"database/sql"

	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/http/middlewares"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"
)

type AdminController struct {
	DB *sql.DB
}

func (c *AdminController) Index(ctx *fiber.Ctx) error {
	var admin []models.Admin
	database.DB.Order("id asc").Find(&admin)

	return ctx.JSON(admin)
	// if err := middlewares.IsAuthenticated(ctx); err != nil {
	//         return err
	// }

	// page, _ := strconv.Atoi(ctx.Query("page", "1"))

	// return ctx.JSON(models.Paginate(database.DB, &models.Admin{}, page))
}

func (c *AdminController) CreateAdmin(ctx *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(ctx, "admin"); err != nil {
		return err
	}
	if err := middlewares.IsAuthenticated(ctx); err != nil {
		return err
	}

	var admin models.Admin

	if err := ctx.BodyParser(&admin); err != nil {
		return err
	}

	database.DB.Create(&admin)

	return ctx.JSON(admin)

}

func (c *AdminController) GetAdmin(ctx *fiber.Ctx) error {
	// if err := middlewares.IsAuthenticated(ctx); err != nil {
	// 	return err
	// }

	id, _ := strconv.Atoi(ctx.Params("name"))

	admin := models.Admin{
		ID: uint(id),
	}

	database.DB.Preload("Role").Find(&admin)

	return ctx.JSON(admin)
}

func (c *AdminController) UpdateAdmin(ctx *fiber.Ctx) error {
	if err := middlewares.IsAuthenticated(ctx); err != nil {
		return err
	}

	id, _ := strconv.Atoi(ctx.Params("id"))

	admin := models.Admin{
		ID: uint(id),
	}

	if err := ctx.BodyParser(&admin); err != nil {
		return err
	}

	database.DB.Model(&admin).Updates(admin)

	return ctx.JSON(admin)
}

func (c *AdminController) DeleteAdmin(ctx *fiber.Ctx) error {
	if err := middlewares.IsAuthenticated(ctx); err != nil {
		return err
	}
	id, _ := strconv.Atoi(ctx.Params("id"))

	admin := models.Admin{
		ID: uint(id),
	}

	database.DB.Delete(&admin)

	return nil
}
