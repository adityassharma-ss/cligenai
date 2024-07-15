package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

var numWords string

var searchCmd = &cobra.Command{
	Use:   "search [your question]",
	Short: "Ask a question and get a response",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		response, err := getApiResponse(args)
		if err != nil {
			log.Fatalf("Error getting API response: %v", err)
		}
		fmt.Println(response)
	},
}

func init() {
	searchCmd.Flags().StringVarP(&numWords, "words", "w", "150", "Number of words in the response")
}

func getApiResponse(args []string) (string, error) {
	userArgs := strings.Join(args, " ")

	num, err := strconv.Atoi(numWords)
	if err != nil || num <= 0 {
		return "", fmt.Errorf("invalid number of words: %s", numWords)
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return "", fmt.Errorf("failed to create GenAI client: %w", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(fmt.Sprintf("%s in %d words.", userArgs, num)))
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no response candidates received")
	}

	finalResponse := resp.Candidates[0].Content.Parts[0]
	return fmt.Sprint(finalResponse), err
}

func checkNilError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
