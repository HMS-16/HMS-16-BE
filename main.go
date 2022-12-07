package main

import (
	"HMS-16-BE/config"
	"HMS-16-BE/database"
	"HMS-16-BE/route"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	config.InitConfig()
	database.InitGorm()
	db := database.InitMySql()

	app := echo.New()
	route.Init(app, db)

	port := fmt.Sprintf(":%s", config.Cfg.API_PORT)
	app.Logger.Fatal(app.Start(port))
}
