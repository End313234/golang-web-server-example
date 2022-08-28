package router

import (
	"web-server-example/api"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	router := fiber.New()

	app.Add("GET", "/", api.GetMovie)
	app.Add("GET", "/all", api.GetAllMovies)
	app.Add("POST", "/", api.AddMovie)
	app.Add("PUT", "/", api.UpdateMovie)
	app.Add("DELETE", "/", api.DeleteMovie)

	app.Mount("/", router)
}
