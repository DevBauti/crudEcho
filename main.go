package main

import (
	"github.com/DevBauti/crudEcho/configs"
	"github.com/DevBauti/crudEcho/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// inicio
	e := echo.New()
	// db
	configs.ConnectMongo()
	// routes
	routes.UserRouter(e)
	//
	e.Logger.Fatal(e.Start(":6000"))
}
