// Package edit contains command for opening a twtxt file in EDITOR
package edit

import (
	"os"
	"os/exec"

	"git.sr.ht/~hjertnes/tw.txt/utils"

	"git.sr.ht/~hjertnes/tw.txt/config"
)

// Command for editing twtxt in EDITOR.
type Command interface {
	Execute(subCommand string)
}

type command struct {
	config *config.Config
}

func (c *command) Execute(subCommand string) {
	/* #nosec */
	cmd := exec.Command(os.Getenv("EDITOR"))
	/* #sec */

	switch subCommand {
	case "internal-config":
		cmd.Args = append(cmd.Args, "~/tw.txt/config.yaml")
	case "common-config":
		cmd.Args = append(cmd.Args, c.config.InternalConfig.ConfigFileLocation)
	default:
		cmd.Args = append(cmd.Args, c.config.CommonConfig.File)
	}

	err := cmd.Start()
	utils.ErrorHandler(err)
}

// New creates new Command.
func New(conf *config.Config) Command {
	return &command{config: conf}
}
