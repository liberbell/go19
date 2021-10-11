package main

import "net/http"

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
}

func main() {
	var People []Person
	http.HandleFunc("/person/", func(rw http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
		}
	})
}
