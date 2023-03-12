package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"github.com/yoyofx/kubectl-gpt/internal/chatgpt"
	"github.com/yoyofx/kubectl-gpt/internal/kubectl"
	"github.com/yoyofx/kubectl-gpt/internal/kubectl/suggestions"
	"os"
)

var (
	version  string = "v1.0"
	revision string = "v0.1"
)

func main() {
	ShowLogo()
	// init chatGPT.
	chatgpt.InitEnv() // init kubectl extension
	kubectl.InitExtension()
	c, err := kube.NewCompleter()
	if err != nil {
		fmt.Println("init error", err)
		os.Exit(1)
	}
	fmt.Printf("welcome, kubectl-gpt %s (rev-%s)\n", version, revision)
	fmt.Println("Please use `exit` or `Ctrl-D` to exit this program.")
	p := prompt.New(
		kube.Executor, // exec command endpoint
		c.Complete,
		prompt.OptionTitle("kubectl-gpt: interactive kubernetes cmd & chatGPT client"),
		prompt.OptionPrefix("gpt >> kubectl >> "),
		prompt.OptionInputTextColor(prompt.Yellow),
		prompt.OptionSuggestionBGColor(prompt.Blue),
		prompt.OptionDescriptionBGColor(prompt.White),
		prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
	)
	p.Run()
}

func ShowLogo() {
	logo := `
                                       ##         .
                                 ## ## ##        ==
                              ## ## ## ## ##    ===
                           /""""""""""""""""\___/ ===
                      ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~
                           \______ o          _,/
                            \      \       _,'
                              '--.._\..--''        
						   
KubeLilin An Cloud-Native application platform for Kubernetes.
Document: https://doc.kubelilin.com/
kube-gpt: interactive kubernetes cmd & chatGPT client
	`
	fmt.Println(logo)
}

var dd = "Here's an example of a Kubernetes deployment YAML file for etcd:\n\n```\napiVersion: apps/v1\nkind: Deployment\nmetadata:\n  name: etcd-deployment\n  labels:\n    app: etcd\nspec:\n  replicas: 1 # number of replicas, you can increase this to scale horizontally\n  selector:\n    matchLabels:\n      app: etcd\n  template:\n    metadata:\n      labels:\n        app: etcd\n    spec:\n      containers:\n      - name: etcd-container\n        image: quay.io/coreos/etcd:v3.4.9 # image location, change this to the version you need \n        ports:\n        - containerPort: 2379 # port for etcd client communication \n          name: client-port  \n        - containerPort: 2380 # port for etcd server-to-server communication   \n          name: server-port                     \n```\n\nYou can save the above content in a file with a `.yaml` or `.yml` extension and deploy it using the `kubectl apply -f <filename>` command. For example, if you save the above content in a file called `etcd-deployment.yaml`, you can deploy it using:\n\n```\nkubectl apply -f etcd-deployment.yaml\n```\n"
