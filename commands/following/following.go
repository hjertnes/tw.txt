// Package following shows who you are following
package following

import (
	"fmt"
	"git.sr.ht/~hjertnes/tw.txt/models"
)

// Command is the exposed interface.
type Command interface {
	Execute()
}

type command struct {
	config *models.Config
}

func (c *command) Execute() {
	for handle, url := range c.config.CommonConfig.Following {
		fmt.Printf("@%s %s\n", handle, url)
	}
}

// New is constructor.
func New(conf *models.Config) Command {
	return &command{config: conf}
}
