package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "waiter [flags]",
	SilenceUsage: true,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("[ERROR] '%#v'", err)
		os.Exit(1)
	}
	os.Exit(0)
}
