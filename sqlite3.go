package main

import (
	_ "database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const SQL = "SELECT `name` FROM `users` WHERE id = :id"

type User struct {
	name string `db:"name"`
}

func main() {
	db, err := sqlx.Connect("sqlite3", "./db.sqlite")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmt, err := db.PrepareNamed(SQL)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	user := &User{}
	args := map[string]interface{}{"id": 1}
	err = stmt.Get(user, args)
	fmt.Println(user)
	fmt.Println(err)

}
