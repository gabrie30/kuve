package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Get logs from pods and containers in a given namespace",
	Long:  `Get logs from pods and containers in a given namespace. If using namespaces to find pods, it will use the first pod in the namespace. If labels it will look for the first pod in all namespaces with that label given`,
	Run: func(cmd *cobra.Command, args []string) {
		getLogs(args[0])
	},
}

func getLogs(arg string) {
	var cmd string
	switch findBy := viper.Get("FIND_PODS_BY"); findBy {
	case "labels":
		pod, namespace := getPodAndNamespaceFromLabel(arg)
		cmd = fmt.Sprintf("kubectl logs -f %s --all-containers=true -n %s --limit-bytes=1000000", pod, namespace)
	case "namespaces":
		podCmd := fmt.Sprintf("kubectl get pods -n %s --no-headers | awk -v x=1 '$4 >= x' | awk '{print $1}' | head -1", arg)
		pod, _ := exec.Command("/bin/sh", "-c", podCmd).Output()

		s := strings.Trim(string(pod), "\n")
		if string(s) == "No resources found." || len(pod) == 0 {
			fmt.Printf("Coud not find pods in namespace: %s", s)
			return
		}
		cmd = fmt.Sprintf("kubectl logs -f %s --all-containers=true -n %s --limit-bytes=1000000", s, arg)
	default:
		log.Fatalf("You must set FIND_PODS_BY in .kuve.yaml to namespaces or labels")
	}

	result, _ := exec.Command("/bin/sh", "-c", cmd).Output()
	fmt.Println(string(result))
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
