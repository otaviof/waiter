package main

import (
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:          "start",
	Short:        "Starts the waiter daemon, and holds until `done` is issued.",
	SilenceUsage: true,
	RunE:         start,
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func start(cmd *cobra.Command, args []string) error {
	w := newWaiter()
	return w.Wait()
}
