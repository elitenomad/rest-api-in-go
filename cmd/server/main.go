package main

import (
	"fmt"
	"net/http"

	"github.com/elitenomad/rest-api/internal/database"
	transportHTTP "github.com/elitenomad/rest-api/internal/transport/http"
)

type App struct {
}

func (app *App) Run() error {
	fmt.Println("Setting up...")

	if _, err := database.NewDatabase(); err != nil {
		return err
	}

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("Error in Starting the Rest API !!!")
		fmt.Println(err)
	}
}
