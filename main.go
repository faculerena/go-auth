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

	authRouter.Handle("/check", token.TokenValidationMiddleware(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Token is valid"))
	}))).Methods("GET")

	server := &http.Server{
		Addr:    "127.0.0.1:9090",
		Handler: mainRouter,
	}
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Server initialization error")
	}

}
