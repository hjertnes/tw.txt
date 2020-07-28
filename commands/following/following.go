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
	Config *config.Config
}

func (c *command) Execute() {
	for handle, url := range c.Config.CommonConfig.Following {
		fmt.Println(fmt.Sprintf("@%s %s", handle, url))
	}

}

// New is constructor.
func New(conf *config.Config) Command {
	return &command{Config: conf}
}
