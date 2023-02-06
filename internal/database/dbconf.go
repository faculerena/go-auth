package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

func Setup(config mysql.Config) (DB *sql.DB, err error) {

	if DB, err = sql.Open("mysql", config.FormatDSN()); err != nil {
		return nil, err
	}
	if pingErr := DB.Ping(); pingErr != nil {
		return nil, err
	}
	return DB, nil
}
