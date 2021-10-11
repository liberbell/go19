package main

import "encoding/xml"

type Person struct {
	FirstName string   `xml:"fname"`
	LastName  string   `xml:"lname"`
	Hobbies   []string `xml:hobbies`
}

func main() {
	joe := Person{FirstName: "Joe", LastName: "Smith"}
	joe.Hobbies = []string{"skies", "Wind Surfing"}

	xmlOutput, _ := xml.MarshalIndent(&joe, " ", "  ")
}
