// Package tweet contains command to post a new tweet
package tweet

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/utils"
)

// Command is the publicly exposed interface.
type Command interface {
	Execute(message string)
}

type command struct {
	Config *config.Config
}

func removeEmptyLines(items []string) []string {
	result := make([]string, 0)

	for _, line := range items {
		if line != "" {
			result = append(result, line)
		}
	}

	return result
}

func (c *command) replaceAtMentions(items []string) []string {
	result := make([]string, 0)

	for _, line := range items {
		for handle, url := range c.Config.CommonConfig.Following {
			line = strings.ReplaceAll(line, fmt.Sprintf("@%s", handle), fmt.Sprintf("@<%s %s>", handle, url))
		}

		result = append(result, line)
	}

	return result
}

func (c *command) Execute(message string) {
	date := time.Now().Format(time.RFC3339)

	content, err := ioutil.ReadFile(utils.ReplaceTilde(c.Config.CommonConfig.File))
	utils.ErrorHandler(err)

	lines := strings.Split(string(content), "\n")
	if message != ""{
		lines = append(lines, fmt.Sprintf("%s\t%s", date, message))
		lines = append(lines, "")
	}


	text := strings.Join(c.replaceAtMentions(removeEmptyLines(lines)), "\n")

	err = ioutil.WriteFile(utils.ReplaceTilde(c.Config.CommonConfig.File), []byte(text), 0)
	utils.ErrorHandler(err)
}

// New creates new Command.
func New(conf *config.Config) Command {
	return &command{Config: conf}
}
