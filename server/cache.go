package main

import "sync"

type cache struct {
	byName map[string]string
	mu sync.Mutex
}

func (cache cache) FindByName(name string) (string, bool) {
	result, ok := cache.byName[name]
	if (!ok) {
		return "", false
	}
	return result, true
}

func (cache cache) AddName(name string, result string) string {
	cache.mu.Lock()
	cache.byName[name] = result
	cache.mu.Unlock()
	return result
}


