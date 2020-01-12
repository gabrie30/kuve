package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
)

func getPodAndNamespaceFromLabel(label string) (string, string) {
	nsPod := fmt.Sprintf("kubectl get pods --all-namespaces -l=%s=%s --no-headers | awk -v x=1 '$4 >= x' | awk '{print $1,$2}' | head -1", viper.Get("POD_LABEL_KEY"), label)
	s, _ := exec.Command("/bin/sh", "-c", nsPod).Output()
	if len(s) == 0 {
		log.Fatalf("Coud not find pod with label %s=%s", viper.Get("POD_LABEL_KEY"), label)
	}
	words := strings.Fields(string(s))
	namespace := words[0]
	pod := words[1]

	return pod, namespace
}
