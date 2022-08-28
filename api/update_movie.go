package api

import (
	"web-server-example/utils"

	"github.com/gofiber/fiber/v2"
)

type UpdateMovieRequest struct {
	Old utils.Movie `json:"old"`
	New utils.Movie `json:"new"`
}

func UpdateMovie(ctx *fiber.Ctx) error {
	var data UpdateMovieRequest
	utils.ReadBody(ctx.Body(), &data)

	// Need specific conversion since "AddMovieRequest" is
	// not the same thing as "utils.Movie"
	_, err := utils.MoviesDatabase.Update(data.Old, data.New)
	if err != nil {
		return utils.NewError(ctx, fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(data.New)
}
