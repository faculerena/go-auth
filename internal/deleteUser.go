package data

import (
	"database/sql"
	"fmt"
)

func DeleteStudent(db *sql.DB, userId int) (int64, error) {
	result, err := db.Exec("DELETE from users WHERE id=?", userId)
	if err != nil {
		return 0, fmt.Errorf("delete users: %v", err)
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("delete users: %v", err)
	}
	return id, nil
}
