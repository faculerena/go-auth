package database

import (
	"github.com/faculerena/goauth/internal/database"
	"github.com/faculerena/goauth/private"
	"testing"
)

func TestGetUser(t *testing.T) {
	// In-memory database for testing purposes
	db, err := database.Setup(private.GetConfig())
	if err != nil {
		t.Fatalf("error opening database: %v", err)
	}
	defer db.Close()

	// Dummy user for testing purposes
	user := database.User{
		Username: "testuserget",
		Passhash: "testpass",
		Roles:    1,
	}

	// Add the user to the database
	_, err = database.AddUser(db, user)
	if err != nil {
		t.Fatalf("error adding user: %v", err)
	}

	// Test retrieving the user
	entries, ok := database.GetUser(db, user.Username)
	if !ok {
		t.Fatalf("error retrieving user")
	}

	if len(entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(entries))
	}
	if entries[0].Username != user.Username || entries[0].Passhash != user.Passhash || entries[0].Roles != user.Roles {
		t.Fatalf("expected %v, got %v", user, entries[0])
	}

	// Test retrieving a user that does not exist
	_, ok = database.GetUser(db, "nonexistentuser")
	if ok {
		t.Fatalf("expected not ok, got ok")
	}
}
