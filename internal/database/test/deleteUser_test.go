package database_test

import (
	"github.com/faculerena/goauth/internal/database"
	"github.com/faculerena/goauth/private"
	"testing"
)

func TestDeleteUser(t *testing.T) {
	// In-memory database for testing purposes
	db, err := database.Setup(private.GetConfig())
	if err != nil {
		t.Fatalf("error opening database: %v", err)
	}
	defer db.Close()

	// Dummy user for testing purposes
	user := database.User{
		Username: "testuserdelete",
		Passhash: "testpass",
		Roles:    1,
	}

	// Add the user to the database
	id, err := database.AddUser(db, user)
	if err != nil {
		t.Fatalf("error adding user: %v", err)
	}

	// Test deleting the user
	rowsAffected, err := database.DeleteUser(db, int(id))
	if err != nil {
		t.Fatalf("error deleting user: %v", err)
	}
	if rowsAffected <= 0 {
		t.Fatalf("expected positive rows affected, got %d", rowsAffected)
	}

	// Test deleting a user that does not exist
	rowsAffected, err = database.DeleteUser(db, int(id))
	if err != nil {
		t.Fatalf("error deleting user: %v", err)
	}
	if rowsAffected > 0 {
		t.Fatalf("expected 0 rows affected, got %d", rowsAffected)
	}
}
