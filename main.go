package main

import (
	"git.sr.ht/~hjertnes/tw.txt/commands/help"
	"os"

	"git.sr.ht/~hjertnes/tw.txt/commands/edit"
	"git.sr.ht/~hjertnes/tw.txt/commands/setup"
	"git.sr.ht/~hjertnes/tw.txt/commands/timeline"
	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/utils"
)

func main() {
	command,subCommand := utils.ParseArgs(os.Args)


	switch command {
	case "setup":
		setup.New().Execute()
	case "timeline":
		conf, err := config.New()
		utils.ErrorHandler(err)
		timeline.New(conf).Execute(subCommand)
	case "edit":
		conf, err := config.New()
		utils.ErrorHandler(err)
		edit.New(conf).Execute(subCommand)
	default:
		help.New().Execute()
	}
}
