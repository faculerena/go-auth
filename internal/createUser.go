package data

import (
	"database/sql"
	"fmt"
)

func AddUser(db *sql.DB, std User) (int64, error) {
	result, err := db.Exec("INSERT INTO users (username, passhash, roles) VALUES (?, ?, ?)", std.Username, std.Passhash, std.Roles)
	if err != nil {
		return 0, fmt.Errorf("add student: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("add student: %v", err)
	}
	return id, nil
}
