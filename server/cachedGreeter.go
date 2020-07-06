package main

import "github.com/pabloos/http/greet"

type cachedGreeter struct {
	id    int
	cache cache
}

func (greeter cachedGreeter) SayHi(greet greet.Greet, f func(greet.Greet) string) string {
	cached, ok := greeter.cache.FindByName(greet.Name)
	if !ok {
		cached = greeter.cache.AddName(greet.Name, f(greet))
	}
	return cached
}

var defaultCachedGreeter = &cachedGreeter{
	cache: cache{
		byName: make(map[string]string),
	},
}
