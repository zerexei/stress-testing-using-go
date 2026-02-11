package main

import (
	"net/http"
	"sync"
)

func main() {
	totalRequests := 1000000
	concurrency := 200 // number of goroutines
	targets := []string{
		"http://localhost:8081",
		"http://localhost:8082",
	}

	var wg sync.WaitGroup
	wg.Add(concurrency)

	requestsPerGoroutine := totalRequests / concurrency

	client := &http.Client{}

	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < requestsPerGoroutine; j++ {
				url := targets[j%len(targets)] // simple round-robin
				resp, err := client.Get(url)
				if err != nil {
					continue
				}
				resp.Body.Close()
			}
		}()
	}

	wg.Wait()
}
