package controllers

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"
)

type RoleController struct {
	DB *sql.DB
}

func (c *RoleController) Index(ctx *fiber.Ctx) error {
	var roles []models.Role

	database.DB.Find(&roles)

	return ctx.JSON(roles)
}

func (c *RoleController) CreateRole(ctx *fiber.Ctx) error {
	var roleDto fiber.Map

	if err := ctx.BodyParser(&roleDto); err != nil {
		return err
	}

	list := roleDto["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))

		permissions[i] = models.Permission{
			ID: uint(id),
		}
	}

	role := models.Role{
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}

	database.DB.Create(&role)

	return ctx.JSON(role)
}

func (c *RoleController) GetRole(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	role := models.Role{
		ID: uint(id),
	}

	database.DB.Preload("Permissions").Find(&role)

	return ctx.JSON(role)
}

func (c *RoleController) UpdateRole(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	var roleDto fiber.Map

	if err := ctx.BodyParser(&roleDto); err != nil {
		return err
	}

	list := roleDto["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := permissionId.(float64)

		permissions[i] = models.Permission{
			ID: uint(id),
		}
	}

	var result interface{}

	database.DB.Table("role_permissions").Where("role_id", id).Delete(&result)

	role := models.Role{
		ID:          uint(id),
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}

	database.DB.Model(&role).Updates(role)

	return ctx.JSON(role)
}

func (c *RoleController) DeleteRole(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	role := models.Role{
		ID: uint(id),
	}

	database.DB.Delete(&role)

	return nil
}
