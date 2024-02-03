package main

import (
	"os"

	"github.com/mesxx/Fiber_Simple_File_Management_API/servers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	port := os.Getenv("PORT")
	server := servers.Server()
	if err := server.Listen(":" + port); err != nil {
		fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
}
