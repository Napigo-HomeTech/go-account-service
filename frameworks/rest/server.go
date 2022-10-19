package rest

import (
	"github.com/Napigo/go-account-service/frameworks/rest/routes"
	"github.com/Napigo/npgc"
	"github.com/Napigo/npglogger"
	"github.com/gofiber/fiber/v2"
)

func ListeningFunc() error {
	config := npgc.Config
	npglogger.Printf("Server %v listening on port %s", config.ServiceName, config.Port)
	return nil
}

type RestServer struct{}

func (rs RestServer) Run() {
	app := fiber.New(fiber.Config{
		ErrorHandler: npgc.DefaultErrorResponse,
	})
	app.Hooks().OnListen(ListeningFunc)
	npgc.CreateRestHook(app)

	routes.HealthCheckRoutes(app)
	routes.AuthRoutes(app)

	port := npgc.Config.Port
	app.Listen(port)
}
