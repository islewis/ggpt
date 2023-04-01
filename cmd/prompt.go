/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"context"
	"os"
	"strings"
	"log"
	"github.com/spf13/cobra"
	openai "github.com/sashabaranov/go-openai"
)

// promptCmd represents the prompt command
var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		home, _ := os.UserHomeDir()
                credPath := home + "/.ggpt/credentials"
                _, err := os.Stat(credPath)
		// Check credential file exists. Could do some more validation to check its a legit key here
		if os.IsNotExist(err) {
			fmt.Print("OpenAI API key not found. Configure key by running 'ggpt configure'")
		}
		// Run 
                if err == nil {
                        // Read in API key
                        fileContents, err := os.ReadFile(credPath)
                        if err != nil {log.Fatal(err)}
                        fileSplit := strings.Split(string(fileContents), "=")
                        key :=  fileSplit[1]
			// Get completion
			client := openai.NewClient(key)
			resp, err := client.CreateChatCompletion(
				context.Background(),
				openai.ChatCompletionRequest{
					Model: openai.GPT3Dot5Turbo,
					Messages: []openai.ChatCompletionMessage{
						{
							Role:    openai.ChatMessageRoleUser,
							Content: args[0],
						},
					},
				},
			)
			if err == nil{
				// print output
				output := resp.Choices[0].Message.Content + "\n" 
				fmt.Print(output)
			}
			if err != nil {log.Fatal(err)}
                }
	},
}

func init() {
	rootCmd.AddCommand(promptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// promptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// promptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
