package controllers

import (
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"

	"strconv"
	"time"

	"gitlab.ci.emalify.com/roamtech/asset_be/util"

	"github.com/gofiber/fiber/v2"
)

////REGISTER
func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//registration validations
	if data["first_name"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": " first name is required!",
		})
	}
	if data["last_name"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "Lastname is required!",
		})
	}
	if data["email"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "Email is required!",
		})
	}
	if data["password"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "Password is required!",
		})
	}
	if len(data["password"]) <= 6 {
		return c.Status(500).JSON(fiber.Map{
			"message": "Password should be more then 6 chars",
		})
	}
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match",
		})
	}

	admin := models.Admin{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	admin.SetPassword(data["password"])

	database.DB.Create(&admin)

	return c.JSON(admin)
}

////LOGIN
func Login(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//Login validations
	if data["email"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "Email is required!",
		})
	}
	if data["password"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "Password is required!",
		})
	}
	if len(data["password"]) <= 6 {
		return c.Status(500).JSON(fiber.Map{
			"message": "Password should be more then 6 chars",
		})
	}

	var admin models.Admin

	database.DB.Where("email = ?", data["email"]).First(&admin)

	if admin.ID == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Invalid Email or Password",
		})
	}

	if err := admin.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid Password",
		})
	}

	token, err := util.GenerateJwt(strconv.Itoa(int(admin.ID)))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
		"token":   token,
	})
}

func Admin(c *fiber.Ctx) error {
	cookie := c.Get("authorization")

	id, _ := util.ParseJwt(cookie)

	var admin models.Admin

	database.DB.Where("id = ?", id).First(&admin)

	return c.JSON(admin)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	cookie := c.Get("authorization")

	id, _ := util.ParseJwt(cookie)

	userId, _ := strconv.Atoi(id)

	admin := models.Admin{
		ID:        uint(userId),
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	database.DB.Model(&admin).Updates(admin)

	return c.JSON(admin)
}

func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match",
		})
	}

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	adminId, _ := strconv.Atoi(id)

	admin := models.Admin{
		ID: uint(adminId),
	}

	admin.SetPassword(data["password"])

	database.DB.Model(&admin).Updates(admin)

	return c.JSON(admin)
}
