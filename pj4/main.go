package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	FirstName string
	LastName  string
}

func main() {
	db, _ := sql.Open("mysql", "root:dbpassword!#edc@tcp(127.0.0.1:3306)/sys")
	defer db.Close()

	db.Query("CREATE TABLE PERSON (first_name varchar(50), last_name varchar(50)")

	db.Query("INSERT INTO PERSON (first_name, last_name) values ('Bob', 'Baker'), ('Betty', 'White') ")
	dataset, _ := db.Query("SELECT * FROM PERSON")

	for dataset.Next() {
		var person Person
		dataset.Scan(&person.FirstName, &person.LastName)

		fmt.Println(person)
	}
}
