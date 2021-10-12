package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
			jsonRequest, _ := json.MarshalIndent(&People, " ", "   ")
			fmt.Fprint(rw, string(jsonRequest))

		case http.MethodPost:
			requestBodyBytes, _ := ioutil.ReadAll(r.Body)
			var newPerson Person
		}
	})
}
