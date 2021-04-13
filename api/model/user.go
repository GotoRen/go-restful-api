package model

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID   int
	Name string
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

	fmt.Println("users:", users)
	return users, nil
}

func FindById(db *sql.DB, id int) (User, error) {
	fmt.Println("id:", id)
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

	fmt.Println("user:", user.ID, user.Name)
	return user, nil
}
