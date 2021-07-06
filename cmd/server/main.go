package main

import (
	"net/http"

	"github.com/elitenomad/rest-api/internal/comment"
	"github.com/elitenomad/rest-api/internal/database"
	transportHTTP "github.com/elitenomad/rest-api/internal/transport/http"
	logger "github.com/sirupsen/logrus"
)

type App struct {
	Name    string
	Version string
}

func (app *App) Run() error {
	logger.SetFormatter(&logger.JSONFormatter{})
	logger.WithFields(
		logger.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		}).Info("Setting Up Our APP")

	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	service := comment.NewService(db)

	handler := transportHTTP.NewHandler(service)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		logger.Error("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	app := App{
		Name:    "Rest API",
		Version: "1.0",
	}

	if err := app.Run(); err != nil {
		logger.Error("Error in Starting the Rest API !!!")
		logger.Fatal(err)
	}
}
