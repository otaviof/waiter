package main

import (
	"log"
	"time"
)

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
