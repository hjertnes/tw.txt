// Package following shows who you are following
package following

import (
	"fmt"
	"git.sr.ht/~hjertnes/tw.txt/config"
)

// Command is the exposed interface.
type Command interface {
	Execute()
}

type command struct {
	config config.Service
}

func (c *command) Execute() {
	for handle, url := range c.config.Get().CommonConfig.Following {
		fmt.Printf("@%s %s\n", handle, url)
	}
}

// New is constructor.
func New(conf config.Service) Command {
	return &command{config: conf}
}
