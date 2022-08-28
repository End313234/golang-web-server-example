package api

import (
	"web-server-example/utils"

	"github.com/gofiber/fiber/v2"
)

type AddMovieRequest utils.Movie

func AddMovie(ctx *fiber.Ctx) error {
	var newMovie AddMovieRequest
	utils.ReadBody(ctx.Body(), &newMovie)

	// Need specific conversion since "AddMovieRequest" is
	// not the same thing as "utils.Movie"
	utils.MoviesDatabase.Add(utils.Movie(newMovie))

	return ctx.Status(fiber.StatusOK).JSON(newMovie)
}
