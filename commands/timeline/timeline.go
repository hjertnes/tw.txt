// Package timeline contains a command for showing the timeline
package timeline

import (
	"fmt"
	"git.sr.ht/~hjertnes/tw.txt/config"
	"regexp"
	"sort"
	"strings"
	"time"

	"git.sr.ht/~hjertnes/tw.txt/loadfeeds"

	"git.sr.ht/~hjertnes/tw.txt/models"
	"git.sr.ht/~hjertnes/tw.txt/output"
	"git.sr.ht/~hjertnes/tw.txt/utils"
)

// Command is the publicly exposed interface.
type Command interface {
	Execute(subCommand string)
}

type command struct {
	config    config.Service
	loadFeeds loadfeeds.Service
}

func (c *command) Execute(subCommand string) {
	timeline := make([]models.Tweet, 0)

	feeds := c.loadFeeds.Execute()

	for _, feed := range feeds {
		lines := strings.Split(feed.Body, "\n")
		timeline = append(timeline, utils.ParseFile(feed.Handle, feed.URL, lines)...)
	}

	sort.SliceStable(timeline, func(i int, j int) bool {
		return timeline[j].Timestamp.After(timeline[i].Timestamp)
	})

	for i, tweet := range timeline {
		if i > len(timeline)-250 || subCommand == "full" {
			c.PrintTweet(tweet, time.Now())
		}
	}
}

func (c *command) PrintTweet(tweet models.Tweet, now time.Time) {
	text := c.ShortenMentions(tweet.Message)

	nick := output.Green(tweet.Handle)
	if tweet.Handle == c.config.Get().CommonConfig.Nick {
		nick = output.BoldGreen(tweet.Handle)
	}

	fmt.Printf("> %s (%s)\n%s\n",
		nick,
		utils.PrettyDuration(now.Sub(tweet.Timestamp)),
		text)
}

func (c *command) ShortenMentions(text string) string {
	re := regexp.MustCompile(`@<([^ ]+) *([^>]+)>`)

	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		nick, url := parts[1], parts[2]
		for fnick, furl := range c.config.Get().CommonConfig.Following {
			if furl == url {
				return c.FormatMention(nick, url, fnick)
			}
		}

		return match
	})
}

func (c *command) FormatMention(nick, url, followednick string) string {
	str := "@" + nick
	if followednick != nick {
		str += fmt.Sprintf("(%s)", followednick)
	}

	if utils.NormalizeURL(url) == utils.NormalizeURL(c.config.Get().CommonConfig.URL) {
		return output.Red(str)
	}

	return output.Blue(str)
}

// New creates new Command.
func New(conf config.Service, lf loadfeeds.Service) Command {
	return &command{config: conf, loadFeeds: lf}
}
