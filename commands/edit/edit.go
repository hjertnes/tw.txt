// Package edit contains command for opening a twtxt file in EDITOR
package edit

import (
	"os"
	"os/exec"

	"git.sr.ht/~hjertnes/tw.txt/config"
)

// Command for editing twtxt in EDITOR.
type Command interface {
	Execute(subCommand string)
}

type command struct {
	Config *config.Config
}

func (c *command) Execute(subCommand string) {
	cmd := exec.Command(os.Getenv("EDITOR"))
	switch subCommand {
	case "internal-config":
		cmd.Args = append(cmd.Args, "~/tw.txt/config.yaml")
	case "common-config":
		cmd.Args = append(cmd.Args, c.Config.InternalConfig.ConfigFileLocation)
	default:
		cmd.Args = append(cmd.Args, c.Config.CommonConfig.File)
	}

	err := cmd.Start()
	if err != nil{
		panic(err)
	}
}

// New creates new Command.
func New(conf *config.Config) Command {
	return &command{Config: conf}
}
