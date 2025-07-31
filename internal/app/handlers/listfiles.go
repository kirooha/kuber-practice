package handlers

import "github.com/gofiber/fiber/v2"

func ListFilesHandler(fiberCtx *fiber.Ctx) error {
	return fiberCtx.SendString("list of files")
}
