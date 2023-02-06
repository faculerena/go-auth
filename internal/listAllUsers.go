package data

import (
	"database/sql"
	"fmt"
)

func ListAll(db *sql.DB) ([]User, error) {

	var students []User

	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		return nil, fmt.Errorf("error in query all student: %v", err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign record to slice
	for rows.Next() {
		var std User
		if err := rows.Scan(&std.ID, &std.Username, &std.Passhash, &std.Roles); err != nil {
			return nil, fmt.Errorf("error in query all student: %v", err)
		}
		students = append(students, std)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in query all student: %v", err)
	}
	return students, nil
}
