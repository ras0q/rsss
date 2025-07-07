package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/ras0q/rsss/internal/database"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rsss",
	Short: "A CLI tool to fetch, process, and summarize RSS feeds.",
	Long: `rsss is a CLI tool that periodically fetches new articles from RSS feeds,
summarizes them using Gemini, and stores the results.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		dbType, _ := cmd.Flags().GetString("database-type")
		dsn, _ := cmd.Flags().GetString("database-dsn")

		var db database.DB
		var err error

		switch dbType {
		case "sqlite":
			db, err = database.NewSQLiteDB(dsn)
		case "mysql":
			db, err = database.NewMySQLDB(dsn)
		default:
			return fmt.Errorf("unsupported database type: %s", dbType)
		}

		if err != nil {
			return fmt.Errorf("failed to initialize database: %w", err)
		}

		cmd.SetContext(database.ToCtx(cmd.Context(), db))

		return nil
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		db, ok := database.FromCtx(cmd.Context())
		if ok && db != nil {
			if err := db.Close(); err != nil {
				cmd.PrintErrf("failed to close database: %v", err)
			}
		}
	},
}

func Execute(ctx context.Context) {
	err := fang.Execute(ctx, rootCmd)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("database-type", "sqlite", "Database type (sqlite or mysql)")
	rootCmd.PersistentFlags().String("database-dsn", "rsss.db", "Database DSN")
}
