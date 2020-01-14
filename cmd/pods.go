/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// podsCmd represents the pods command
var podsCmd = &cobra.Command{
	Use:   "pods",
	Short: "Returns pods given settings in .kuve.conf",
	Long:  `Returns pods given settings in .kuve.conf`,
	Run: func(cmd *cobra.Command, args []string) {
		getPods(args[0])
	},
}

func getPods(arg string) {

	var podCmd string

	switch fetchType := viper.Get("FIND_PODS_BY"); fetchType {
	case "labels":
		podCmd = fmt.Sprintf("kubectl get pods --all-namespaces -l=%s=%s", viper.Get("POD_LABEL_KEY"), arg)
	case "namespaces":
		podCmd = fmt.Sprintf("kubectl get pods -n %s", arg)
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

	fmt.Println(string(pod))

}

func init() {
	rootCmd.AddCommand(podsCmd)
}
