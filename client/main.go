package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/sony/gobreaker/v2"
)

type CircuitBreakerConfiguration struct {
	MaxRequests       uint32
	Interval          time.Duration
	Timeout           time.Duration
	BucketPeriod      time.Duration
	MinRequestsToOpen uint32
	FailureRateToOpen float64
}

func main() {
	cbConfig := CircuitBreakerConfiguration{
		MaxRequests:       3,
		Interval:          60 * time.Second,
		Timeout:           60 * time.Second,
		BucketPeriod:      10 * time.Second,
		MinRequestsToOpen: 5,
		FailureRateToOpen: 0.6,
	}

	settings := gobreaker.Settings{
		Name:        "BladeRunnerAPI",
		MaxRequests: cbConfig.MaxRequests,
		Interval:    cbConfig.Interval,
		Timeout:     cbConfig.Timeout,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.Requests >= cbConfig.MinRequestsToOpen &&
				float64(counts.TotalFailures)/float64(counts.Requests) >= cbConfig.FailureRateToOpen
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.Printf("Circuit breaker '%s' changed from %s to %s", name, from, to)
		},
	}

	cb := gobreaker.NewCircuitBreaker[any](settings)

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	serverURL := "http://localhost:8080/blade-runner"

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	log.Println("Starting client")

	for {
		select {
		case <-ticker.C:
			makeRequest(cb, client, serverURL)
		}
	}
}

func makeRequest(cb *gobreaker.CircuitBreaker[any], client *http.Client, url string) {
	start := time.Now()
	
	result, err := cb.Execute(func() (interface{}, error) {
		resp, err := client.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
		}

		return string(body), nil
	})

	duration := time.Since(start)

	if err != nil {
		log.Printf("Request failed after %v: %v", duration, err)
	} else {
		log.Printf("Request succeeded after %v: %s", duration, result)
	}
}