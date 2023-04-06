/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"encoding/json"
	"strings"
	"os"
	"io/ioutil"
	"strconv"
	"io/fs"
	"path/filepath"
	"github.com/spf13/cobra"
	"github.com/islewis/ggpt/common"
)

// lastCmd represents the last command
var lastCmd = &cobra.Command{
	Use:   "last",
	Short: "Returns the output of the previous query.",
	Long: `Last repeats the output of the previous GPT query.
In comparison to querying again with the same prompt,
last guarentees the previous output, and prevents
unnecessary usage of OpenAI's API. This saves time
and money when a simple reuse is desired.
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Make sure theres a record to use
		home, _ := os.UserHomeDir()
		histDir := home+"/.ggpt/history"
		// Get list of files, via timestamped name
		var records []int64
		filepath.WalkDir(histDir, func(s string, d fs.DirEntry, err error) error {
			if err != nil {
				fmt.Println("Unable to find any previous queries.")
				log.Fatal(err)
			}
			if filepath.Ext(d.Name()) == ".json" {
				file := strings.Split(s, "/")
				timestamp := strings.Split(file[len(file)-1], ".")
				// take first slice, conv to int64
				timestampI64, err := strconv.ParseInt(timestamp[0], 10, 64)
				if err != nil {log.Fatal(err)}
				records = append(records, timestampI64)
			}
			return nil
		})
		// check there is a record to return
		if len(records) == 0 {
			fmt.Println("No previous queries found. Make one with \"ggpt prompt\"")
			os.Exit(0)
		}
		// get latest record
		latest := records[0]
		for _, value := range records {
			if value > latest {
			latest = value
			}
		}
		// return output of latest record
		latestPath := histDir + "/" + strconv.FormatInt(latest, 10) + ".json"
		jsonFile, err := os.Open(latestPath)
		if err != nil {log.Fatal(err)}
		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {log.Fatal(err)}
		var record common.Record
		json.Unmarshal(byteValue, &record)
		fmt.Println(record.Output)
	},
}

func init() {
	rootCmd.AddCommand(lastCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lastCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lastCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
