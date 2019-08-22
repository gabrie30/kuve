package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	eventType string
	allTypes  bool
)

// eventsCmd represents the events command
var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Get and filter events based off type from current context",
	Long:  `Get and filter events based off type from current context`,
	Run: func(cmd *cobra.Command, args []string) {
		filterEvents()
	},
}

func init() {
	rootCmd.AddCommand(eventsCmd)
	eventsCmd.Flags().StringVarP(&eventType, "type", "c", "Warning", "type of event to filter")
	eventsCmd.Flags().BoolVarP(&allTypes, "all", "a", false, "show all events")
}

func filterEvents() {

	var cmdStr string
	if allTypes == true {
		cmdStr = fmt.Sprintf("kubectl get events --all-namespaces")
	} else {
		cmdStr = fmt.Sprintf("kubectl get events --all-namespaces --field-selector type=%s", eventType)
	}

	out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(out))
}
