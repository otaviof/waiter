package main

import (
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:          "done",
	Aliases:      []string{"stop"},
	Short:        "Interrupts the waiting.",
	SilenceUsage: true,
	RunE:         done,
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

func done(cmd *cobra.Command, args []string) error {
	p := newWaiter()
	return p.Done()
}
