package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cligenai",
	Short: "A CLI tool to interact using the Gemini API",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			log.Printf("Error displaying help: %v", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		log.Printf("Error executing root command: %v", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
