package data

import (
	"database/sql"
	"fmt"
)

func UpdateStudent(db *sql.DB, userId int, std User) (int64, error) {
	result, err := db.Exec("UPDATE users SET username =?,  passhash= ?,  roles= ? WHERE id_user=?", std.Username, std.Passhash, std.Roles, std.ID)
	if err != nil {
		return 0, fmt.Errorf("update user: %v", err)
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("update user: %v", err)
	}
	return id, nil
}
