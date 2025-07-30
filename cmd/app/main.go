package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func listFilesHandler(fiberCtx *fiber.Ctx) error {
	return fiberCtx.SendString("list of files")
}

func main() {
	app := fiber.New()

	app.Get("/files", listFilesHandler)

	log.Fatal(app.Listen(":8080"))
}
