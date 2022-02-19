package routes

import (
	"github.com/labstack/echo"
	"github.com/marlonmp/postsAPI/components/user"
)

func RoutesUp(app *echo.Echo) {

	user.RoutesUp(app)
}
