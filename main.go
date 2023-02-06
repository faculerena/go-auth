package main

import (
	"bufio"
	"database/sql"
	"fmt"
	data "github.com/faculerena/goauth/internal"
	"github.com/faculerena/goauth/private"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"strings"
)

var db *sql.DB

func main() {

	if err := setupDB(private.GetConfig()); err != nil {
		log.Fatal(err)
	}

	check, err := checkUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(check)
}

func setupDB(config mysql.Config) (err error) {
	if db, err = sql.Open("mysql", config.FormatDSN()); err != nil {
		return err
	}
	if pingErr := db.Ping(); pingErr != nil {
		return err
	}
	return nil
}

func checkUser() (b bool, err error) {

	username, password, err := getCredentials()
	if err != nil {
		return false, nil
	}

	onlyentry := data.GetUser(db, username)
	if onlyentry == nil {
		return false, nil
	}

	return onlyentry[0].ValidateHash(password), err
}

func getCredentials() (username string, password []byte, err error) {

	fmt.Print("> ")
	usernameRaw, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", nil, err
	}
	username = strings.TrimSuffix(usernameRaw, "\n")

	fmt.Print("> ")
	password, err = terminal.ReadPassword(0)
	if err != nil {
		return "", nil, err

	}

	return username, password, nil
}
