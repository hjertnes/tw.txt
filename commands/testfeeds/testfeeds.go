// Package testfeeds tests status code of a feed
package testfeeds

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"git.sr.ht/~hjertnes/tw.txt/models"

	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/constants"
	"git.sr.ht/~hjertnes/tw.txt/output"
	"github.com/schollz/progressbar/v3"
)

// Command is the publicly exposed interface.
type Command interface {
	Execute()
}

type command struct {
	Config *config.Config
}

const maxfetchers = 50

func (c *command) Execute() {
	statuses := make([]models.Status, 0)

	feeds := c.Config.CommonConfig.Following
	feeds[c.Config.CommonConfig.Nick] = c.Config.CommonConfig.URL

	bar := progressbar.Default(int64(len(feeds)), "Testing feeds...")

	tweetsch := make(chan models.Status, len(feeds))

	var wg sync.WaitGroup
	// max parallel http fetchers
	fetchers := make(chan struct{}, maxfetchers)

	for handle, url := range feeds {
		wg.Add(1)
		fetchers <- struct{}{}

		go func(handle string, url string) {
			status := c.GetFeed(url)

			tweetsch <- models.Status{Handle: handle, Status: status, URL: url}

			<-fetchers

			_ = bar.Add(1)

			wg.Done()
		}(handle, url)
	}

	go func() {
		wg.Wait()
		close(tweetsch)
	}()

	for status := range tweetsch {
		statuses = append(statuses, status)
	}

	for _, status := range statuses {
		if status.Status {
			fmt.Println(output.Green(fmt.Sprintf("@<%s %s>: OK", status.Handle, status.URL)))
		} else {
			fmt.Println(output.Red(fmt.Sprintf("@<%s %s>: Error", status.Handle, status.URL)))
		}
	}
}

// GetFeed Fetches a feed.
func (c *command) GetFeed(url string) bool {
	client := http.Client{Timeout: time.Second * 2}
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
		return false
	}

	_ = resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

// New creates new Command.
func New(conf *config.Config) Command {
	return &command{Config: conf}
}
