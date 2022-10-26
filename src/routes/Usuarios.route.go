package routes

import (
	"github.com/albinovejar/api_templateV1/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func RoutesUsuarios(app *fiber.App) {
	route := app.Group("/usuarios")

	route.Get("/:id", controllers.Usuarios_GetOne)
	route.Get("/", controllers.Usuarios_GetAll)
	route.Put("/", controllers.Usuarios_Edit)
	route.Delete("/", controllers.Usuarios_Delete)
}
