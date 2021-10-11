package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusBadRequest)
	})

	http.HandleFunc("/withargs/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "This is a response from a %s request. The arguments are: args1:%s, args2:%s", r.Method, r.URL.Query().Get("arg1"), r.URL.Query().Get("arg2"))
	})
}
