/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"context"
	"encoding/json"
	"os"
	"strconv"
	"time"
	"strings"
	"log"
	"github.com/spf13/cobra"
	"github.com/islewis/ggpt/common"
	openai "github.com/sashabaranov/go-openai"
)

// promptCmd represents the prompt command
var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Call GPT autocomplete with the given string as a prompt",
	Long: `This command is the meat of ggpt. Pass in a prompt, get an output from GPT.

For example:
	'ggpt prompt "Write me a haiku"'

The output of this command can be piped out, allowing for flexible manipulation directly in the terminal.
	'ggpt prompt "Output a sample csv of 5 apartments, including cost and location" | tee apartments.csv'

Command substition allows for full integration into CLI workflows. 
	'ggpt prompt "Briefly summarize the content of the following csv: $(cat apartments.csv)"'
`,
	Run: func(cmd *cobra.Command, args []string) {
		// read in debug arg, if exists
		home, _ := os.UserHomeDir()
                credPath := home + "/.ggpt/credentials"
		// Check credential file exists. Could do some more validation to check its a legit key here
		if Verbose == true {
			fmt.Println("Looking for API credentials at " + credPath)
		}
                _, err := os.Stat(credPath)
		if os.IsNotExist(err) {
			fmt.Print("OpenAI API key not found. Configure key by running 'ggpt configure'\n")
		}
		// Run 
                if err == nil {
                        // Read in API key
			if Verbose == true {
				fmt.Println("Credential file found at " + credPath)
			}
                        fileContents, err := os.ReadFile(credPath)
                        if err != nil {log.Fatal(err)}
                        fileSplit := strings.Split(string(fileContents), "=")
                        key :=  fileSplit[1]
			// Get completion
			if Verbose == true {
				fmt.Println("\nMaking GPT query")
				fmt.Printf("PROMPT: '%s'\n\n", args[0])
			}
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
				if Verbose == true {
					fmt.Println("Query returned successfully")
				}
				// Print output
				output := resp.Choices[0].Message.Content
				fmt.Println("\nOUTPUT:")
				fmt.Println(output+"\n")
				// Log request
				currentTime := time.Now().Unix()
				data := common.Record {
					Time : currentTime,
					Prompt : args[0],
					Output : output,
				}
				file, _ := json.MarshalIndent(data, "", " ")
				recordPath := home+"/.ggpt/history/"+strconv.FormatInt(time.Now().Unix(),10)+".json"
				if Verbose == true {
					fmt.Println("Storing query record locally at " + recordPath)
				}
				err = ioutil.WriteFile(recordPath, file, 0644)
				if err != nil {log.Fatal(err)}
			if err != nil {log.Fatal(err)}
			}
		if err != nil {log.Fatal(err)}
                }
	},
}

var Verbose bool

func init() {
	rootCmd.AddCommand(promptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// promptCmd.PersistentFlags().String("foo", "", "A help for foo")
	promptCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// promptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
