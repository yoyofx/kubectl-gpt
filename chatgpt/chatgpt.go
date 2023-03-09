package chatgpt

import (
	"context"
	"fmt"
	"github.com/alecthomas/chroma/quick"
	"github.com/sashabaranov/go-openai"
	"os"
	"strings"
)

const (
	KUBERNETES_COMMAND        = " kubectl command:"
	KUBERNETES_RESOURCES_YAML = " k8s yaml file \n no description and note required \n"
)

var (
	client       *openai.Client
	tempYamlFile = ""
)

func InitEnv() {
	envToken, exists := os.LookupEnv("CHATGPT_TOKEN")
	if !exists {
		fmt.Println("CHATGPT_TOKEN is not set! initializing error , program will be shut down!")
		os.Exit(1)
	}
	client = openai.NewClient(envToken)
}

func Question(pref string, content string) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:            openai.GPT3Dot5Turbo,
			FrequencyPenalty: 0.0,
			PresencePenalty:  0.0,
			MaxTokens:        250,
			Temperature:      0,
			TopP:             1,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "kubectl k8s yaml command",
				},
				{
					Role: openai.ChatMessageRoleUser,
					Content: "Convert this text to a " + pref + "\n\n" +
						content,
				},
			},
		},
	)

	if err != nil {
		return
	}

	gptResponse := resp.Choices[0].Message.Content
	gptResponse, isYaml := ParseCommandAndYAML(gptResponse)
	fmt.Println("chatGPT response: \n")
	if isYaml {
		yamlFormatter := quick.Highlight(os.Stdout, gptResponse, "yaml", "terminal", "monokai")
		fmt.Println(yamlFormatter)
		fmt.Println("")
		fmt.Println("This is a YAML , you can save it by Press save command , such as save ./res.yaml")
	} else {
		yamlFormatter := quick.Highlight(os.Stdout, gptResponse, "bash", "terminal", "monokai")
		fmt.Println(yamlFormatter)
	}
}

func ParseCommandAndYAML(content string) (string, bool) {
	content = strings.Trim(content, "\n")
	if strings.HasPrefix(content, "apiVersion:") {
		//resource yaml
		tempYamlFile = content
		return content, true
	} else {
		tempYamlFile = ""
	}
	return content, false
}
