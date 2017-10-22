package main

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const SQL = "SELECT `name` FROM `users` WHERE id = :id"

type User struct {
	name string `db:"name"`
}

func main() {
	db, err := sqlx.Connect("mysql", "root:@unix(/tmp/mysql.sock)/foobar_m0cchi____duplicate_guard?parseTime=true")

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
