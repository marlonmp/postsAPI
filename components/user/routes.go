package user

import "github.com/labstack/echo"

func RoutesUp(app *echo.Echo) {

	accountGroup := app.Group("/account")

	accountGroup.POST("/", signUp, validateSignUp, hashPassword)

	accountGroup.POST("/signIn", signIn, validateSignIn)

	accountGroup.PATCH("/", updateUser, validateUpdateAccount, hashPassword)

	accountGroup.DELETE("/", deleteUser, validateDeleteAccount)

	accountGroup.GET("/", getUser, validateGetAccount)

	userGroup := app.Group("/user")

	userGroup.GET("/:userName", getUser)
}
