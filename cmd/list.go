package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"rsss/internal/database"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all RSS feeds",
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := database.InitDB("rsss.db")
		if err != nil {
			return fmt.Errorf("failed to initialize database: %w", err)
		}
		defer db.Close()

		feeds, err := database.GetFeeds(db)
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
