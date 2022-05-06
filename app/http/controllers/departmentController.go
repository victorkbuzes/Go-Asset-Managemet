package controllers

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"
)

type DepartmentController struct {
	DB *sql.DB
}

func (c *DepartmentController) Index(ctx *fiber.Ctx) error {
	var departments []models.Department
	database.DB.Order("id asc").Find(&departments)

	return ctx.JSON(departments)

	// page, _ := strconv.Atoi(ctx.Query("page", "1"))

	// return ctx.JSON(models.Paginate(database.DB, &models.Department{}, page))

}

func (c *DepartmentController) CreateDepartment(ctx *fiber.Ctx) error {
	var department models.Department

	if err := ctx.BodyParser(&department); err != nil {
		return err
	}

	database.DB.Create(&department)

	return ctx.JSON(department)
}

func (c *DepartmentController) GetDepartment(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	department := models.Department{
		ID: uint(id),
	}

	database.DB.Find(&department)

	return ctx.JSON(department)
}

func (c *DepartmentController) UpdateDepartment(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	department := models.Department{
		ID: uint(id),
	}

	if err := ctx.BodyParser(&department); err != nil {
		return err
	}

	database.DB.Model(&department).Updates(department)

	return ctx.JSON(department)
}

func (c *DepartmentController) DeleteDepartment(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	department := models.Department{
		ID: uint(id),
	}

	database.DB.Delete(&department)

	return nil
}
