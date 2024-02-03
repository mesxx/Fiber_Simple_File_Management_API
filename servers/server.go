package servers

import (
	"github.com/mesxx/Fiber_Simple_File_Management_API/configs"
	"github.com/mesxx/Fiber_Simple_File_Management_API/middlewares"
	"github.com/mesxx/Fiber_Simple_File_Management_API/models"
	"github.com/mesxx/Fiber_Simple_File_Management_API/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Server() *fiber.App {
	db, err := configs.DatabaseConfig()
	if err != nil {
		fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	db.AutoMigrate(&models.File{})

	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorMiddleware,
	})
	app.Use(logger.New())
	app.Use(recover.New())

	routes.FileRouter(app, db)

	return app
}
