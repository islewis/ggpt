package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		dirPath := home+"/.ggpt"
		credPath := dirPath + "/credentials"
		_, err := os.Stat(credPath)
		// if credential file isnt set, set it
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
			if err != nil {log.Fatal(err)}
			fmt.Println("Key set")
		}
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
			fmt.Print("If new key is entered, previous key will be overwritten.\n")
			fmt.Printf("OpenAI API Key [%s]: ", censoredKey)
			fmt.Scanln(&newKey)
			// Do some very validation on new key
			if len(newKey) == 51 {
				f, err := os.Create(credPath)
				if err != nil {log.Fatal(err)}
				newFileContents := "openai_key=" + newKey
				_, err = f.WriteString(newFileContents)
				if err != nil {log.Fatal(err)}
				fmt.Println("Key replaced")
			} else {
				fmt.Println("Input doesnt look like an OpenAI API key, try again")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	// set config info
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	home, _ := os.UserHomeDir()
	dirPath := home+"/.ggpt"
	histPath := dirPath+"/history"
	viper.AddConfigPath(dirPath)
	err := viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		// make sure dirs are created
		_= os.MkdirAll(histPath, os.ModePerm)
		// todo handle error
		// write file with any defaults
		viper.SetDefault("model_name", "GPT3Dot5Turbo")
		viper.SafeWriteConfig()
		fmt.Print("ggpt config files initialized at ~/.ggpt \n")
	} else {
		return
		// Config file was found but another error was produced
	}
}

