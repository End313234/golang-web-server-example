package api

import (
	"web-server-example/utils"

	"github.com/gofiber/fiber/v2"
)

type GetMovieRequest utils.Movie

func GetMovie(ctx *fiber.Ctx) error {
	var target GetMovieRequest
	utils.ReadBody(ctx.Body(), &target)

	// Need specific conversion since "AddMovieRequest" is
	// not the same thing as "utils.Movie"
	_, err := utils.MoviesDatabase.Get(utils.Movie(target))
	if err != nil {
		return utils.NewError(ctx, fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(target)
}
