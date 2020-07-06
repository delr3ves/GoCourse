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
	executor := getEffectiveGreeter(r)

	fmt.Fprint(w, executor(t))
}

func getEffectiveGreeter(r *http.Request) func(greet.Greet) string {
	var decorator = r.Context().Value("greetDecorator")
	var executor func(greet.Greet) string
	if decorator == nil {
		executor = sayHi
	} else {
		executor = decorator.(func(func(greet.Greet) string) func(greet.Greet) string)(sayHi)
	}
	return executor
}
