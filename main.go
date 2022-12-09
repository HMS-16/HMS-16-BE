package main

import (
	"HMS-16-BE/config"
	"HMS-16-BE/database"
	"HMS-16-BE/route"
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	appMode := os.Getenv("APP_MODE")
	if appMode == "" {
		config.InitConfig()
	}
	database.InitGorm()
	db := database.InitMySql()

	app := echo.New()
	route.Init(app, db)

	app_port := os.Getenv("PORT")
	if app_port == "" {
		app_port = config.Cfg.API_PORT
	}

	port := fmt.Sprintf(":%s", app_port)
	app.Logger.Fatal(app.Start(port))
}
