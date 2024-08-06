package ollama

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"log"
	"os/exec"
	"time"
)

var (
	llm    llms.LLM
	ctx    context.Context
	cancel context.CancelFunc
)

/**
 * InitOllama initializes the Ollama client with the specified model.
 */
func InitOllama(model string) error {
	var err error
	llm, err = ollama.New(ollama.WithModel(model))
	if err != nil {
		return err
	}

	ctx = context.Background()
	return nil
}

/**
 * LaunchOllama starts the Ollama server in the background.
 */
func LaunchOllama() error {
	ctx, cancel = context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "ollama", "serve")
	err := cmd.Start()
	if err != nil {
		cancel()
		return fmt.Errorf("failed to start Ollama: %v", err)
	}

	go func() {
		err = cmd.Wait()
		if err != nil {
			log.Printf("Ollama exited with error: %v", err)
		}
	}()

	// Wait for a few seconds to ensure the server has started
	time.Sleep(5 * time.Second)
	return nil
}

/**
 * QueryOllama sends a query to the Ollama server and returns the response.
 */
func QueryOllama(prompt string) (map[string]interface{}, error) {
	response, err := llm.Call(ctx, prompt)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{
		"response": response,
	}

	return result, nil
}
