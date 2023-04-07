package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/islewis/ggpt/common"
	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// promptCmd represents the prompt command
var promptCmd = &cobra.Command{
	Use:   "prompt \"This is my prompt.\"",
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
		// make sure theres a prompt
		if len(args) == 0 {
			fmt.Println("Prompt not found. See \"ggpt prompt --help\" for more details")
			os.Exit(0)
		}
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
		if err == nil {
			// Read in API key
			if Verbose == true {
				fmt.Println("Credential file found at " + credPath)
			}
			fileContents, err := os.ReadFile(credPath)
			if err != nil {
				log.Fatal(err)
			}
			fileSplit := strings.Split(string(fileContents), "=")
			key := fileSplit[1]
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
			if err == nil {
				// Print output
				output := resp.Choices[0].Message.Content
				if Verbose == true {
					fmt.Println("Query returned successfully")
					fmt.Println("\nOUTPUT:")
				}
				fmt.Println(output)
				// Log request
				currentTime := time.Now().Unix()
				data := common.Record{
					Time:   currentTime,
					Prompt: args[0],
					Output: output,
				}
				file, _ := json.MarshalIndent(data, "", " ")
				recordPath := home + "/.ggpt/history/" + strconv.FormatInt(time.Now().Unix(), 10) + ".json"
				if Verbose == true {
					fmt.Println("\nStoring query record locally at " + recordPath)
				}
				err = ioutil.WriteFile(recordPath, file, 0644)
				if err != nil {
					log.Fatal(err)
				}
				if err != nil {
					log.Fatal(err)
				}
			}
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

var Verbose bool

func init() {
	rootCmd.AddCommand(promptCmd)
	cobra.OnInitialize(checkHistoryDir)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// promptCmd.PersistentFlags().String("foo", "", "A help for foo")
	promptCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// promptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func checkHistoryDir() {
	home, _ := os.UserHomeDir()
	dirPath := home + "/.ggpt"
	histPath := dirPath + "/history"
	_, err := os.Stat(histPath)
	if err != nil {
		if os.IsNotExist(err) {
			_ = os.MkdirAll(histPath, os.ModePerm)
		} else {
			fmt.Println("Unable to check if history dir exists at" + dirPath)
			log.Fatal(err)
		}
	}
}
