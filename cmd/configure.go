/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure your OpenAI API key",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		home, _ := os.UserHomeDir()
		credPath := home + "/.ggpt/credentials"
		_, err := os.Stat(credPath)
		if err == nil {
			// Read in current API key
			fileContents, err := os.ReadFile(credPath)
		        if err != nil {log.Fatal(err)}
			fileSplit := strings.Split(string(fileContents), "=")
			key :=  fileSplit[1]
			// Get blurred key to show
			censoredKey := "**************" + key[len(key)-5:]
			// Prompt for new key
			var newKey string
			fmt.Printf("OpenAI API Key [%s]: ", censoredKey)
			fmt.Scanln(&newKey)
			// Overwrite existing file
			f, err := os.Create(credPath)
		        if err != nil {log.Fatal(err)}
			newFileContents := "openai_key=" + newKey
			_, err = f.WriteString(newFileContents)
			if err != nil {log.Fatal(err)}
			fmt.Println("Key replaced")


		}
		if os.IsNotExist(err) {
			// create credential file
			f, err := os.Create(credPath)
		        if err != nil {log.Fatal(err)}
			// take prompt for key
			var newKey string
			fmt.Print("OpenAI API Key: ")
			fmt.Scanln(&newKey)
			// write key to credential file
			fileContents := "openai_key=" + newKey
			_, err = f.WriteString(fileContents)
			//fmt.Fprintf(credPath, key)
			if err != nil {log.Fatal(err)}
			fmt.Println("Key set")
		}
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
