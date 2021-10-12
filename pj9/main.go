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
	var people []Person
	http.HandleFunc("/person/", func(rw http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			jsonRequest, _ := json.MarshalIndent(&people, " ", "   ")
			fmt.Fprint(rw, string(jsonRequest))

		case http.MethodPost:
			requestBodyBytes, _ := ioutil.ReadAll(r.Body)
			var newPerson Person
			json.Unmarshal(requestBodyBytes, &newPerson)

			for i := range people {
				if people[i].ID == newPerson.ID {
					rw.WriteHeader(http.StatusConflict)
					return
				}
			}
			people = append(people, newPerson)

		case http.MethodPut:
			requestBodyBytes, _ := ioutil.ReadAll(r.Body)
			var newPerson Person
			json.Unmarshal(requestBodyBytes, &newPerson)

			for i := range people {
				if people[i].ID == newPerson.ID {
					people[i].FirstName = newPerson.FirstName
					people[i].LastName = newPerson.LastName
					return
				}
			}
			rw.WriteHeader(http.StatusNotFound)
		}
	})
}
