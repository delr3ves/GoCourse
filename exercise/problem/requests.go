package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	sites := []string{
		//"https://wwsdgsdgw.google.com",
		"https://drive.google.com",
		"https://maps.google.com",
		"https://hangouts.google.com",
	}
	wg.Add(len(sites))

	for _, site := range sites {
		//time.Sleep(time.Millisecond * 250)
		go func(site string) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				res, err := http.Get(site)
				if err != nil {
					io.WriteString(os.Stdout, "Error: "+err.Error()+"\n")
					cancel()
				} else {
					io.WriteString(os.Stdout, res.Status+"\n")
				}
			}
		}(site)
	}
	wg.Wait()
}
