package main

import (
	"os"

	"git.sr.ht/~hjertnes/tw.txt/commands/edit"
	"git.sr.ht/~hjertnes/tw.txt/commands/follow"
	"git.sr.ht/~hjertnes/tw.txt/commands/following"
	"git.sr.ht/~hjertnes/tw.txt/commands/help"
	"git.sr.ht/~hjertnes/tw.txt/commands/html"
	"git.sr.ht/~hjertnes/tw.txt/commands/setup"
	"git.sr.ht/~hjertnes/tw.txt/commands/testfeeds"
	"git.sr.ht/~hjertnes/tw.txt/commands/timeline"
	"git.sr.ht/~hjertnes/tw.txt/commands/tweet"
	"git.sr.ht/~hjertnes/tw.txt/commands/unfollow"
	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/loadfeeds"
	"git.sr.ht/~hjertnes/tw.txt/loadfeeds/cache"
	"git.sr.ht/~hjertnes/tw.txt/loadfeeds/getfeeds"
	"git.sr.ht/~hjertnes/tw.txt/loadfeeds/headfeeds"
	"git.sr.ht/~hjertnes/tw.txt/utils"
)

const two = 2

func configErrorHandler(command string, err error){
	if  command == "timeline" ||
		command == "html" ||
		command == "follow" ||
		command == "unfollow" ||
		command == "following" ||
		command == "replace-mentions" ||
		command == "tweet" ||
		command == "test-feeds" {
		utils.ErrorHandler(err)
	}
}

func runProgram(args []string) {
	command, subCommand, params := utils.ParseArgs(args)

	conf, err := config.New()

	configErrorHandler(command, err)

	lf := buildLoadFeeds(conf)

	if len(params) < two && (command == "follow" || command == "unfollow" || command == "tweet"){
		help.New().Execute()
		return
	}

	switch command {
	case "setup":
		setup.New().Execute()
	case "html":
		html.New(conf, lf).Execute()
	case "timeline":
		timeline.New(conf, lf).Execute(subCommand)
	case "edit":
		edit.New(conf).Execute(subCommand)
	case "follow":
		follow.New(conf).Execute(params[0], params[1])
	case "unfollow":
		unfollow.New(conf).Execute(params[0])
	case "tweet":
		tweet.New(conf).Execute(params[0])
	case "replace-mentions":
		tweet.New(conf).Execute("")
	case "following":
		following.New(conf).Execute()
	case "test-feeds":
		testfeeds.New(lf).Execute()
	default:
		help.New().Execute()
	}
}

func buildLoadFeeds(conf config.Service) loadfeeds.Service{
	c, _ := cache.New()
	//utils.ErrorHandler(err)

	hf := headfeeds.New(conf)
	gf := getfeeds.New(conf)
	lf := loadfeeds.New(conf, c, hf, gf)

	return lf
}

func main() {
	runProgram(os.Args)
}
