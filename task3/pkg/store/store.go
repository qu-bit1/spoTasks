package store

import (
	"database/sql"
	"fmt"
)

type Users struct {
	UserID int64
	Pwd    string
}

func Conn() *sql.DB {
	db, err := sql.Open("mysql", "root:campus_fora@tcp(localhost:3306)/spoTask")
	if err != nil {
		fmt.Println("error validating sql.open")
		panic(err.Error())
	}
	return db
}
