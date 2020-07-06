package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

func POST(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func Debug(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer h.ServeHTTP(w, r)

		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(dump))
	}
}

func Delay(delay time.Duration, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer h.ServeHTTP(w, r)

		time.Sleep(delay)
	}
}


func Cached(h http.HandlerFunc) http.HandlerFunc {
	var cachedGreeter = &cachedGreeter{
		cache: cache{
			byName: make(map[string]string),
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		cacheableRequest := r.WithContext(context.WithValue(r.Context(), "greetDecorator", cachedGreeter.SayHi))
		h.ServeHTTP(w, cacheableRequest)
	}
}
