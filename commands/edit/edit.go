// Package edit contains command for opening a twtxt file in EDITOR
package edit

import (
	"fmt"
	"os"
	"os/exec"

	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/utils"
)

// Command for editing twtxt in EDITOR.
type Command interface {
	Execute()
}

type command struct {
	Config *config.Config
}

func (c *command) Execute() {
	_ = exec.Command(fmt.Sprintf("%s %s", os.Getenv("EDITOR"), utils.ReplaceTilde(c.Config.CommonConfig.File))).Start()
}

// New creates new Command.
func New(conf *config.Config) Command {
	return &command{Config: conf}
}
