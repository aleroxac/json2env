package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var file string

func convertToENV(c *cobra.Command, args []string) error {
	if file == "" {
		fmt.Println("Please, inform some file to be parsed.")
		os.Exit(0)
	}

	file_content, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error to read file: %s", err)
	}

	var struct_content []map[string]interface{}
	if err := json.Unmarshal(file_content, &struct_content); err != nil {
		fmt.Printf("Error to unmarshal file: %s", err)
	}

	for _, value := range struct_content {
		for key, value := range value {
			fmt.Printf("%s=%s\n", key, value)
		}
	}

	return nil
}

var rootCmd = &cobra.Command{
	Use:   "json2env",
	Short: "Command line tool to convert JSON into envfile.",
	RunE:  convertToENV,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&file, "from-file", "f", "", "JSON file to be converted into envfile")
}
