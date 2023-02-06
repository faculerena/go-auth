package token

import (
	"errors"
	"fmt"
	"github.com/faculerena/goauth/internal/database"
	"github.com/faculerena/goauth/private"
	"net/http"
)

func getSignedToken(username string, passphrase string) (string, error) {

	claimsMap := private.Payload(username, passphrase)
	secret := private.Secret()
	header := private.Header()
	tokenString, err := GenerateToken(header, claimsMap, secret)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

func validateUser(username string, passwordHash []byte) (bool, error) {
	db, _ := database.Setup(private.GetConfig())

	user, exists := database.GetUser(db, username)
	if !exists {
		return false, errors.New("user does not exist")
	}
	passwordCheck := user[0].ValidateHash(passwordHash)

	if !passwordCheck {
		return false, nil
	}
	return true, nil
}

func SigninHandler(rw http.ResponseWriter, r *http.Request) {

	if _, ok := r.Header["Username"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Username Missing"))
		return
	}
	if _, ok := r.Header["Passwordhash"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Passwordhash Missing"))
		return
	}

	valid, err := validateUser(r.Header["Username"][0], []byte(r.Header["Passwordhash"][0]))

	if err != nil {
		// user does not exist
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("User Does not Exist"))
		return
	}

	if !valid {
		// wrong password
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("Incorrect Password"))
		return
	}
	tokenString, err := getSignedToken(r.Header["Username"][0], r.Header["Passwordhash"][0])
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Internal Server Error"))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(tokenString))
}
