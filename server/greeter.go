package main

import (
	"fmt"
	"github.com/pabloos/http/greet"
)

func sayHi(greet greet.Greet) string {
	if greet.Name == "" || greet.Location == "" {
		return ("Tell us what is your name and where do you come from!\n")
	}
	return fmt.Sprintf("Hello %s, from %s\n", greet.Name, greet.Location)

}
