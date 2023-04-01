/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"fmt"
	"log"
	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ggpt",
	Short: "CLI utility for GPT completion, written in go",
	Long: `ggpt is a CLI tool to interact with OpenAI's GPT language model. ggpt wraps OpenAI's completion feature, via their API, outputting the result directly in the terminal. `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ggpt.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// make .ggpt dir
	fmt.Print("inside init config")
	home, _ := os.UserHomeDir()
        dirPath := home + "/.ggpt"
        _, err := os.Stat(dirPath)
        if err == nil { return }
	if os.IsNotExist(err) {
		_ = os.MkdirAll(dirPath, os.ModePerm)
		fmt.Print("making dir")
        if err != nil {log.Fatal(err)}
	}
	// make .ggpt history dir
        histPath := dirPath + "/history"
        _, err = os.Stat(histPath)
        if err == nil { return }
	if os.IsNotExist(err) {
		n := os.MkdirAll(histPath, os.ModePerm)
		fmt.Print("making history")
		fmt.Print(n)
		if err != nil {log.Fatal(err)}
	}
}
