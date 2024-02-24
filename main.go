package main

import (
	"bufio"
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
		N:         gpt3.IntPtr(1),
		// Stop:      []string{",", "."},
		Echo: false,
	}
	fmt.Printf("Question: \n%s\n", request.Prompt[0])
	resp, err := client.Completion(ctx, request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Response \n %s", resp.Choices[0].Text)
	}

	for true {
		fmt.Println("Ask a question\n\n>")
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		request := gpt3.CompletionRequest{
			Prompt:    []string{line},
			MaxTokens: gpt3.IntPtr(50),
			N:         gpt3.IntPtr(1),
			// Stop:      []string{",", "."},
			Echo: false,
		}
		response, _ := client.Completion(ctx, request)
		fmt.Printf("Response \n %s \n", response.Choices[0].Text)
	}
}
