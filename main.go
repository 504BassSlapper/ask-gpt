package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PullRequestInc/go-gpt3"
)

type Config struct {
	OpenAIAPIKey string
	DefaultModel string
}

func main() {
	apiKey := os.Getenv("OPENAI_PRIVATE_KEY")
	if apiKey == "" {
		log.Fatalln("Missing API KEY")
	}
	//options := gp3.ClientOption
	ctx := context.Background()
	config := Config{
		OpenAIAPIKey: os.Getenv("OPENAI_PRIVATE_KEY"),
		DefaultModel: "gpt-3.5-turbo-instruct",
	}
	httpClient := &http.Client{

		Timeout: 0,
	}
	client := gpt3.NewClient(config.OpenAIAPIKey, gpt3.WithHTTPClient(httpClient), gpt3.WithDefaultEngine(config.DefaultModel))
	request := gpt3.CompletionRequest{
		Prompt:    []string{"How many coffee should i drink today"},
		MaxTokens: gpt3.IntPtr(50),
		Stop:      []string{",", "."},
		Echo:      false,
	}
	fmt.Printf("Question: \n%s\n", request.Prompt[0])
	resp, err := client.Completion(ctx, request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Response \n %s", resp.Choices[0].Text)
	}
}
