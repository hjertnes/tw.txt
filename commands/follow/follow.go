// Package follow contains a command for following a user
package follow

import (
	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/models"
)

// Command is the exposed interface.
type Command interface {
	Execute(nick string, url string)
}

type command struct {
	config *models.Config
}

func (c *command) Execute(nick string, url string) {
	c.config.CommonConfig.Following[nick] = url
	config.Save(c.config)
}

// New is constructor.
func New(conf *models.Config) Command {
	return &command{config: conf}
}
