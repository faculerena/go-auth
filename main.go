package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/faculerena/goauth/internal/database"
	"github.com/faculerena/goauth/internal/token"
	"github.com/faculerena/goauth/private"
	"github.com/gorilla/mux"
	"net/http"
)

var err error

func main() {
	/*
		db, err := database.Setup(private.GetConfig())
		if err != nil {
			log.Fatal(err)
		}
			check, err := database.CheckUser(db)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(check)

	*/
	/*
		myToken, err := token.GenerateToken(private.Header(), private.Payload(), private.Secret())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(myToken)

			isOk, err := token.ValidateToken(myToken, private.Secret())

			fmt.Println(isOk)
	*/
	db, err := database.Setup(private.GetConfig())
	var s []string
	s = append(s, "perez", "gomez", "lopez")
	for i, si := range s {
		h := sha1.New()
		h.Write([]byte(si))
		hashencoder := hex.EncodeToString(h.Sum(nil))

		fmt.Println(s[i], hashencoder)
	}
	fmt.Println("------------------------")

	database.DeleteUser(db, 4)
	database.DeleteUser(db, 7)
	/*
		id, err := database.AddUser(db, juancito)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id)
	*/
	mainRouter := mux.NewRouter()

	authRouter := mainRouter.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/signup", token.SignupHandler) //returns token

	authRouter.HandleFunc("/signin", token.SigninHandler)

	authRouter.Use(token.TokenValidationMiddleware)

	server := &http.Server{
		Addr:    "127.0.0.1:9090",
		Handler: mainRouter,
	}
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Server initialization error")
	}

}
