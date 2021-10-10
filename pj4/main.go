package main

import "database/sql"

type Person struct {
	FirstName string
	LastName  string
}

func main() {
	db, _ := sql.Open("mysql", "root:dbpassword!#edc@tcp(127.0.0.1:3306)/sys")
	defer db.Close()
}
