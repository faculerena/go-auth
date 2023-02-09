package database

import (
	"github.com/faculerena/goauth/internal/database"
	"github.com/faculerena/goauth/private"
	"testing"
)

func TestUpdateStudent(t *testing.T) {

	db, err := database.Setup(private.GetConfig())
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	// Create the test data
	_, err = db.Exec("CREATE TABLE users (id_user INTEGER PRIMARY KEY, username TEXT, passhash TEXT, roles TEXT)")
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}
	_, err = db.Exec("INSERT INTO users (id_user, username, passhash, roles) VALUES (1, 'johndoe', 'password123', 'student')")
	if err != nil {
		t.Fatalf("failed to insert data: %v", err)
	}

	// Call the UpdateStudent function
	std := database.User{
		ID:       1,
		Username: "johndoe",
		Passhash: "newpassword456",
		Roles:    1,
	}
	id, err := database.UpdateUser(db, 1, std)
	if err != nil {
		t.Fatalf("failed to update student: %v", err)
	}
	if id != 1 {
		t.Errorf("expected 1 row affected, but got %d", id)
	}

	// Check the results
	var username, passhash, roles string
	err = db.QueryRow("SELECT username, passhash, roles FROM users WHERE id_user = 1").Scan(&username, &passhash, &roles)
	if err != nil {
		t.Fatalf("failed to query data: %v", err)
	}
	if username != "johndoe" {
		t.Errorf("expected username 'johndoe', but got '%s'", username)
	}
	if passhash != "newpassword456" {
		t.Errorf("expected passhash 'newpassword456', but got '%s'", passhash)
	}
	if roles != "teacher" {
		t.Errorf("expected roles 'teacher', but got '%s'", roles)
	}
}
