package infra

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:example0000@tcp(127.0.0.1:3306)/sluck")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}
