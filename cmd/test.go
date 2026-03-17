package cmd

import (
	"fmt"

	"github.com/apialerts/cli/internal/config"
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

		apialerts.ConfigureWithConfig(apiKey, apialerts.Config{Debug: true})
		apialerts.SetIntegration(IntegrationName)

		event := apialerts.Event{
			Event:   "cli.test",
			Title:   "CLI Test Event",
			Message: "Test event from API Alerts CLI",
			Tags:    []string{"test", "cli"},
		}

		result, err := apialerts.SendAsync(event)
		if err != nil {
			return fmt.Errorf("test failed: %w", err)
		}

		fmt.Printf("✓ Test event sent to %s (%s)\n", result.Workspace, result.Channel)
		for _, w := range result.Warnings {
			fmt.Printf("! Warning: %s\n", w)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
