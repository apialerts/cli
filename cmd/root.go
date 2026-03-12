package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "apialerts",
	Short:   "API Alerts CLI — send events from your terminal",
	Long:    "A command-line interface for apialerts.com. Configure your API key, send events, and test connectivity.",
	Version: Version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
