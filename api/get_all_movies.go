package api

import (
	"web-server-example/utils"

	"github.com/gofiber/fiber/v2"
)

func GetAllMovies(ctx *fiber.Ctx) error {
	movies := utils.MoviesDatabase.GetAll()

	return ctx.Status(fiber.StatusOK).JSON(movies)
}
