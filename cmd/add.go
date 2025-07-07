package cmd

import (
	"fmt"

	"github.com/ras0q/rsss/internal/database"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [url]",
	Short: "Add a new RSS feed",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		db, ok := database.FromCtx(cmd.Context())
		if !ok {
			return fmt.Errorf("database not found in context")
		}

		url := args[0]
		if err := db.AddFeed(cmd.Context(), url); err != nil {
			return fmt.Errorf("failed to add feed: %w", err)
		}

		fmt.Printf("Feed '%s' added successfully.\n", url)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
