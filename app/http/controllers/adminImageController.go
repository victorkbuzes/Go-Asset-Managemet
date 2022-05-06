package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"
)

type AdminImageController struct {
	DB *sql.DB
}

func (c *AdminImageController) Index(ctx *fiber.Ctx) error {
	var image []models.Image
	database.DB.Find(&image)

	return ctx.JSON(image)
	// page, _ := strconv.Atoi(ctx.Query("page", "1"))

	// return ctx.JSON(models.Paginate(database.DB, &models.Image{}, page))
}

//FUNCTIONALITY TO CREATE IMAGE

func (ic *AdminImageController) UploadImage(ctx *fiber.Ctx) error {

	//parse incoming image file
	var image models.ImagePost
	file, err := ctx.FormFile("image")
	if err := ctx.BodyParser(&image); err != nil {
		return err
	}

	name := image.Name

	if err != nil {
		log.Println("image upload error --> ", err)
		return ctx.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	// generate new uuid for image name
	uniqueId := uuid.New()

	// remove "- from imageName"

	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	// extract image extension from original file filename

	fileExt := strings.Split(file.Filename, ".")[1]

	// generate image from filename and extension
	image_name := fmt.Sprintf("%s.%s", filename, fileExt)

	// save image to ./images dir
	err = ctx.SaveFile(file, fmt.Sprintf("./resources/uploads/%s", image_name))

	if err != nil {
		log.Println("image save error --> ", err)
		return ctx.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	// generate image url to serve to client using CDN

	imageUrl := fmt.Sprintf("http://localhost:8000/resources/uploads/%s", image_name)

	// create meta data and send to client

	data := &models.Image{
		Image:     image_name,
		ImageUrl:  imageUrl,
		ImageType: name,
	}

	database.DB.Create(data)

	return ctx.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})

}

//FUNCTIONALITY TO DELETE IMAGE

func (ic *AdminImageController) DeleteImage(ctx *fiber.Ctx) error {

	// extract image name from params
	imageName := ctx.Params("image_name")

	// delete image from ./images
	err := os.Remove(fmt.Sprintf("./resources/uploads/%s", imageName))
	if err != nil {
		log.Println(err)
		return ctx.JSON(fiber.Map{"status": 500, "message": "Server Error", "data": nil})
	}

	return ctx.JSON(fiber.Map{"status": 201, "message": "Image deleted successfully", "data": nil})
}
