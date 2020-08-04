// Package html contains a command for showing the timeline
package html

import (
	"fmt"
	"html/template"
	"os"
	"sort"
	"strings"
	"time"

	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/constants"

	"git.sr.ht/~hjertnes/patterns"
	"git.sr.ht/~hjertnes/tw.txt/loadfeeds"
	"git.sr.ht/~hjertnes/tw.txt/models"
	"git.sr.ht/~hjertnes/tw.txt/utils"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

// Command is the publicly exposed interface.
type Command interface {
	Execute()
}

type command struct {
	config    config.Service
	loadFeeds loadfeeds.Service
}

func (c *command) Execute() {
	timeline := make([]models.Tweet, 0)
	feeds := c.loadFeeds.Execute()

	for _, feed := range feeds {
		lines := strings.Split(feed.Body, "\n")
		timeline = append(timeline, utils.ParseFile(feed.Handle, feed.URL, lines)...)
	}

	newTimeline := c.replaceStuff(timeline)

	sort.SliceStable(newTimeline, func(i int, j int) bool {
		return newTimeline[j].Timestamp.Before(newTimeline[i].Timestamp)
	})

	t, err := template.ParseFiles(utils.ReplaceTilde(c.config.Get().InternalConfig.TemplateFileLocation))
	utils.ErrorHandler(err)

	if !utils.Exist("timeline.html") {
		f, err := os.Create("timeline.html")
		utils.ErrorHandler(err)
		err = f.Close()
		utils.ErrorHandler(err)
	}

	f, err := os.OpenFile("timeline.html", os.O_RDWR, 0600)

	utils.ErrorHandler(err)

	err = t.Execute(f, models.HTMLModel{
		Timestamp: time.Now().UTC(),
		Timeline:  newTimeline,
	})
	utils.ErrorHandler(err)
}

func rewriteOrgModeLinks(input string) string {
	for {
		parts, err := patterns.FindAndSplit(input, "[[", "]]", "][")
		if err != nil {
			break
		}

		if len(parts) == 1 {
			input = strings.ReplaceAll(
				input,
				fmt.Sprintf("[[%s]]", parts[0]),
				fmt.Sprintf(`<a href="%s">%s</a>`, parts[0], parts[0]))
		} else {
			input = strings.ReplaceAll(
				input,
				fmt.Sprintf("[[%s][%s]", parts[0], parts[1]),
				fmt.Sprintf(`<a href="%s">%s</a>`, parts[0], parts[1]))
		}
	}

	return input
}

func rewriteMentions(input string) string {
	for {
		match, err := patterns.FindAndSplit(input, "@<", ">", " ")
		if err != nil {
			break
		}

		if len(match) < constants.Two {
			break
		}

		input = strings.ReplaceAll(
			input,
			fmt.Sprintf("@<%s %s>", match[0], match[1]),
			fmt.Sprintf(`<a href="%s">@%s</a>`, match[1], match[0]))
	}

	return input
}

func (c *command) replaceStuff(timeline []models.Tweet) []models.HTMLTweet {
	result := make([]models.HTMLTweet, 0)

	for _, tweet := range timeline {
		classes := make(map[string]string)

		if c.config.Get().CommonConfig.Nick == tweet.Handle {
			classes["by-myself"] = "by-myself"
		}

		if strings.Contains(
			tweet.Message,
			fmt.Sprintf("@<%s %s>", c.config.Get().CommonConfig.Nick, c.config.Get().CommonConfig.URL),
		) {
			classes["mentioned"] = "mentioned"
		}

		tweet.Message = strings.ReplaceAll(tweet.Message, "<script>", "script")
		tweet.Message = strings.ReplaceAll(tweet.Message, "</script>", "script")
		tweet.Message = rewriteMentions(tweet.Message)
		tweet.Message = rewriteOrgModeLinks(tweet.Message)

		html := markdown.ToHTML([]byte(tweet.Message), parser.New(), nil)

		result = append(result, models.HTMLTweet{
			Timestamp: tweet.Timestamp,
			Handle:    tweet.Handle,
			URL:       tweet.URL,
			/* #nosec */
			Message: template.HTML(html),
			/* #sec */
			Classes: mapToString(classes),
		})
	}

	return result
}

func mapToString(input map[string]string) string {
	res := ""

	for key := range input {
		if !strings.Contains(res, key) {
			if res == "" {
				res = key
			} else {
				res = fmt.Sprintf("%s, %s", res, key)
			}
		}
	}

	return res
}

// New creates new Command.
func New(conf config.Service, lf loadfeeds.Service) Command {
	return &command{config: conf, loadFeeds: lf}
}
