package controllers

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"
)

type CategorieController struct {
	DB *sql.DB
}

func (c *CategorieController) Index(ctx *fiber.Ctx) error {
	var categories []models.Categorie
	database.DB.Order("id asc").Find(&categories)

	return ctx.JSON(categories)
	// page, _ := strconv.Atoi(ctx.Query("page", "1"))

	// return ctx.JSON(models.Paginate(database.DB, &models.Categorie{}, page))

}

func (c *CategorieController) CreateCategorie(ctx *fiber.Ctx) error {
	var categorie models.Categorie

	if err := ctx.BodyParser(&categorie); err != nil {
		return err
	}

	database.DB.Create(&categorie)

	return ctx.JSON(categorie)
}

func (c *CategorieController) GetCategorie(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	categorie := models.Categorie{
		ID: uint(id),
	}

	database.DB.Find(&categorie)

	return ctx.JSON(categorie)
}

func (c *CategorieController) UpdateCategories(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	categorie := models.Categorie{
		ID: uint(id),
	}

	if err := ctx.BodyParser(&categorie); err != nil {
		return err
	}

	database.DB.Model(&categorie).Updates(categorie)

	return ctx.JSON(categorie)
}

func (c *CategorieController) DeleteCategorie(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	categorie := models.Categorie{
		ID: uint(id),
	}

	database.DB.Delete(&categorie)

	return nil
}
