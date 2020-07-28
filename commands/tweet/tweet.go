package tweet

import (
	"fmt"
	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/utils"
	"io/ioutil"
	"strings"
	"time"
)

// Command is the publicly exposed interface.
type Command interface {
	Execute(message string)
}

type command struct {
	Config *config.Config
}

func removeEmptyLines(items []string) []string{
	result := make([]string, 0)

	for _, line := range items{
		if line != ""{
			result = append(result, line)
		}
	}

	return result
}

func (c *command) Execute(message string){
	date := time.Now().Format(time.RFC3339)

	content, err := ioutil.ReadFile(utils.ReplaceTilde(c.Config.CommonConfig.File))
	utils.ErrorHandler(err)

	lines := strings.Split(string(content), "\n")
	lines = append(lines, fmt.Sprintf("%s\t%s", date, message))
	lines = append(lines, "")

	text := strings.Join(removeEmptyLines(lines), "\n")

	err = ioutil.WriteFile(utils.ReplaceTilde(c.Config.CommonConfig.File), []byte(text), 0)
	utils.ErrorHandler(err)
}

// New creates new Command.
func New(conf *config.Config) Command {
	return &command{Config: conf}
}
