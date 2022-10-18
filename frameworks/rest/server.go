package rest

import (
	"os"

	"github.com/Napigo/go-account-service/frameworks/rest/routes"
	commonrest "github.com/Napigo/npgcommon/rest"
	"github.com/Napigo/npglogger"
	"github.com/gofiber/fiber/v2"
)

func ListeningFunc() error {
	service_name := os.Getenv("SERVICE_NAME")
	port := os.Getenv("SERVICE_PORT")

	npglogger.Printf("Server %v listening on port %s", service_name, port)
	return nil
}

type RestServer struct{}

func (rs RestServer) Run() {
	port := os.Getenv("SERVICE_PORT")

	app := fiber.New(fiber.Config{
		ErrorHandler: commonrest.DefaultErrorResponse,
	})
	app.Hooks().OnListen(ListeningFunc)
	commonrest.CreateRestHook(app)

	routes.HealthCheckRoutes(app)
	routes.AuthRoutes(app)

	app.Listen(port)
}
