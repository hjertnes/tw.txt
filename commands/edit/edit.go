// Package edit contains command for opening a twtxt file in EDITOR
package edit

import (
	"git.sr.ht/~hjertnes/tw.txt/config"
	"os"
	"os/exec"

	"git.sr.ht/~hjertnes/tw.txt/utils"
)

// Command for editing twtxt in EDITOR.
type Command interface {
	Execute(subCommand string)
}

type command struct {
	config config.Service
}

func (c *command) Execute(subCommand string) {
	/* #nosec */
	cmd := exec.Command(os.Getenv("EDITOR"))
	/* #sec */

	switch subCommand {
	case "internal-config":
		cmd.Args = append(cmd.Args, "~/tw.txt/config.yaml")
	case "common-config":
		cmd.Args = append(cmd.Args, c.config.Get().InternalConfig.ConfigFileLocation)
	default:
		cmd.Args = append(cmd.Args, c.config.Get().CommonConfig.File)
	}

	err := cmd.Start()
	utils.ErrorHandler(err)
}

// New creates new Command.
func New(conf config.Service) Command {
	return &command{config: conf}
}
