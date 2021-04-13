package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

var DB *sql.DB

func Conn() *sql.DB {
	return DB
}

func Close() error {
	return DB.Close()
}

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

func DBConnecter() {
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
