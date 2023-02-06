package main

import (
	"database/sql"
	"fmt"
	data "github.com/faculerena/goauth/internal"
	"github.com/faculerena/goauth/private"
	"github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func main() {

	if err := setupDB(private.GetConfig()); err != nil {
		log.Fatal(err)
	}

	check, err := data.CheckUser(db)
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
