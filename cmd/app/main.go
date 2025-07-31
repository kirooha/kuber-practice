package main

import (
	"log"

	"github.com/kirooha/kuber-practice/internal/app/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/files", handlers.ListFilesHandler)
	app.Post("/file", handlers.SaveFileHandler)

	log.Fatal(app.Listen(":8080"))
}
