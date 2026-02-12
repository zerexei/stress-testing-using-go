package main

import (
	"net/http"
	"sync"
	"fmt"
)

func main() {
	// totalRequests := 1000000
	totalRequests := 1000
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
					fmt.Println("Server not reachable:", err)
					continue
				}
				resp.Body.Close()
			}
		}()
	}

	wg.Wait()
}
