package main

import (
	"fmt"
	"github.com/faculerena/goauth/internal/token"
	"github.com/faculerena/goauth/private"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

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

	myToken, err := token.GenerateToken(private.Header(), private.Payload(), private.Secret())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(myToken)

	isOk, err := token.ValidateToken(myToken, private.Secret())

	fmt.Println(isOk)

	mainRouter := mux.NewRouter()
	authRouter := mainRouter.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/signup", token.SignupHandler)

	// The Signin will send the JWT back as we are making microservices.
	// The JWT token will make sure that other services are protected.
	// So, ultimately, we would need a middleware
	authRouter.HandleFunc("/signin", token.SigninHandler)

	// Add the middleware to different subrouter
	// HTTP server
	// Add time outs
	server := &http.Server{
		Addr:    "127.0.0.1:9090",
		Handler: mainRouter,
	}
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Error Booting the Server")
	}

}
