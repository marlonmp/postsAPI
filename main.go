package main

import (
	"github.com/labstack/echo"
	"github.com/marlonmp/postsAPI/routes"
)

func main() {

	app := echo.New()

	routes.RoutesUp(app)

	err := app.Start(":5000")

	if err != nil {
		panic(err)
	}
}
