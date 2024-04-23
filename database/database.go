package database

import "database/sql"

func InitDb() *sql.DB {
	dsn := "root@tcp(localhost:3306)/eduworks"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	return db
}
