package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func main() {
	// Load any variables from a local .env file (optional)
	godotenv.Load()

	// Fetch your OpenAI API key from the environment
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is not set")
	}

	systemContext := os.Getenv()

	// Create a new OpenAI client
	client := openai.NewClient(apiKey)

	// Prepare the prompt we want to send to the LLM
	hard_prompt := "Hello, can you provide a brief overview of Go's concurrency model?"

	// Create a ChatCompletion request. 
	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo, // or "gpt-4" if you have access
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleSystem,
				Content: 
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})
	if err != nil {
		log.Fatalf("OpenAI API error: %v", err)
	}

	// Print the first choice returned by the LLM
	if len(resp.Choices) > 0 {
		fmt.Println("LLM Response:")
		fmt.Println(resp.Choices[0].Message.Content)
	} else {
		fmt.Println("No response from LLM")
	}
}