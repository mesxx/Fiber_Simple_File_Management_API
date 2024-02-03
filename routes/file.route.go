package routes

import (
	"github.com/mesxx/Fiber_Simple_File_Management_API/handlers"
	"github.com/mesxx/Fiber_Simple_File_Management_API/repositories"
	"github.com/mesxx/Fiber_Simple_File_Management_API/usecases"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func FileRouter(app fiber.Router, db *gorm.DB) {
	fr := repositories.NewFileRepository(db)
	fu := usecases.NewFileUsecase(fr)
	fh := handlers.NewFileHandler(fu)

	app.Post("/", fh.Upload)
	app.Get("/", fh.GetAll)
	app.Get("/:id", fh.RenderFileByID)
	app.Delete("/:id", fh.Delete)
}
