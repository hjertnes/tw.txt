// Package timeline contains a command for showing the timeline
package timeline

import (
	"fmt"
	"git.sr.ht/~hjertnes/tw.txt/constants"
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/context"

	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/models"
	"git.sr.ht/~hjertnes/tw.txt/output"
	"git.sr.ht/~hjertnes/tw.txt/utils"

	"github.com/schollz/progressbar/v3"
)

// Command is the publicly exposed interface.
type Command interface {
	Execute(subCommand string)
}

type command struct {
	Config *config.Config
}

const maxfetchers = 50

func (c *command) Execute(subCommand string) {
	timeline := make([]models.Tweet, 0)

	feeds := c.Config.CommonConfig.Following
	feeds[c.Config.CommonConfig.Nick] = c.Config.CommonConfig.URL

	bar := progressbar.Default(int64(len(feeds)), "Updating feeds...")

	tweetsch := make(chan []models.Tweet, len(feeds))

	var wg sync.WaitGroup
	// max parallel http fetchers
	fetchers := make(chan struct{}, maxfetchers)

	for handle, url := range feeds {
		wg.Add(1)
		fetchers <- struct{}{}

		go func(handle string, url string) {
			defer func() {
				<-fetchers

				_ = bar.Add(1)

				wg.Done()
			}()

			feed, _ := c.GetFeed(url)
			tweets := utils.ParseFile(handle, url, feed)
			tweetsch <- tweets
		}(handle, url)
	}

	go func() {
		wg.Wait()
		close(tweetsch)
	}()

	for tweets := range tweetsch {
		timeline = append(timeline, tweets...)
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
	if tweet.Handle == c.Config.CommonConfig.Nick {
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
		for fnick, furl := range c.Config.CommonConfig.Following {
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

	if utils.NormalizeURL(url) == utils.NormalizeURL(c.Config.CommonConfig.URL) {
		return output.Red(str)
	}

	return output.Blue(str)
}

// GetFeed Fetches a feed.
func (c *command) GetFeed(url string) ([]string, error) {
	client := http.Client{}
	ctx := context.TODO()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if c.Config.CommonConfig.DiscloseIdentity {
		req.Header.Set(
			"User-Agent",
			fmt.Sprintf(
				"%s/%s (+%s; @%s)",
				constants.Name,
				constants.Version,
				c.Config.CommonConfig.URL,
				c.Config.CommonConfig.Nick,
			),
		)
	}

	resp, err := client.Do(req)
	if err != nil {
		return make([]string, 0), err
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return make([]string, 0), err
	}

	_ = resp.Body.Close()

	return strings.Split(string(content), "\n"), nil
}

// New creates new Command.
func New(conf *config.Config) Command {
	return &command{Config: conf}
}
