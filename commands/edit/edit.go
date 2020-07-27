// Package edit contains command for opening a twtxt file in EDITOR
package edit

import (
	"os"
	"os/exec"

	"git.sr.ht/~hjertnes/tw.txt/config"
)

// Command for editing twtxt in EDITOR.
type Command interface {
	Execute()
}

type command struct {
	Config *config.Config
}

func (c *command) Execute() {
	cmd := exec.Command(os.Getenv("EDITOR"))
	cmd.Args = append(cmd.Args, c.Config.CommonConfig.File)
	err := cmd.Start()
	if err != nil{
		panic(err)
	}
}

// New creates new Command.
func New(conf *config.Config) Command {
	return &command{Config: conf}
}
