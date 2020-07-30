// Package timeline contains a command for showing the timeline
package timeline

import (
	"fmt"
	"git.sr.ht/~hjertnes/tw.txt/services/fetchfeeds"
	"regexp"
	"sort"
	"strings"
	"time"

	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/models"
	"git.sr.ht/~hjertnes/tw.txt/output"
	"git.sr.ht/~hjertnes/tw.txt/utils"
)

// Command is the publicly exposed interface.
type Command interface {
	Execute(subCommand string)
}

type command struct {
	config *config.Config
	fetchFeeds fetchfeeds.Command
}



func (c *command) Execute(subCommand string) {
	timeline := make([]models.Tweet, 0)
	feeds := c.fetchFeeds.Execute("Fetching feeds...")

	for _, feed := range feeds{
		lines := strings.Split(feed.Body, "\n")
		timeline = append(timeline, utils.ParseFile(feed.Handle, feed.URL, lines)...)
	}

	sort.SliceStable(timeline, func(i int, j int) bool {
		return timeline[j].Timestamp.After(timeline[i].Timestamp)
	})

	for i, tweet := range timeline {
		if i > len(timeline)-1000 || subCommand == "full" {
			c.PrintTweet(tweet, time.Now())
		}
	}
}

func (c *command) PrintTweet(tweet models.Tweet, now time.Time) {
	text := c.ShortenMentions(tweet.Message)

	nick := output.Green(tweet.Handle)
	if tweet.Handle == c.config.CommonConfig.Nick {
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
		for fnick, furl := range c.config.CommonConfig.Following {
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

	if utils.NormalizeURL(url) == utils.NormalizeURL(c.config.CommonConfig.URL) {
		return output.Red(str)
	}

	return output.Blue(str)
}

// New creates new Command.
func New(conf *config.Config, ff fetchfeeds.Command) Command {
	return &command{config: conf, fetchFeeds: ff}
}
