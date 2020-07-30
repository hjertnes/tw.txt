// Package fetchfeeds fetches feeds
package fetchfeeds

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"git.sr.ht/~hjertnes/tw.txt/models"

	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/constants"
	"github.com/schollz/progressbar/v3"
)

// Command is the publicly exposed interface.
type Command interface {
	Execute(progressBarText string) []models.Feed
}

type command struct {
	config *config.Config
}

const maxfetchers = 50

func (c *command) Execute(progressBarText string) []models.Feed{
	feeds := c.config.CommonConfig.Following
	feeds[c.config.CommonConfig.Nick] = c.config.CommonConfig.URL

	bar := progressbar.Default(int64(len(feeds)), progressBarText)

	tweetsch := make(chan models.Feed, len(feeds))

	var wg sync.WaitGroup
	// max parallel http fetchers
	fetchers := make(chan struct{}, maxfetchers)

	for handle, url := range feeds {
		wg.Add(1)
		fetchers <- struct{}{}

		go func(handle string, url string) {
			status, body := c.GetFeed(url)

			tweetsch <- models.Feed{Handle: handle, Status: status, URL: url, Body: body}

			<-fetchers

			_ = bar.Add(1)

			wg.Done()
		}(handle, url)
	}

	go func() {
		wg.Wait()
		close(tweetsch)
	}()
	
	result := make([]models.Feed, 0)

	for feed := range tweetsch {
		result = append(result, feed)
	}

	return result
}

// GetFeed Fetches a feed.
func (c *command) GetFeed(url string) (bool, string) {
	client := http.Client{Timeout: time.Second * 2}
	ctx := context.TODO()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if c.config.CommonConfig.DiscloseIdentity {
		req.Header.Set(
			"User-Agent",
			fmt.Sprintf(
				"%s/%s (+%s; @%s)",
				constants.Name,
				constants.Version,
				c.config.CommonConfig.URL,
				c.config.CommonConfig.Nick,
			),
		)
	}

	resp, err := client.Do(req)
	if err != nil {
		return false, ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		body = []byte("")
	}
	_ = resp.Body.Close()

	return resp.StatusCode == http.StatusOK, string(body)
}

// New creates new Command.
func New(conf *config.Config) Command {
	return &command{config: conf}
}
