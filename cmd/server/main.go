package main

import "fmt"

type App struct {
}

func (app *App) Run() error {
	fmt.Println("Run method...")
	return nil
}

func main() {
	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("Error in Starting the Rest API !!!")
		fmt.Println(err)
	}
}
