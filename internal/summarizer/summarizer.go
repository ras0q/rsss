package summarizer

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func Summarize(apiKey, modelName, prompt, text string) (string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel(modelName)
	fullPrompt := fmt.Sprintf("%s\n\n%s", prompt, text)
	resp, err := model.GenerateContent(ctx, genai.Text(fullPrompt))
	if err != nil {
		return "", err

	}

	if len(resp.Candidates) > 0 {
		for _, part := range resp.Candidates[0].Content.Parts {
			if txt, ok := part.(genai.Text); ok {
				return string(txt), nil
			}
		}
	}

	return "", fmt.Errorf("no summary generated")
}
