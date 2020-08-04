// Package unfollow contains a command for unfollowing a user
package unfollow

import (
	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/models"
)

// Command is the exposed interface.
type Command interface {
	Execute(nick string)
}

type command struct {
	config *models.Config
}

func (c *command) Execute(nick string) {
	delete(c.config.CommonConfig.Following, nick)
	config.Save(c.config)
}

// New is the constructor.
func New(conf *models.Config) Command {
	return &command{config: conf}
}
