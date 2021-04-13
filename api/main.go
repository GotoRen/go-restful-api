package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

func dbClient(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func init() {
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
}

func main() {
	e := echo.New()
	e.GET("/users", GetUsers)
	e.GET("/users/:id", GetUser)

	e.Logger.Fatal(e.Start(":1323"))
}

func GetUsers(c echo.Context) error {
	users, err := FindAll(DB)
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
	user, err := FindById(DB, userid)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func FindAll(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT * FROM users")
	defer rows.Close()
	if err != nil {
		return []User{}, err
	}

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return []User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}

func FindById(db *sql.DB, id int) (User, error) {
	rows, err := db.Query("SELECT * FROM users WHERE id = ?", id)
	defer rows.Close()
	if err != nil {
		return User{}, err
	}

	rows.Next()
	var user User
	if err := rows.Scan(&user.ID, &user.Name); err != nil {
		return User{}, err
	}
	return user, nil
}
