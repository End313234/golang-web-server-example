package utils

import "github.com/gofiber/fiber/v2"

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewError(ctx *fiber.Ctx, code int, message string) error {
	err := Error{
		Message: message,
		Code:    code,
	}

	return ctx.Status(code).JSON(err)
}
