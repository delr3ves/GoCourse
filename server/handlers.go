package main

import (
	"encoding/json"
	"fmt"
	"github.com/pabloos/http/greet"
	"io"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "You are on the index page\n")
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	var t greet.Greet

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	greeter := defaultCachedGreeter
	fmt.Fprint(w, greeter.SayHi(t, sayHi))
}
