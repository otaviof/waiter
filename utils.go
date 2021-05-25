package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

func waitForSignal() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh)

	log.Print("Waiting for `done` signal...")
	s := <-signalCh
	fmt.Printf("\n")
	log.Printf("Exit signal '%#v'", s)
}

func retry(attempts int, sleep time.Duration, fn func() bool) {
	log.Printf("Attempts '%d'...", attempts)
	shouldRetry := fn()
	if !shouldRetry {
		log.Print("Done!")
		return
	}
	attempts--
	if attempts <= 0 {
		return
	}

	time.Sleep(sleep)
	retry(attempts, sleep, fn)
}

func waitForFileDeletion(pidFilePath string) {
	retry(600, 1*time.Second, func() bool {
		_, err := os.Stat(pidFilePath)
		if err == nil {
			return true
		}
		return !os.IsNotExist(err)
	})
}
