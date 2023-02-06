package data

import (
	"crypto/sha256"
	"fmt"
)

type User struct {
	ID       int64
	Username string
	Passhash string
	Roles    int
}

func (u *User) ValidateHash(pswhash []byte) bool {
	return u.Passhash == fmt.Sprintf("%x", sha256.Sum256([]byte(pswhash)))
}
