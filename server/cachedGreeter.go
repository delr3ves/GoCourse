package main

import "github.com/pabloos/http/greet"

type cachedGreeter struct {
	id    int
	cache cache
}

func (greeter cachedGreeter) SayHi(f func(greet.Greet) string) func(greet2 greet.Greet) string {
	return func (greet greet.Greet) string{
		cached, ok := greeter.cache.FindByName(greet.Name)
		if !ok {
			cached = greeter.cache.AddName(greet.Name, f(greet))
		}
		return cached
	}
}
