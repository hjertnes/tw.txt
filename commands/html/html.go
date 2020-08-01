// Package timeline contains a command for showing the timeline
package html

import (
	"fmt"
	"git.sr.ht/~hjertnes/tw.txt/services/fetchfeeds"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"html/template"
	"os"
	"regexp"
	"sort"
	"strings"

	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/models"
	"git.sr.ht/~hjertnes/tw.txt/utils"
)

// Command is the publicly exposed interface.
type Command interface {
	Execute()
}

type command struct {
	config     *config.Config
	fetchFeeds fetchfeeds.Command
}

func (c *command) Execute() {
	timeline := make([]models.Tweet, 0)
	feeds := c.fetchFeeds.Execute("Fetching feeds...")

	for _, feed := range feeds {
		lines := strings.Split(feed.Body, "\n")
		timeline = append(timeline, utils.ParseFile(feed.Handle, feed.URL, lines)...)
	}

	newTimeline := replaceStuff(timeline)

	sort.SliceStable(newTimeline, func(i int, j int) bool {
		return newTimeline[j].Timestamp.Before(newTimeline[i].Timestamp)
	})

	t, err := template.ParseFiles(utils.ReplaceTilde(c.config.InternalConfig.TemplateFileLocation))
	utils.ErrorHandler(err)

	if !utils.Exist("timeline.html"){
		f, err := os.Create("timeline.html")
		utils.ErrorHandler(err)
		err = f.Close()
		utils.ErrorHandler(err)
	}

	f, err := os.OpenFile("timeline.html", os.O_RDWR, 0600)

	utils.ErrorHandler(err)

	err = t.Execute(f, newTimeline)
	utils.ErrorHandler(err)
}

var r1 = regexp.MustCompile(`@<(\w*) (.*)>`)

func rewriteOrgModeLinks(input string) string{
	for {
		if !strings.Contains(input, "[["){
			break
		}

		parts := strings.Split(strings.Split(strings.Split(input, "[[")[1], "]]")[0], "][")

		if len(parts) == 1{
			input = strings.ReplaceAll(input, fmt.Sprintf("[[%s]]", parts[0]), fmt.Sprintf(`<a href="%s">%s</a>`, parts[0], parts[0]))
		} else {
			input = strings.ReplaceAll(input, fmt.Sprintf("[[%s][%s]", parts[0], parts[1]), fmt.Sprintf(`<a href="%s">%s</a>`, parts[0], parts[1]))
		}
	}

	return input
}



func replaceStuff(timeline []models.Tweet) []models.HTMLTweet{
	result := make([]models.HTMLTweet, 0)

	for _, tweet := range timeline {
		tweet.Message = r1.ReplaceAllString(tweet.Message, `<a href="$2">@$1</a>`)
		tweet.Message = strings.ReplaceAll(tweet.Message, "<script>", "script")
		tweet.Message = strings.ReplaceAll(tweet.Message, "</script>", "script")


		tweet.Message = rewriteOrgModeLinks(tweet.Message)

		html := markdown.ToHTML([]byte(tweet.Message), parser.New(), nil)

		result = append(result, models.HTMLTweet{
			Timestamp: tweet.Timestamp,
			Handle: tweet.Handle,
			URL: tweet.URL,
			Message: template.HTML(html),
		})
	}

	return result
}

// New creates new Command.
func New(conf *config.Config, ff fetchfeeds.Command) Command {
	return &command{config: conf, fetchFeeds: ff}
}
