package main

import (
	"simple-gin-api/app"
	"simple-gin-api/app/controller"
)

func main() {

	// Initializing application
	app := app.App{}
	app.Init()

	user := controller.User{}
	user.Init(app.HTTPServer)

	// Start application
	app.Run()
}
