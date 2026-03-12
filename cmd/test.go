package cmd

import (
	"fmt"

	"github.com/apialerts/apialerts-cli/internal/config"
	"github.com/apialerts/apialerts-go"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Send a test event",
	Long:  "Send a test event to verify your API key and connectivity.",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey, err := config.GetAPIKey()
		if err != nil {
			return err
		}

		apialerts.Configure(apiKey)
		apialerts.SetIntegration("cli")

		event := apialerts.Event{
			Event:   "cli.test",
			Title:   "CLI Test Event",
			Message: "Test event from API Alerts CLI",
			Tags:    []string{"test", "cli"},
		}

		fmt.Println("Sending test event...")
		if err := apialerts.SendAsync(event); err != nil {
			return fmt.Errorf("test failed: %w", err)
		}

		fmt.Println("Test event sent successfully.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
