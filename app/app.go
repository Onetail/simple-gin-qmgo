package app

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type App struct {
	HTTPServer *HTTPServer
	Database   *Database
}

func (app *App) Init() {

	// Initializing configuration
	if os.Getenv("GO_ENV") == "production" {
		viper.SetConfigName("config.prod")
		viper.AddConfigPath("./config")
	} else {
		viper.SetConfigName("config.local")
		viper.AddConfigPath("./config")
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	app.Database = &Database{
		app: app,
	}
	app.Database.Init()

	app.HTTPServer = &HTTPServer{}
	app.HTTPServer.Init(app)
}

func (app *App) Run() {
	app.HTTPServer.Start()
}
