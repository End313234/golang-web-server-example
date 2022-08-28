package main

import (
	"log"
	"web-server-example/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.SetupRoutes(app)

	log.Print(app.Listen(":3000"))
}
