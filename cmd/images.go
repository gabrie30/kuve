package cmd

import (
	"fmt"
	"os/exec"

	"github.com/gabrie30/kuve/colorlog"
	"github.com/spf13/cobra"
)

// imagesCmd represents the images command
var imagesCmd = &cobra.Command{
	Use:   "images  [namespace]",
	Short: "Returns a list of images deployed into namespace",
	Long:  `Returns a list of images deployed into namespace`,
	Run: func(cmd *cobra.Command, args []string) {
		namespace := args[0]
		printImages(namespace)
	},
}

func init() {
	rootCmd.AddCommand(imagesCmd)
}

func printImages(namespace string) {
	imgStr := fmt.Sprintf("kubectl get pods -n %s -o jsonpath='{..image}' | tr -s '[[:space:]]' '\n' | sort | uniq", namespace)

	images, err := exec.Command("/bin/sh", "-c", imgStr).Output()
	if err != nil {
		fmt.Println(err)
	}

	if len(images) <= 0 {
		colorlog.FatalError(fmt.Sprintf("Cannot find images in namespace: %s", namespace))
	}
	fmt.Println("")
	colorlog.Info(fmt.Sprintf("*************** Images Deployed in Namespace: %s ***************", namespace))
	fmt.Println("")
	fmt.Println("")
	fmt.Println(string(images))
}
