package controller

import (
	"net/http"
	"strconv"

	"github.com/GotoRen/go-restful-api/api/database"
	"github.com/GotoRen/go-restful-api/api/model"
	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	users, err := model.FindAll(database.Conn())
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
	user, err := model.FindById(database.Conn(), userid)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
