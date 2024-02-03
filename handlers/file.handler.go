package handlers

import (
	"strconv"

	"github.com/mesxx/Fiber_Simple_File_Management_API/helpers"
	"github.com/mesxx/Fiber_Simple_File_Management_API/models"
	"github.com/mesxx/Fiber_Simple_File_Management_API/usecases"

	"github.com/gofiber/fiber/v2"
)

type (
	FileHandler interface {
		Upload(c *fiber.Ctx) error
		GetAll(c *fiber.Ctx) error
		RenderFileByID(c *fiber.Ctx) error
		Delete(c *fiber.Ctx) error
	}

	fileHandler struct {
		FileUsecase usecases.FileUsecase
	}
)

func NewFileHandler(fu usecases.FileUsecase) FileHandler {
	return &fileHandler{
		FileUsecase: fu,
	}
}

func (fh fileHandler) Upload(c *fiber.Ctx) error {
	var fileRequest models.File
	file, err := c.FormFile("image")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	//
	fileType := file.Header.Get("Content-Type")
	if err := helpers.UploadSettingType(fileType); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	fileName, err := helpers.UploadSettingName(file.Filename)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	//

	//
	destination := "./public/images/" + fileName
	if err := c.SaveFile(file, destination); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	//

	fileRequest.Name = fileName
	res, err := fh.FileUsecase.Create(&fileRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success upload",
		"data":    res,
	})
}

func (fh fileHandler) GetAll(c *fiber.Ctx) error {
	files, err := fh.FileUsecase.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    files,
	})
}

func (fh fileHandler) RenderFileByID(c *fiber.Ctx) error {
	id := c.Params("id")

	value, err := strconv.Atoi(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	file, err := fh.FileUsecase.GetByID(uint(value))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	//
	filename := file.Name
	destination := "./public/images/" + filename
	//

	return c.SendFile(destination)
}

func (fh fileHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	value, err := strconv.Atoi(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	fileDeleted, err := fh.FileUsecase.Delete(uint(value))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success delete",
		"data":    fileDeleted,
	})
}
