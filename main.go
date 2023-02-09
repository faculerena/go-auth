package main

import (
	"fmt"
	"github.com/faculerena/goauth/internal/token"
	"github.com/gorilla/mux"
	"net/http"
)

var err error

func main() {

	mainRouter := mux.NewRouter()

	authRouter := mainRouter.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/signup", token.SignupHandler) //returns token

	authRouter.HandleFunc("/signin", token.SigninHandler)

	//	authRouter.Use(token.TokenValidationMiddleware) // remove comment if you want to test the token received

	server := &http.Server{
		Addr:    "127.0.0.1:9090",
		Handler: mainRouter,
	}
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Server initialization error")
	}

}
