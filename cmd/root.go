package cmd

import(
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"TextMetrics/cmd/helper"
)

var rootCmd = &cobra.Command{
	Use:   "tm",
	Short: "TextMetrics is simple cmd to retrive file content features",
	Long: `A simple command line tools to get any features about file content`,
	Run: func(cmd *cobra.Command, args []string) {
	  helper.GetLineCount("go.mod")
	},
  }
  
  func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
  }