package model

import (
	"database/sql"
)

var DB *sql.DB

type User struct {
	ID   int    `required:"true"`
	Name string `required:"true"`
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
