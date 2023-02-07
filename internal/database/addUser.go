package database

import (
	"database/sql"
	"fmt"
)

func AddUser(db *sql.DB, user User) (int64, error) {

	_, ok := GetUser(db, user.Username)
	if ok {
		return -1, nil //handling -1 as user already exists
	}

	result, err := db.Exec("INSERT INTO users (username, passhash, roles) VALUES (?, ?, ?)", user.Username, user.Passhash, user.Roles)
	if err != nil {
		return 0, fmt.Errorf("add user: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("add user: %v", err)
	}
	return id, nil
}
