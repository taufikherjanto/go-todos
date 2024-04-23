package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitDb() *sql.DB {
	const (
		// config database
		username = "root"
		password = ""
		host     = "localhost"
		port     = "3306"
		dbname   = "eduwork"
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	return db
}
