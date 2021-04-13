package main

import (
	"github.com/GotoRen/go-restful-api/api/database"
	"github.com/GotoRen/go-restful-api/api/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func main() {
	// DB接続
	database.DBConnecter()
	defer database.Close()

	// ルーティング
	e := echo.New()
	router.Router(e)

	e.Logger.Fatal(e.Start(":8080"))
}
