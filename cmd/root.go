package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rsss",
	Short: "A CLI tool to fetch, process, and summarize RSS feeds.",
	Long: `rsss is a CLI tool that periodically fetches new articles from RSS feeds,
summarizes them using Gemini, and stores the results.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
