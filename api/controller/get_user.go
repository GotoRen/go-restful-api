package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/GotoRen/go-restful-api/api/model"
	"github.com/labstack/echo"
)

var DB *sql.DB

type User struct {
	ID   int    `required:"true"`
	Name string `required:"true"`
}

func SayHello() {
	fmt.Println("Hello")
}

func GetUsers(c echo.Context) error {
	users, err := model.FindAll(DB)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	userid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	user, err := model.FindById(DB, userid)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
