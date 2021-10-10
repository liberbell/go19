package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Person struct {
	FirstName string
	LastName  string
}

func writeBinaryFile(data interface{}, file string) {
	buf := new(bytes.Buffer)
	encorder := gob.NewEncoder(buf)
	encorder.Encode(data)
	ioutil.WriteFile(file, buf.Bytes(), 0600)
}

func readBinaryFile(data interface{}, file string) {
	raw, _ := ioutil.ReadFile(file)
	buf := bytes.NewBuffer(raw)
	decoder := gob.NewDecoder(buf)
	decoder.Decode(data)
}
func main() {
	file := "person.txt"

	person := Person{"Bob", "Baker"}
	writeBinaryFile(person, file)

	var readPerson Person
	readBinaryFile(&readPerson, file)

	fmt.Println(readPerson)
}
