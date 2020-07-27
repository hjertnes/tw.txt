package main

import (
	"fmt"
	"os"

	"git.sr.ht/~hjertnes/tw.txt/commands/edit"
	"git.sr.ht/~hjertnes/tw.txt/commands/setup"
	"git.sr.ht/~hjertnes/tw.txt/commands/timeline"
	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/utils"
)

func main() {
	command, _ := utils.ParseArgs(os.Args)

	switch command {
	case "setup":
		setup.New().Execute()
	case "timeline":
		conf, err := config.New()
		utils.ErrorHandler(err)
		timeline.New(conf).Execute()
	case "edit":
		conf, err := config.New()
		utils.ErrorHandler(err)
		edit.New(conf).Execute()
	default:
		fmt.Println("tw.txt is another twtxt client -- https://twtxt.readthedocs.org/en/stable")
		fmt.Println("Usage:")
		fmt.Println("\t tw.txt command")
		fmt.Println()
		fmt.Println("Commands:")
		fmt.Println("\ttimeline - prints last 1000 items in your timeline")
		fmt.Println("\tedit - opens your twtxt in $EDITOR")
		fmt.Println("\tsetup - creates config file and opens it in $EDITOR")
	}
}
