package main

import (
	"fmt"
	"os"

	"github.com/arshad12/budget-app/internal/api"
	"github.com/arshad12/budget-app/internal/db"
	"github.com/arshad12/budget-app/internal/handlers"
	middlewares "github.com/arshad12/budget-app/internal/middleware"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal("Error loading .env file")
	}

	db, err := db.ConnectDB()
	if err != nil {
		e.Logger.Fatal("Error connecting to database")
	}

	handler := handlers.Handler{
		DB: db,
	}

	api.Routes(&api.Application{
		Logger:  e.Logger,
		Server:  e,
		Handler: &handler,
	})

	e.Use(middleware.Logger())

	e.Use(middlewares.CustomMiddleware)

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
