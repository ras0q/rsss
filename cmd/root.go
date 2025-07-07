package cmd

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rsss",
	Short: "A CLI tool to fetch, process, and summarize RSS feeds.",
	Long: `rsss is a CLI tool that periodically fetches new articles from RSS feeds,
summarizes them using Gemini, and stores the results.`,
}

func Execute(ctx context.Context) {
	err := fang.Execute(ctx, rootCmd)
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
