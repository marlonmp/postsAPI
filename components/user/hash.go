package user

import (
	"github.com/labstack/echo"
	"github.com/marlonmp/argon2"
	"github.com/marlonmp/postsAPI/util"
)

func hashPassword(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		user := ctx.Get("checked.user").(*Model)

		user.Password = argon2.GenerateHash(user.Password, util.Argon2Options)

		return next(ctx)
	}
}
