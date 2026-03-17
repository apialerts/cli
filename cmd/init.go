package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/apialerts/cli/internal/config"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var initKey string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Set up your API key",
	Long:  "Set your API key interactively or via flag. The key is stored in ~/.apialerts/config.json.",
	Example: `  apialerts init
  apialerts init --key "your-api-key"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		key := initKey

		if key == "" {
			if !term.IsTerminal(int(os.Stdin.Fd())) {
				return fmt.Errorf("no terminal detected — use: apialerts init --key \"your-api-key\"")
			}
			fmt.Print("Enter your API key: ")
			keyBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
			fmt.Println()
			if err != nil {
				return fmt.Errorf("failed to read input: %w", err)
			}
			key = strings.TrimSpace(string(keyBytes))
		}

		if key == "" {
			return fmt.Errorf("API key cannot be empty")
		}

		cfg := &config.CLIConfig{APIKey: key}
		if err := config.Save(cfg); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}

		fmt.Println("API key saved.")
		return nil
	},
}

func init() {
	initCmd.Flags().StringVar(&initKey, "key", "", "Your API Alerts API key")
	rootCmd.AddCommand(initCmd)
}
