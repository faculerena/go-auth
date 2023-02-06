package database

type User struct {
	ID       int64
	Username string
	Passhash string
	Roles    int
}

func (u *User) ValidateHash(pswhash []byte) bool {

	return u.Passhash == string(pswhash)
}
