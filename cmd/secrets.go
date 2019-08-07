package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/gabrie30/kuve/colorlog"
	"github.com/spf13/cobra"
)

// secretsCmd represents the secrets command
var secretsCmd = &cobra.Command{
	Use:   "secrets [namespace]",
	Short: "Base64 decode and view secrets from a given namespace",
	Long:  `Base64 decode and view secrets from a given namespace`,
	Run: func(cmd *cobra.Command, args []string) {
		decodeSecrets(args[0])
	},
}

func init() {
	rootCmd.AddCommand(secretsCmd)
}

func decodeSecrets(namespace string) {
	printSecrets(namespace)

	colorlog.Info("Enter a secret from the above list do you wish to view. Please type in a value below and hit enter.")
	fmt.Println("")
	fmt.Print("Secret name to decode: ")

	var input string
	fmt.Scanln(&input)

	if input == "" {
		colorlog.FatalError("You must enter a secret name, try again.")
	}

	secretsCmd := fmt.Sprintf("kubectl get secret %s --namespace=%s -o json", input, namespace)
	rawSecrets, err := exec.Command("/bin/sh", "-c", secretsCmd).Output()

	if err != nil {
		colorlog.FatalError(err)
	}

	var secretz map[string]interface{}
	json.Unmarshal(rawSecrets, &secretz)

	data := secretz["data"].(map[string]interface{})
	fmt.Println("")
	colorlog.Info(fmt.Sprintf("**************************************************"))
	colorlog.Info(fmt.Sprintf("  namespace: %s | secret: %s \n", namespace, input))

	for k, v := range data {
		decodedValue, _ := base64.StdEncoding.DecodeString(v.(string))
		fmt.Printf("  %s: %v\n", k, string(decodedValue))
	}
	fmt.Println("")
	colorlog.Info(fmt.Sprintf("**************************************************"))
	fmt.Println("")
}

func printSecrets(namespace string) {
	secretsCmd := fmt.Sprintf("kubectl get secrets -n %s", namespace)

	secrets, err := exec.Command("/bin/sh", "-c", secretsCmd).Output()
	if err != nil {
		fmt.Println(err)
	}

	if len(secrets) <= 0 {
		colorlog.FatalError(fmt.Sprintf("Cannot find secrets in namespace: %s", namespace))
	}

	fmt.Println(string(secrets))
}
