package main

import (
	"bytes"
	"encoding/gob"
)

type Person struct {
	FirstName string
	LastName  string
}

func writeBinaryFile(data interface{}, file string) {
	buf := new(bytes.Buffer)
	encorder := gob.NewEncoder(buf)
}
func main() {
	a
}
