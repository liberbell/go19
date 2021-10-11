package main

import "github.com/jmoiron/sqlx"

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

func main() {
	db, _ := sqlx.Open("mysql", "root:dbpassword!#edc@tcp(127.0.0.1:3306)/sys")
	defer db.Close()

	db.Query("CREATE TABLE PERSON (first_name varchar(50), last_name varchar(50))")
	db.Query("INSERT INTO PERSON (first_name, last_name) values ('Bob', 'Barker'),('Betty', 'White')")
}
