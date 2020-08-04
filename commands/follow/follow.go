// Package follow contains a command for following a user
package follow

import (
	"git.sr.ht/~hjertnes/tw.txt/config"
)

// Command is the exposed interface.
type Command interface {
	Execute(nick string, url string)
}

type command struct {
	config config.Service
}

func (c *command) Execute(nick string, url string) {
	c.config.Get().CommonConfig.Following[nick] = url
	c.config.Save()
}

// New is constructor.
func New(conf config.Service) Command {
	return &command{config: conf}
}
