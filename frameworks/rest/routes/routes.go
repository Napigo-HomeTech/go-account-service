package routes

import (
	"github.com/Napigo/go-account-service/frameworks/rest/controllers"
	"github.com/gofiber/fiber/v2"
)

func HealthCheckRoutes(app *fiber.App) {
	// all GET method for Health Checks endpoint
	app.Get("/health-check", controllers.HealthCheckController)
}

func AuthRoutes(app *fiber.App) {
	app.Get("/jwt", controllers.GetJwtController)
	app.Get("/test-jwt", controllers.GetTestJwtController)
}
