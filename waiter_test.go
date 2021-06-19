package main

import (
	"os"
	"testing"
	"time"
)

func TestWaiter_save_read(t *testing.T) {
	expect := 0

	w := NewWaiter("/var/tmp/lock.pid", 0, 0)
	err := w.save(expect)
	if err != nil {
		t.Error(err)
	}

	payload, err := w.read()
	if err != nil {
		t.Error(err)
	}
	if expect != payload {
		t.Errorf("payload '%v' is different than expected '%v'", payload, expect)
	}
}

func TestWaiter_Wait_Done(t *testing.T) {
	tests := []struct {
		name         string
		lockFilePath string
		retries      int64
		interval     time.Duration
		sleep        time.Duration
		callsDone    bool
		expectErr    bool
	}{{
		"default scenario, lock is started and finished before timeout",
		"/var/tmp/lock.pid",
		5,
		1 * time.Second,
		2 * time.Second,
		true,
		false,
	}, {
		"error scenario, lock is started but is NOT finshed before timout",
		"/var/tmp/lock.pid",
		3,
		1 * time.Second,
		5 * time.Second,
		false,
		true,
	}}

	for _, tt := range tests {
		t.Logf("lockFilePath='%s', retries='%d', interval='%v'", tt.lockFilePath, tt.retries, tt.interval)
		t.Run(tt.name, func(t *testing.T) {
			errCh := make(chan error, 1)
			w := NewWaiter(tt.lockFilePath, tt.retries, tt.interval)
			go func() {
				errCh <- w.Wait()
			}()

			time.Sleep(tt.sleep)
			if tt.callsDone {
				_ = w.Done()
			}

			_ = os.RemoveAll(tt.lockFilePath)

			err := <-errCh
			t.Logf("err='%v'", err)
			if (tt.expectErr && err == nil) || (!tt.expectErr && err != nil) {
				t.Errorf("Waiter.Wait(): err='%v', expect error '%v'", err, tt.expectErr)
			}
			close(errCh)
		})
	}
}
