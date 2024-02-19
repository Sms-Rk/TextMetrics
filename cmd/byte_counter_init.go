package cmd

import (
        "github.com/spf13/cobra"
		"TextMetrics/cmd/helper"
)

// deleteCmd represents the delete command
var byteCountCmd = &cobra.Command{
        Use:   "bytes",
        Short: "A brief description of your command",
        Long: "init function for couting file bytes or file sizes",
        Run: func(cmd *cobra.Command, args []string) {
			helper.GetByteCount("go.mod")
        },
}

func init() {
        rootCmd.AddCommand(byteCountCmd)

}