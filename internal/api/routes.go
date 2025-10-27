package api

import (
	"github.com/arshad12/budget-app/internal/handlers"
	"github.com/labstack/echo/v4"
)

type Application struct {
	Logger  echo.Logger
	Server  *echo.Echo
	Handler *handlers.Handler
}

func Routes(app *Application) {
	app.Logger.Info("Routes Initialized")
	app.Server.GET("/", app.Handler.HealthCheck)
}
