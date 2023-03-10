package debug

import (
	"log"
	"os"
)

const (
	envEnableLog = "KUBECTL_GPT_ENABLE_LOG"
	logFileName  = "kubectl-gpt.log"
)

var (
	logger *log.Logger
)

func init() {
	enableLog := os.Getenv(envEnableLog)
	if enableLog == "true" || enableLog == "1" {
		logger = log.New(os.Stdout, "", log.Lshortfile)
	}
}

// Log to output message
func Log(msg string) {
	if logger != nil {
		logger.Println(msg)
	}
}
