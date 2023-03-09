package kube

import (
	"bytes"
	"fmt"
	"kubectl-gpt/chatgpt"
	"kubectl-gpt/internal/debug"
	"os"
	"os/exec"
	"strings"
)

func Executor(s string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	} else if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}
	if strings.HasPrefix(s, "gen") { // ask chatGPT , generate k8s resources and commands. .
		args := strings.Split(s, " ")
		if len(args) > 2 {
			cmdPref := ""
			switch args[1] {
			case "yaml":
				cmdPref = chatgpt.KUBERNETES_RESOURCES_YAML
			case "command":
				cmdPref = chatgpt.KUBERNETES_COMMAND
			}

			fmt.Println("Requesting chatGPT please wait ................")
			cmd := strings.Join(args[2:len(args)], " ")
			chatgpt.Question(cmdPref, cmd)
		} else {
			fmt.Println("Error, please enter the question!")
		}
	} else {
		cmdStr := "kubectl " + s
		debug.Log(cmdStr)
		cmd := exec.Command("/bin/sh", "-c", cmdStr)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Got error: %s\n", err.Error())
		}
	}
	return
}

func ExecuteAndGetResult(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		debug.Log("you need to pass the something arguments")
		return ""
	}

	out := &bytes.Buffer{}
	cmd := exec.Command("/bin/sh", "-c", s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		debug.Log(err.Error())
		return ""
	}
	r := string(out.Bytes())
	return r
}
