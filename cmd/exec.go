package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
)

var (
	shell     string
	container string
	label     string
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec [namespace]",
	Short: "Execs into the first running pod and container of a namespace",
	Long:  `Execs into the first running pod and container of a namespace`,
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

func execPod(namespace string) {
	pod := getPod(namespace)

	execCmd := fmt.Sprintf("kubectl exec -it %s --namespace=%s -- %s", pod, namespace, shell)

	if container != "" {
		execCmd = fmt.Sprintf("kubectl exec -it %s -c=%s --namespace=%s -- %s", pod, container, namespace, shell)
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

func getPod(namespace string) string {
	podCmd := fmt.Sprintf("kubectl get pods -n %s --field-selector=status.phase=Running --no-headers | head -n1 | cut -d ' ' -f1", namespace)

	if label != "" {
		podCmd = fmt.Sprintf("kubectl get pods -n %s --selector=%s --field-selector=status.phase=Running --no-headers | head -n1 | cut -d ' ' -f1", namespace, label)
	}

	pod, err := exec.Command("/bin/sh", "-c", podCmd).Output()
	if err != nil {
		fmt.Println(err)
	}

	if len(pod) <= 0 {
		log.Fatalf("Cannot find running pods in namespace: %s, matching that criteria", namespace)
	}

	return strings.TrimSpace(string(pod))
}
