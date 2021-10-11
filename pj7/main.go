package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	FirstName string   `xml:"fname"`
	LastName  string   `xml:"lname"`
	Hobbies   []string `xml:hobbies`
}

func main() {
	joe := Person{FirstName: "Joe", LastName: "Smith"}
	joe.Hobbies = []string{"skies", "Wind Surfing"}

	xmlOutput, _ := xml.MarshalIndent(&joe, " ", "  ")
	fmt.Println(xmlOutput)
	fmt.Println(string(xmlOutput))

	joeFromxml := Person{}

	xml.Unmarshal(xmlOutput, &joeFromxml)
	fmt.Println(joeFromxml)
}
