package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/ras0q/rsss/internal/database"
	"github.com/ras0q/rsss/internal/rss"
	"github.com/ras0q/rsss/internal/summarizer"
	"github.com/spf13/cobra"
)

var summarizeCmd = &cobra.Command{
	Use:   "summarize",
	Short: "Summarize new articles from all RSS feeds",
	RunE: func(cmd *cobra.Command, args []string) error {
		model, _ := cmd.Flags().GetString("model")
		prompt, _ := cmd.Flags().GetString("prompt")
		apiKey := os.Getenv("GEMINI_API_KEY")
		if apiKey == "" {
			return fmt.Errorf("GEMINI_API_KEY environment variable not set")
		}

		db, err := database.InitDB("rsss.db")
		if err != nil {
			return fmt.Errorf("failed to initialize database: %w", err)
		}
		defer db.Close()

		feeds, err := database.GetFeeds(db)
		if err != nil {
			return fmt.Errorf("failed to get feeds: %w", err)
		}

		for _, feedURL := range feeds {
			fmt.Printf("Processing feed: %s\n", feedURL)
			feed, err := rss.ParseFeed(feedURL)
			if err != nil {
				log.Printf("Failed to parse feed %s: %v", feedURL, err)
				continue
			}

			for _, item := range feed.Items {
				processed, err := database.IsArticleProcessed(db, item.GUID)
				if err != nil {
					log.Printf("Failed to check if article is processed: %v", err)
					continue
				}

				if !processed {
					fmt.Printf("  Summarizing new article: %s\n", item.Title)

					contentToSummarize := item.Description
					if item.Content != "" {
						contentToSummarize = item.Content
					}

					summary, err := summarizer.Summarize(apiKey, model, prompt, contentToSummarize)
					if err != nil {
						log.Printf("    Failed to summarize: %v", err)
						continue
					}

					fmt.Printf("    Summary: %s\n", summary)

					if err := database.MarkArticleAsProcessed(db, item.GUID); err != nil {
						log.Printf("    Failed to mark article as processed: %v", err)
					}
				}
			}
		}
		return nil
	},
}

func init() {
	summarizeCmd.Flags().String("model", "gemini-1.5-flash", "The model to use for summarization")
	summarizeCmd.Flags().String("prompt", "次の記事を日本語で3文に要約してください。", "The prompt to use for summarization")
	rootCmd.AddCommand(summarizeCmd)
}
