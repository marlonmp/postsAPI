package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/marlonmp/argon2"
	"github.com/marlonmp/postsAPI/util"
)

func getUser(ctx echo.Context) error {

	userFilters := ctx.Get("checked.user")

	res, err := util.NewQuery("posts", "user").FindOne(userFilters)

	if err != nil {
		fmt.Println("err:", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	user := new(Model)

	err = res.Decode(user)

	if err == mongo.ErrNoDocuments {
		return ctx.NoContent(http.StatusNotFound)
	}

	return ctx.JSON(http.StatusOK, user)
}

func signUp(ctx echo.Context) error {

	newUser := ctx.Get("checked.user")

	res, err := util.NewQuery("posts", "user").InsertOne(newUser)

	if err != nil {
		fmt.Println("err:", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	obId, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		fmt.Println("error in ObjectID")
		return ctx.NoContent(http.StatusInternalServerError)
	}

	if obId.IsZero() {
		return ctx.NoContent(http.StatusConflict)
	}

	return ctx.NoContent(http.StatusCreated)
}

func signIn(ctx echo.Context) error {

	userFilters := ctx.Get("checked.user").(*Model)

	password := userFilters.Password

	userFilters.Password = ""

	res, err := util.NewQuery("posts", "user").FindOne(userFilters)

	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	user := new(Model)

	err = res.Decode(user)

	if err == mongo.ErrNoDocuments {
		return ctx.NoContent(http.StatusNotFound)
	}

	equal, err := argon2.VerifyHash(password, user.Password)

	if err != nil {
		fmt.Println("err:", err)
		ctx.NoContent(http.StatusInternalServerError)
	}

	if !equal {
		ctx.NoContent(http.StatusForbidden)
	}

	return ctx.NoContent(http.StatusOK)
}

func updateUser(ctx echo.Context) error {

	userFilters := ctx.Get("checked.oldUser")
	user := ctx.Get("checked.user")

	res, err := util.NewQuery("posts", "user").UpdateOne(userFilters, user)

	if err != nil {
		fmt.Println("err:", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	if res.MatchedCount == 0 {
		return ctx.NoContent(http.StatusNotFound)
	}

	if res.ModifiedCount == 0 {
		return ctx.NoContent(http.StatusConflict)
	}

	return ctx.NoContent(http.StatusOK)
}

func deleteUser(ctx echo.Context) error {

	userFilters := ctx.Get("checked.user")

	res, err := util.NewQuery("posts", "user").DeleteOne(userFilters)

	if err != nil {
		fmt.Println("err:", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	if res.DeletedCount == 0 {
		return ctx.NoContent(http.StatusNotFound)
	}

	return ctx.NoContent(http.StatusOK)
}
