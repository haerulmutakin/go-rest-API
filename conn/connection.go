package connection

import (
	"database/sql"
)

// ConnectToDb ...
func ConnectToDb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/godb")

	if err != nil {
		panic(err.Error())
	}

	return db
}
