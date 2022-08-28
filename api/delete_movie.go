package api

import (
	"web-server-example/utils"

	"github.com/gofiber/fiber/v2"
)

type DeleteMovieRequest utils.Movie

func DeleteMovie(ctx *fiber.Ctx) error {
	var target DeleteMovieRequest
	utils.ReadBody(ctx.Body(), &target)

	// Need specific conversion since "AddMovieRequest" is
	// not the same thing as "utils.Movie"
	_, err := utils.MoviesDatabase.Delete(utils.Movie(target))
	if err != nil {
		return utils.NewError(ctx, fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(target)
}
