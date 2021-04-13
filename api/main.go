package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/GotoRen/go-restful-api/api/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
)

type User struct {
	ID   int    `required:"true"`
	Name string `required:"true"`
}

var DB *sql.DB

type dbConfig struct {
	Host     string `required:"true"`
	User     string `required:"true"`
	Password string `required:"true"`
	Database string `required:"true"`
}

func main() {
	dbConf := dbConfig{}
	err := envconfig.Process("db", &dbConf)
	if err != nil {
		log.Printf("db init error: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConf.User, dbConf.Password, dbConf.Host, dbConf.Database)
	db, err := dbClient(dsn)
	if err != nil {
		log.Printf("db init error: %v", err)
	}
	DB = db

	e := echo.New()
	e.GET("/users", GetUsers)
	e.GET("/users/:id", GetUser)

	e.Logger.Fatal(e.Start(":8080"))
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
