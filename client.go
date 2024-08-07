package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run client.go <input_file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Sending prompt...\n")
	resp, err := http.Post("http://localhost:8080/prompt", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error response from server: %v, %s\n", resp.Status, resp.Body)
		os.Exit(2)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error closing response body: %v\n", err)
		}
	}(resp.Body)

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("Error decoding response: %v\n", err)
		os.Exit(1)
	}

	response, ok := result["response"].(string)
	if !ok {
		fmt.Println("Invalid response format")
		os.Exit(1)
	}

	fmt.Println(response)
}
