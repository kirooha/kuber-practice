package handlers

import "github.com/gofiber/fiber/v2"

func SaveFileHandler(fiberCtx *fiber.Ctx) error {
	return fiberCtx.SendString("save file")
}
