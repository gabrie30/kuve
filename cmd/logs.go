package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Get logs from pods and containers in a given namespace",
	Long:  `Get logs from pods and containers in a given namespace`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logs called")
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
