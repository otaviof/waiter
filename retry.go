package main

import (
	"log"
	"time"
)

// retry re-executes the informed function for the desired amount of times, waiting for the informed
// duration. It returns a boolean indicating if timeout was reached.
func retry(attempts int64, sleep time.Duration, fn func() bool) bool {
	log.Printf("Attempts '%d'...", attempts)
	shouldRetry := fn()
	if !shouldRetry {
		log.Print("Done!")
		return true
	}
	attempts--
	if attempts <= 0 {
		return false
	}

	time.Sleep(sleep)
	return retry(attempts, sleep, fn)
}
