package data

import (
	"database/sql"
	"fmt"
	"os"
)

func GetUser(db *sql.DB, username string) []User {
	entry, err := db.Query("SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
	var onlyentry []User
	for entry.Next() {
		var std User
		if err := entry.Scan(&std.ID, &std.Username, &std.Passhash, &std.Roles); err != nil {
			_ = fmt.Errorf("error: %v", err)
		}

		onlyentry = append(onlyentry, std)
	}

	return onlyentry
}
