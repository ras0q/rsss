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
		db, err := database.InitDB("rsss.db")
		if err != nil {
			return fmt.Errorf("failed to initialize database: %w", err)
		}
		defer db.Close()

		url := args[0]
		if err := database.AddFeed(db, url); err != nil {
			return fmt.Errorf("failed to add feed: %w", err)
		}

		fmt.Printf("Feed '%s' added successfully.\n", url)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
