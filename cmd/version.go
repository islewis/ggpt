package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "v0.2.0"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the version number",
	Long:  "Display the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ggpt", version)
	},
}
