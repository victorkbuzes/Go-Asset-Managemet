package controllers

import (
	"fmt"
	"strconv"

	"database/sql"

	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"
)

type createUserReq struct {
	ID           uint   `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Assets       string `json:"assets"`
	Accessories  string `json:"accessories"`
	AssetID      uint   `json:"asset_id"`
	IsActive     bool   `json:"is_active" gorm:"default:true"`
	AccesorieID  uint   `json:"accesorie_id" `
	DepartmentID uint   `json:"department_id"`
}

type UserController struct {
	DB *sql.DB
}

func (c *UserController) Index(ctx *fiber.Ctx) error {
	var users []models.User
	database.DB.Preload("Assets").Preload("Accessories").Preload("Departments").Order("id asc").Find(&users)

	return ctx.JSON(users)

}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {

	var userReq createUserReq

	if err := ctx.BodyParser(&userReq); err != nil {
		return err

	}
	// fmt.Printf(" Userreq %v userreq \n", userReq)

	//asset
	asset := models.Asset{
		ID: userReq.AssetID,
	}
	database.DB.Where("is_assigned = ?", false).Where("is_cleared_of = ?", false).Where("is_damaged = ?", false).Find(&asset)

	// fmt.Printf("log asset %v", asset)
	asset.AssignedTo = userReq.Name

	//accesorie
	var acccesorie models.Accessorie

	if userReq.AccesorieID > 0 {
		acccesorie = models.Accessorie{
			ID: userReq.AccesorieID,
		}

	}
	fmt.Print("log accesorie id ", userReq.AccesorieID)

	database.DB.Where("is_assigned = ?", false).Where("is_cleared_of = ?", false).Where("is_damaged = ?", false).Find(&acccesorie)

	acccesorie.AssignedTo = userReq.Name

	// fmt.Printf("log accesore %v", acccesorie)

	// department
	department := models.Department{
		ID: userReq.DepartmentID,
	}
	database.DB.Find(&department)
	// fmt.Printf("log asset %v", department)

	user := models.User{
		ID:       userReq.ID,
		Name:     userReq.Name,
		Email:    userReq.Email,
		IsActive: userReq.IsActive,
	}

	// Update with conditions and model value

	//assigned_to
	database.DB.Model(&asset).Update("assigned_to", userReq.Name)
	database.DB.Model(&acccesorie).Update("assigned_to", userReq.Name)
	//asset
	asset.IsAssigned = true
	assetresult := database.DB.Model(&asset).Where("id = ?", userReq.AssetID).Update("is_assigned", true)
	if assetresult.Error != nil {
		fmt.Printf("Error in updating %v", assetresult.Error)
		//return result.Error
	}
	// fmt.Printf(" asset after update %v", asset)

	//accesorie update

	acccesorie.IsAssigned = true
	accesorieresult := database.DB.Model(&acccesorie).Where("id = ?",
		userReq.AccesorieID).Update("is_assigned", true)
	if accesorieresult.Error != nil {
		fmt.Printf("Error in updating %v", accesorieresult.Error)
		//return result.Error
	}
	fmt.Printf(" acessorie after update %v", acccesorie)
	database.DB.Create(&user)

	database.DB.Model(&user).Association("Assets").Append(&asset)
	// database.DB.Model(&user).Association("Tags").Append(&tag)
	if userReq.AccesorieID > 0 {
		database.DB.Model(&user).Association("Accesories").Append(&acccesorie)
	}

	database.DB.Model(&user).Association("Departments").Append(&department)

	return ctx.JSON(&user)
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	user := models.User{
		ID: uint(id),
	}

	database.DB.Preload("Assets").Preload("Departments").Preload("Accesories").Find(&user)
	return ctx.JSON(user)
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))
	// var asset models.Asset
	// var acccesorie models.Accesorie

	user := models.User{
		ID: uint(id),
	}

	if err := ctx.BodyParser(&user); err != nil {
		return err
	}
	// if user.NotActive {
	// 	database.DB.Model(&asset).Update("is_assigned", false)

	// }

	database.DB.Model(&user).Preload("Department").Preload("Asset").Updates(user)

	return ctx.JSON(user)
}

func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {
	// if err := middlewares.IsAuthenticated(ctx); err != nil {
	// 	return err
	// }
	id, _ := strconv.Atoi(ctx.Params("id"))

	user := models.User{
		ID: uint(id),
	}

	database.DB.Delete(&user)

	return nil
}
