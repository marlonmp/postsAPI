package user

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo"
)

var (
	re_userName, _ = regexp.Compile(`^[\w\d\_\-]{3,35}$`)
	re_nickName, _ = regexp.Compile(`^[\w\d\_\-\ ]{3,35}$`)
	re_password, _ = regexp.Compile(`^[\@\&\%\$\"\'\d\w\_\-]{8,40}$`)
)

func validateGetAccount(next echo.HandlerFunc) echo.HandlerFunc {

	return func(ctx echo.Context) error {

		userName, err := ctx.Cookie("userName")

		if err != nil {
			ctx.NoContent(http.StatusBadRequest)
		}

		isValid := re_userName.MatchString(userName.Value)

		if !isValid {
			return ctx.NoContent(http.StatusBadRequest)
		}

		user := new(Model)

		user.UserName = userName.Value

		ctx.Set("checked.user", user)

		return next(ctx)
	}
}

func validateSignUp(next echo.HandlerFunc) echo.HandlerFunc {

	return func(ctx echo.Context) error {

		user := new(Model)

		err := ctx.Bind(user)

		if err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}

		isValid := re_nickName.MatchString(user.NickName)
		isValid = isValid && re_userName.MatchString(user.UserName)
		isValid = isValid && re_password.MatchString(user.Password)

		if !isValid {
			return ctx.NoContent(http.StatusBadRequest)
		}

		user.Id = nil

		ctx.Set("checked.user", user)

		return next(ctx)
	}
}

func validateSignIn(next echo.HandlerFunc) echo.HandlerFunc {

	return func(ctx echo.Context) error {

		user := new(Model)

		err := ctx.Bind(user)

		if err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}

		isValid := re_userName.MatchString(user.UserName) && re_password.MatchString(user.Password)

		if !isValid {
			return ctx.NoContent(http.StatusBadRequest)
		}

		user.Id = nil
		user.NickName = ""

		ctx.Set("checked.user", user)

		return next(ctx)
	}
}

func validateUpdateAccount(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		user := new(Model)

		userName, err := ctx.Cookie("userName")

		if err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}

		err = ctx.Bind(user)

		if err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}

		isValid := re_userName.MatchString(userName.Value)

		if user.NickName != "" && isValid {
			isValid = isValid && re_nickName.MatchString(user.NickName)
		}

		if user.UserName != "" && isValid {
			isValid = isValid && re_userName.MatchString(user.UserName)
		}

		if user.Password != "" && isValid {
			isValid = isValid && re_password.MatchString(user.Password)
		}

		if !isValid {
			return ctx.NoContent(http.StatusBadRequest)
		}

		oldUser := new(Model)

		oldUser.UserName = userName.Value

		user.Id = nil

		ctx.Set("checked.oldUser", oldUser)
		ctx.Set("checked.user", user)

		return next(ctx)
	}
}

func validateDeleteAccount(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		user := new(Model)

		err := ctx.Bind(user)

		if err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}

		isValid := re_userName.MatchString(user.UserName)

		if !isValid {
			return ctx.NoContent(http.StatusBadRequest)
		}

		user.Id = nil
		user.NickName = ""
		user.Password = ""

		ctx.Set("checked.user", user)

		return next(ctx)
	}
}
