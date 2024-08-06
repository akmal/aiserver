package main

import (
	"flag"
	"fmt"
	"github.com/akmal/aiserver/ollama"
	"github.com/akmal/aiserver/server"
	"os"
)

func main() {
	modelName := flag.String("model", "llama3.1", "Name of the LLM model to use")
	debug := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	err := ollama.InitOllama(*modelName)
	if err != nil {
		fmt.Printf("Error initializing Ollama: %v\n", err)
		os.Exit(1)
	}

	err = ollama.LaunchOllama()
	if err != nil {
		fmt.Printf("Error launching Ollama: %v\n", err)
		os.Exit(1)
	}

	server.StartServer(*debug)
}
