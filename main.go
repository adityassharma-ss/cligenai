package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adityassharma-ss/cligenai/cmd"
)

func getEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("environment variable %s not set", key)
	}
	return value, nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	geminiAPIKey, err := getEnv("GEMINI_API_KEY")
	if err != nil {
		log.Fatalf("Error: %v. Please set the GEMINI_API_KEY environment variable. Check the README for more information.", err)
	}

	// Execute the root command and handle errors
	cmd.Execute()
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}

	fmt.Println("GEMINI_API_KEY:", geminiAPIKey) // Optional: For debugging purposes, can be removed in production
}
