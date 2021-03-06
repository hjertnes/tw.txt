// Package tweet contains command to post a new tweet
package tweet

import (
	"fmt"
	"io/ioutil"
	"regexp"
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
	config config.Service
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

	c.config.Get().CommonConfig.Following[c.config.Get().CommonConfig.Nick] = c.config.Get().CommonConfig.URL

	for _, line := range items {
		re1 := regexp.MustCompile(`\s@(\w*)\s`)
		re2 := regexp.MustCompile(`\s@(\w*)$`)
		re3 := regexp.MustCompile(`@(\w*)\s`)

		matches := re1.FindAllStringSubmatch(line, -1)
		matches = append(matches, re2.FindAllStringSubmatch(line, -1)...)
		matches = append(matches, re3.FindAllStringSubmatch(line, -1)...)

		for _, match := range matches {
			line = strings.ReplaceAll(
				line,
				fmt.Sprintf("@%s", match[1]),
				fmt.Sprintf("@<%s %s>", match[1], c.config.Get().CommonConfig.Following[match[1]]))
		}

		result = append(result, line)
	}

	return result
}

func (c *command) Execute(message string) {
	date := time.Now().Format(time.RFC3339)

	content, err := ioutil.ReadFile(utils.ReplaceTilde(c.config.Get().CommonConfig.File))
	utils.ErrorHandler(err)

	lines := strings.Split(string(content), "\n")
	if message != "" {
		lines = append(lines, fmt.Sprintf("%s\t%s", date, message))
		lines = append(lines, "")
	}

	text := strings.Join(c.replaceAtMentions(removeEmptyLines(lines)), "\n")

	err = ioutil.WriteFile(utils.ReplaceTilde(c.config.Get().CommonConfig.File), []byte(text), 0)
	utils.ErrorHandler(err)
}

// New creates new Command.
func New(conf config.Service) Command {
	return &command{config: conf}
}
