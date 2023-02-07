package database_test

import (
	"github.com/faculerena/goauth/internal/database"
	"github.com/faculerena/goauth/private"
	"testing"
)

func TestAddUser(t *testing.T) {
	// In-memory database for testing purposes
	db, err := database.Setup(private.GetConfig())
	if err != nil {
		t.Fatalf("error opening database: %v", err)
	}
	defer db.Close()

	// Dummy user for testing purposes
	user := database.User{
		Username: "testuser",
		Passhash: "testpass",
		Roles:    1,
	}

	// Test adding a user that does not exist
	id, err := database.AddUser(db, user)
	if err != nil {
		t.Fatalf("error adding user: %v", err)
	}
	if id <= 0 {
		t.Fatalf("expected positive id, got %d", id)
	}

	// Test adding a user that already exists
	idduplicated, err := database.AddUser(db, user)
	if err == nil && idduplicated != -1 {
		t.Fatalf("expected error adding user, got nil")
	}
}
