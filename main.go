package main

import (
	"git.sr.ht/~hjertnes/tw.txt/commands/html"
	"git.sr.ht/~hjertnes/tw.txt/loadfeeds"
	"git.sr.ht/~hjertnes/tw.txt/loadfeeds/cache"
	"git.sr.ht/~hjertnes/tw.txt/loadfeeds/getfeeds"
	"os"

	"git.sr.ht/~hjertnes/tw.txt/commands/edit"
	"git.sr.ht/~hjertnes/tw.txt/commands/follow"
	"git.sr.ht/~hjertnes/tw.txt/commands/following"
	"git.sr.ht/~hjertnes/tw.txt/commands/help"
	"git.sr.ht/~hjertnes/tw.txt/commands/setup"
	"git.sr.ht/~hjertnes/tw.txt/commands/testfeeds"
	"git.sr.ht/~hjertnes/tw.txt/commands/timeline"
	"git.sr.ht/~hjertnes/tw.txt/commands/tweet"
	"git.sr.ht/~hjertnes/tw.txt/commands/unfollow"
	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/loadfeeds/headfeeds"
	"git.sr.ht/~hjertnes/tw.txt/utils"
)

const two = 2

func runProgram(args []string) {
	command, subCommand, params := utils.ParseArgs(args)

	switch command {
	case "setup":
		setup.New().Execute()
	case "html":
		conf, err := config.New()
		utils.ErrorHandler(err)

		c, err := cache.New()

		hf := headfeeds.New(conf)
		gf := getfeeds.New(conf)
		lf := loadfeeds.New(conf, c, hf, gf)

		html.New(conf, lf).Execute()
	case "timeline":
		conf, err := config.New()
		utils.ErrorHandler(err)

		c, err := cache.New()
		utils.ErrorHandler(err)

		hf := headfeeds.New(conf)
		gf := getfeeds.New(conf)
		lf := loadfeeds.New(conf, c, hf, gf)

		timeline.New(conf, lf).Execute(subCommand)
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

		c, err := cache.New()

		hf := headfeeds.New(conf)
		gf := getfeeds.New(conf)
		lf := loadfeeds.New(conf, c, hf, gf)

		testfeeds.New(lf).Execute()
	default:
		help.New().Execute()
	}
}

func main() {
	runProgram(os.Args)
}
