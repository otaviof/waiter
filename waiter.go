package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

// Waiter actor that creates a lock-file and wait for the file to be deleted, or returns error upon
// timeout (maximum retries).
type Waiter struct {
	lockFilePath string        // path to the lock-file
	retries      int64         // amount of times to retry
	interval     time.Duration // sleep duration between retries
}

// ErrTimeout when maximum amount of retries is reached.
var ErrTimeout = errors.New("reached maximum amount of retries")

// save writes the lock-file with informed PID.
func (w *Waiter) save(pid int) error {
	f, err := os.Create(w.lockFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	log.Printf("Saving '%d' (PID) on '%s' lock-file", pid, w.lockFilePath)
	pidStr := strconv.Itoa(pid)
	if _, err = f.WriteString(pidStr); err != nil {
		return err
	}
	return f.Sync()
}

// read reads the lock-file, must contain an integer.
func (w *Waiter) read() (int, error) {
	if _, err := os.Stat(w.lockFilePath); err != nil {
		return -1, err
	}
	data, err := ioutil.ReadFile(w.lockFilePath)
	if err != nil {
		return -1, err
	}
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return -1, err
	}
	return pid, nil
}

// Wait wait for the lock-file to be removed, or timeout.
func (w *Waiter) Wait() error {
	pid := os.Getpid()
	if err := w.save(pid); err != nil {
		return err
	}

	// will inspect the lock file and only return "false", which stops retry loop, if the lock-file
	// is not found
	done := retry(w.retries, w.interval, func() bool {
		_, err := os.Stat(w.lockFilePath)
		return err == nil || !os.IsNotExist(err)
	})
	if !done {
		_ = os.RemoveAll(w.lockFilePath)
		return fmt.Errorf("%w: elapsed %v", ErrTimeout, time.Duration(w.retries)*w.interval)
	}
	return nil
}

// Done removes the lock-file.
func (w *Waiter) Done() error {
	pid, err := w.read()
	if err != nil {
		return err
	}
	log.Printf("Removing lock-file at '%s' (%d PID)", w.lockFilePath, pid)
	return os.Remove(w.lockFilePath)
}

// NewWaiter instantiate the Waiter.
func NewWaiter(lockFilePath string, retries int64, interval time.Duration) *Waiter {
	return &Waiter{
		lockFilePath: lockFilePath,
		retries:      retries,
		interval:     interval,
	}
}
