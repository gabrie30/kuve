package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	shell     string
	container string
	label     string
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec [namespace/label value]",
	Short: "Execs into the first running pod and container",
	Long:  `Execs into the first running pod and container. If you use the label flag it will over write what is in your .kuve.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		execPod(args[0])
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
	execCmd.Flags().StringVarP(&shell, "shell", "s", "/bin/sh", "shell to exec with")
	execCmd.Flags().StringVarP(&container, "container", "c", "", "container to exec into")
	execCmd.Flags().StringVarP(&label, "selector", "l", "", "selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)")
}

func execPod(arg string) {
	var execCmd string

	pod := getPod(arg)

	switch fetchType := viper.Get("FIND_PODS_BY"); fetchType {
	case "labels":
		_, namespace := getPodAndNamespaceFromLabel(arg)
		execCmd = fmt.Sprintf("kubectl exec -it %s --namespace=%s -- %s", pod, namespace, shell)

		if container != "" {
			execCmd = fmt.Sprintf("kubectl exec -it %s --namespace=%s -c=%s -- %s", pod, namespace, container, shell)
		}
	case "namespaces":
		execCmd = fmt.Sprintf("kubectl exec -it %s --namespace=%s -- %s", pod, arg, shell)

		if container != "" {
			execCmd = fmt.Sprintf("kubectl exec -it %s -c=%s --namespace=%s -- %s", pod, container, arg, shell)
		}
	default:
		log.Fatalf("You must set FIND_PODS_BY in .kuve.yaml to namespaces or labels")
	}

	binary, lookErr := exec.LookPath("kubectl")

	if lookErr != nil {
		panic(lookErr)
	}
	env := os.Environ()

	execErr := syscall.Exec(binary, strings.Split(execCmd, " "), env)
	if execErr != nil {
		panic(execErr)
	}
}

func getPod(arg string) string {
	var podCmd string

	switch fetchType := viper.Get("FIND_PODS_BY"); fetchType {
	case "labels":
		if label != "" {
			podCmd = fmt.Sprintf("kubectl get pods -n %s --selector=%s --field-selector=status.phase=Running --no-headers | head -n1 | cut -d ' ' -f1", arg, label)
		} else {
			podCmd = fmt.Sprintf("kubectl get pods --all-namespaces -l=%s=%s --field-selector=status.phase=Running --no-headers | head -n1 | cut -d ' ' -f4", viper.Get("POD_LABEL_KEY"), arg)
		}
	case "namespaces":
		podCmd = fmt.Sprintf("kubectl get pods -n %s --field-selector=status.phase=Running --no-headers | head -n1 | cut -d ' ' -f1", arg)

		if label != "" {
			podCmd = fmt.Sprintf("kubectl get pods -n %s --selector=%s --field-selector=status.phase=Running --no-headers | head -n1 | cut -d ' ' -f1", arg, label)
		}
	default:
		log.Fatalf("You must set FIND_PODS_BY in .kuve.yaml to namespaces or labels")
	}

	pod, err := exec.Command("/bin/sh", "-c", podCmd).Output()
	if err != nil {
		fmt.Println(err)
	}

	if len(pod) <= 0 {
		log.Fatalf("Cannot find running pods with command generated: %s", podCmd)
	}

	return strings.TrimSpace(string(pod))
}
