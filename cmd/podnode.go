package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/gabrie30/kuve/colorlog"
	"github.com/spf13/cobra"
)

// podnodeCmd represents the exec command
var podnodeCmd = &cobra.Command{
	Use:   "podnode [pod_name] [namespace]",
	Short: "View which node a given pod in a given namespace is running on (gcp clusters only)",
	Long:  `View which node a given pod in a given namespace is running on (gcp clusters only)`,
	Run: func(cmd *cobra.Command, args []string) {
		podNode(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(podnodeCmd)
}

func podNode(pod string, namespace string) {
	cmd := fmt.Sprintf("kubectl get pod/%s --namespace=%s -o custom-columns=NODE:.spec.nodeName --no-headers", pod, namespace)
	f, _ := exec.Command("/bin/sh", "-c", cmd).Output()

	nodeip := strings.Split(string(f), "\n")
	nodeip = removeEmptyString(nodeip)

	if len(nodeip) < 1 {
		colorlog.Error(fmt.Sprintf("No node found for pod: %v in namespace: %s\n", pod, namespace))
		colorlog.Info("Check that your context, pod name and namespace are correct")
		return
	}
	colorlog.Success(fmt.Sprintf("Pod %s is on node %s", pod, strings.Join(nodeip, ", ")))
	colorlog.SubtleInfo("------------------------------------------------ NODE DEBUG INFO ------------------------------------------------")

	customstring := "-o=custom-columns=NAME:.metadata.name,INSTANCE-TYPE:.metadata.labels.'beta\\.kubernetes\\.io/instance-type',NODEPOOL:.metadata.labels.'cloud\\.google\\.com/gke-nodepool',INTERNAL-IP:'.status.addresses[0].address'"
	ni := fmt.Sprintf("kubectl get node %s %s", strings.Join(nodeip, ", "), customstring)
	do, _ := exec.Command("/bin/sh", "-c", ni).Output()
	fmt.Printf("%s", do)
}

func removeEmptyString(s []string) []string {
	retVal := []string{}
	for _, item := range s {
		if item != "" {
			retVal = append(retVal, item)
		}
	}
	return retVal
}
