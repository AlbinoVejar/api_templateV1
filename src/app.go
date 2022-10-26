package src

import (
	"os"

	"github.com/albinovejar/api_templateV1/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}
	return port
}

func InitServer() {
	app := fiber.New()

	app.Use(cors.New(
		cors.Config{
			AllowOrigins:     "*",
			AllowMethods:     "GET,POST,HEAD,PUT,DELETE",
			AllowHeaders:     "",
			AllowCredentials: false,
			Next:             nil,
		},
	))

	app.Use(logger.New())

	routes.RoutesUsuarios(app)

	app.Listen(getPort())
}
