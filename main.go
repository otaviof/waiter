package main

import (
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd primary cobra.Command instance, sub-commands will be attached to it.
var rootCmd = &cobra.Command{
	Short:        "Waits until `done` is issued.",
	Long:         rootCmdLongDesc,
	Use:          "waiter [flags]",
	SilenceUsage: true,
}

const rootCmdLongDesc = `
Idle loop to hold a container (or maybe a Kubernetes POD) running while some other action happens in
the background. It is started by issuing "waiter start" and can be stopped with "waiter done", or
after timeout. Please check "--help" to inspect the flags.

	## starts the waiting, use flags to tweak the timeout
	$ waiter start

	## when done, the lock-file must be removed
	$ waiter done
	## (or)
	$ rm -f <lock-file-path>

In the case of timeout, the waiter will return error, it only exists gracefully via "waiter done", or
the removal of the lock-file before timeout.
`

var (
	// lockFilePath path to lock file
	lockFilePath string
	// retries amount of retries
	retries int64
	// interval amount of time to sleep bettween attempts
	interval time.Duration
)

func init() {
	flags := rootCmd.PersistentFlags()

	flags.StringVar(&lockFilePath, "lock-file-path", "/var/tmp/waiter.pid", "path to the lock-file")
	flags.Int64Var(&retries, "retries", 120, "amount of attempts")
	flags.DurationVar(&interval, "interval", 1*time.Second, "sleep between attempts")
}

// newWaiter based on variables backing up command-line flags.
func newWaiter() *Waiter {
	return NewWaiter(lockFilePath, retries, interval)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("[ERROR] %v", err)
		os.Exit(1)
	}
	os.Exit(0)
}
