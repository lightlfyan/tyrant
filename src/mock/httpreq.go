package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var wg sync.WaitGroup
	var url = "http://127.0.0.1:8888/api/url1"
	workcount := 100
	wg.Add(workcount)
	for i := 0; i < workcount; i++ {
		go func(url string) {
			resp, err := http.Get(url)
			defer resp.Body.Close()
			if err != nil {
				fmt.Println("err: ", err)
			} else {
				fmt.Println("status: ", resp.StatusCode)
			}
			defer wg.Done()
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}
