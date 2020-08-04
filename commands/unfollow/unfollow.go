// Package unfollow contains a command for unfollowing a user
package unfollow

import (
	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/utils"
)

// Command is the exposed interface.
type Command interface {
	Execute(nick string)
}

type command struct {
	config config.Service
}

func (c *command) Execute(nick string) {
	delete(c.config.Get().CommonConfig.Following, nick)
	err := c.config.Save()
	if err != nil{
		utils.ErrorHandler(err)
	}
}

// New is the constructor.
func New(conf config.Service) Command {
	return &command{config: conf}
}
