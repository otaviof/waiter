package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type Waiter struct {
	pidPath string
}

const pidPath = "/var/tmp/waiter.pid"

func (w *Waiter) savePid(pid int) error {
	f, err := os.Create(w.pidPath)
	if err != nil {
		return err
	}
	defer f.Close()

	log.Printf("Saving PID '%d' at '%s'", pid, w.pidPath)
	pidStr := strconv.Itoa(pid)
	if _, err = f.WriteString(pidStr); err != nil {
		return err
	}
	return f.Sync()
}

func (w *Waiter) readPid() (int, error) {
	if _, err := os.Stat(w.pidPath); err != nil {
		return -1, err
	}
	data, err := ioutil.ReadFile(w.pidPath)
	if err != nil {
		return -1, err
	}
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return -1, err
	}
	return pid, nil
}

func (w *Waiter) Wait() error {
	pid := os.Getpid()
	if err := w.savePid(pid); err != nil {
		return err
	}
	waitForFileDeletion(w.pidPath)
	return nil
}

func (w *Waiter) Done() error {
	pid, err := w.readPid()
	if err != nil {
		return err
	}
	log.Printf("Removing pid-file '%s' (%d)", w.pidPath, pid)
	return os.Remove(w.pidPath)
}

func NewWaiter() *Waiter {
	return &Waiter{
		pidPath: pidPath,
	}
}
