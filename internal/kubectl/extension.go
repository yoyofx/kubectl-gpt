package kubectl

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"os"
	"path/filepath"
)

var (
	KUBE_CONFIG_MAPS []string
)

func InitExtension() {
	kubeConfigPath, exists := os.LookupEnv("KUBE_CHATGPT_CONFIG")

	if exists {
		fmt.Println("Loaded kubeconfig path:", kubeConfigPath)
		err := filepath.Walk(kubeConfigPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				KUBE_CONFIG_MAPS = append(KUBE_CONFIG_MAPS, path)
				fmt.Println(path) // file path
			}
			return nil
		})
		if err != nil {
			return
		}
		fmt.Println("You can press 'switch' command , to switch kubeconfig. \n")
	} else {
		fmt.Println("Loaded kubeconfig path: /$home/.kube/config \n")
	}
}

func GetConfigSuggestions(args []string) []prompt.Suggest {
	var subcommands []prompt.Suggest
	for _, configPath := range KUBE_CONFIG_MAPS {
		subcommands = append(subcommands, prompt.Suggest{Text: configPath, Description: filepath.Base(configPath)})
	}

	if len(args) == 2 {
		return prompt.FilterHasPrefix(subcommands, args[1], true)
	}
	return nil
}
