package main

import (
	"bytes"
	"encoding/gob"
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
func main() {
	a
}
