package kube

import (
	"fmt"
	"github.com/yoyofx/kubectl-gpt/internal/chatgpt"
	"github.com/yoyofx/kubectl-gpt/internal/debug"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Executor(s string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	} else if s == "quit" || s == "exit" {
		os.Exit(0)
		return
	}
	if strings.HasPrefix(s, "gen") { // ask chatGPT , generate k8s resources and commands. .
		args := strings.Split(s, " ")
		cmd := strings.Join(args[2:len(args)], " ")
		if len(args) > 2 {
			cmdPref := ""
			additional := ""
			switch args[1] {
			case "yaml":
				cmdPref = chatgpt.KUBERNETES_RESOURCES_YAML
			case "command":
				cmdPref = chatgpt.KUBERNETES_COMMAND
				additional = "\n\n" + "output only kubectl command"
				cmd = cmd + additional
			}

			fmt.Println("Requesting chatGPT please wait ................")
			chatgpt.Question(cmdPref, cmd)
		} else {
			fmt.Println("Error, please enter the question!")
		}
	} else if strings.HasPrefix(s, "switch") {
		args := strings.Split(s, " ")
		if len(args) > 1 {
			err := os.Setenv("KUBECONFIG", args[1])
			if err == nil {
				fmt.Println("ENV: KUBECONFIG=", args[1])
			}
		}
	} else if strings.HasPrefix(s, "save") {
		args := strings.Split(s, " ")
		if len(args) > 1 {
			if chatgpt.TempYamlFile != "" {
				//save yaml to file

			} else {
				fmt.Println("Not Yaml to save , please re-quest chatGPT.")
			}
		}
	} else {
		cmdStr := "kubectl " + s
		debug.Log(cmdStr)
		var shell, flag string
		if runtime.GOOS == "windows" {
			shell = "cmd"
			flag = "/c"
		} else {
			shell = "/bin/sh"
			flag = "-c"
		}

		cmd := exec.Command(shell, flag, cmdStr)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Got error: %s\n", err.Error())
		}
	}

	return
}
