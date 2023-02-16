package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func errorResponse(ctx *fiber.Ctx, err error) error {
	log.Println(err)
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"success": false,
		"data":    err,
	})
}
