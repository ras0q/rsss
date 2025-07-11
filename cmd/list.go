package cmd

import (
	"fmt"

	"github.com/ras0q/rsss/internal/database"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all RSS feeds",
	RunE: func(cmd *cobra.Command, args []string) error {
		db, ok := database.FromCtx(cmd.Context())
		if !ok {
			return fmt.Errorf("database not found in context")
		}

		feeds, err := db.GetFeeds(cmd.Context())
		if err != nil {
			return fmt.Errorf("failed to get feeds: %w", err)
		}

		if len(feeds) == 0 {
			fmt.Println("No feeds found.")
			return nil
		}

		fmt.Println("Registered feeds:")
		for _, feed := range feeds {
			fmt.Printf("- %s\n", feed)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
