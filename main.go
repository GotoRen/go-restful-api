package main

import (
	"database/sql"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

var connection *sql.DB

func main() {
	/*** ローカル環境のMySQLに接続 ***/
	//db, err := sql.Open("mysql", "root:password@tcp(:3306)/sample_db")

	/*** DockerコンテナのMySQLに接続 ***/
	db, err := sql.Open("mysql", os.Getenv("DATA_SOURCE_NAME"))
	if err != nil {
		panic(err.Error())
	}
	connection = db
	defer db.Close()

	e := echo.New()
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})
	e.GET("/users", GetUsers)
	e.GET("/users/:id", GetUser)

	e.Logger.Fatal(e.Start(":1323"))
}

// User型
type User struct {
	ID   int
	Name string
}

// /users でアクセス
func GetUsers(c echo.Context) error {
	users, err := FindAll(connection)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

// /users/id でアクセス
func GetUser(c echo.Context) error {
	userid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	user, err := FindById(connection, userid)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// Fetch All
func FindAll(db *sql.DB) ([]User, error) {   // []*user  ユーザ型のポインタの配列を返す
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

// Fetch ID
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
