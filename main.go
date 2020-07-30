package main

import (
	"git.sr.ht/~hjertnes/tw.txt/services/fetchfeeds"
	"os"

	"git.sr.ht/~hjertnes/tw.txt/commands/following"
	"git.sr.ht/~hjertnes/tw.txt/commands/testfeeds"
	"git.sr.ht/~hjertnes/tw.txt/commands/tweet"

	"git.sr.ht/~hjertnes/tw.txt/commands/follow"
	"git.sr.ht/~hjertnes/tw.txt/commands/unfollow"

	"git.sr.ht/~hjertnes/tw.txt/commands/help"

	"git.sr.ht/~hjertnes/tw.txt/commands/edit"
	"git.sr.ht/~hjertnes/tw.txt/commands/setup"
	"git.sr.ht/~hjertnes/tw.txt/commands/timeline"
	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/utils"
)

const two = 2

func main() {
	command, subCommand, params := utils.ParseArgs(os.Args)

	switch command {
	case "setup":
		setup.New().Execute()
	case "timeline":
		conf, err := config.New()
		utils.ErrorHandler(err)

		ff := fetchfeeds.New(conf)

		timeline.New(conf, ff).Execute(subCommand)
	case "edit":
		conf, err := config.New()
		utils.ErrorHandler(err)
		edit.New(conf).Execute(subCommand)
	case "follow":
		if len(params) < two {
			help.New().Execute()
		} else {
			conf, err := config.New()
			utils.ErrorHandler(err)
			follow.New(conf).Execute(params[0], params[1])
		}
	case "unfollow":
		if len(params) < 1 {
			help.New().Execute()
		} else {
			conf, err := config.New()
			utils.ErrorHandler(err)
			unfollow.New(conf).Execute(params[0])
		}
	case "tweet":
		if len(params) < 1 {
			help.New().Execute()
		} else {
			conf, err := config.New()
			utils.ErrorHandler(err)
			tweet.New(conf).Execute(params[0])
		}
	case "replace-mentions":
		conf, err := config.New()
		utils.ErrorHandler(err)
		tweet.New(conf).Execute("")
	case "following":
		conf, err := config.New()
		utils.ErrorHandler(err)
		following.New(conf).Execute()
	case "test-feeds":
		conf, err := config.New()
		utils.ErrorHandler(err)

		ff := fetchfeeds.New(conf)

		testfeeds.New(ff).Execute()
	default:
		help.New().Execute()
	}
}
