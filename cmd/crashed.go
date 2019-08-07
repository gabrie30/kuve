package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/gabrie30/kuve/colorlog"
	"github.com/spf13/cobra"
)

// crashedCmd represents the crashed command
var crashedCmd = &cobra.Command{
	Use:   "crashed [namespace]",
	Short: "Get logs from crashed pods in a given namespace, uses --previous command",
	Long:  `Get logs from crashed pods in a given namespace, uses --previous command`,
	Run: func(cmd *cobra.Command, args []string) {
		namespace := args[0]
		getCrashedLogs(namespace)
	},
}

func init() {
	logsCmd.AddCommand(crashedCmd)
	crashedCmd.Flags().StringVarP(&container, "container", "c", "", "container to grab logs from (default is first container found)")
}

func getCrashedLogs(namespace string) {

	// Iterate over those pods getting their logs
	fmt.Println("")

	targetContainer := container

	if targetContainer == "" {
		targetContainer = getFirstContainer(namespace)
	}

	// Get a list of pods with restarts
	podsWithRestarts := podsWithRestarts(namespace)
	restartingPods := strings.Split(podsWithRestarts, "\n")

	for _, pod := range restartingPods {

		if pod == "\n" || len(pod) <= 0 {
			continue
		}

		s := fmt.Sprintf("kubectl logs %s -n %s -c %s --previous", pod, namespace, targetContainer)
		container, _ := exec.Command("/bin/sh", "-c", s).Output()

		colorlog.Info(fmt.Sprintf("--------------------------- pod: %s container: %s ---------------------------", pod, targetContainer))
		fmt.Println(string(container))
	}
}

func podsWithRestarts(namespace string) string {
	rsPodsString := fmt.Sprintf("kubectl get pods -n %s --no-headers | awk -v x=1 '$4 >= x' | awk '{print $1}'", namespace)

	pods, err := exec.Command("/bin/sh", "-c", rsPodsString).Output()
	if err != nil {
		fmt.Println(err)
	}

	if len(pods) <= 0 {
		log.Fatalf("Cannot find pods with restarts in namespace: %s", namespace)
	}

	return string(pods)
}

func getFirstContainer(namespace string) string {

	containerString := fmt.Sprintf("kubectl get pods -n %s -o jsonpath='{.items[0].spec.containers[0].name}'", namespace)

	container, err := exec.Command("/bin/sh", "-c", containerString).Output()
	if err != nil {
		fmt.Println(err)
	}

	if len(container) <= 0 {
		log.Fatalf("Cannot find container in namespace: %s", namespace)
	}

	return string(container)
}
