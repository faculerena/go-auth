package token

import (
	"github.com/faculerena/goauth/internal/database"
	"github.com/faculerena/goauth/private"
	"log"
	"net/http"
)

func SignupHandler(rw http.ResponseWriter, r *http.Request) {

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

	db, err := database.Setup(private.GetConfig())
	if err != nil {
		log.Fatal(err)
	}

	id, err := database.AddUser(db, database.User{
		Username: r.Header["Username"][0],
		Passhash: "r.Header[\"Passwordhash\"][0]",
	})

	if id == -1 { //if -1 user already exists
		rw.WriteHeader(http.StatusConflict)
		rw.Write([]byte("Username already exists"))
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("User Created"))
}
