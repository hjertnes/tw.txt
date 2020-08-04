// Package getfeeds fetches feeds
package getfeeds

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"

	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/constants"
	"git.sr.ht/~hjertnes/tw.txt/models"
	"github.com/schollz/progressbar/v3"
)

// Command is the publicly exposed interface.
type Command interface {
	Execute(feeds map[string]string) []models.Feed
}

type command struct {
	config config.Service
}

const maxfetchers = 50

func (c *command) Execute(feeds map[string]string) []models.Feed {
	bar := progressbar.Default(int64(len(feeds)), "Loading...")

	tweetsch := make(chan models.Feed, len(feeds))

	var wg sync.WaitGroup
	// max parallel http fetchers
	fetchers := make(chan struct{}, maxfetchers)

	for handle, url := range feeds {
		wg.Add(1)
		fetchers <- struct{}{}

		go func(handle string, url string) {
			status, body, headers := c.GetFeed(url)

			lm := time.Now()
			cl := 0

			if headers["Last-Modified"] != nil {
				lm, _ = time.Parse(time.RFC1123, headers["Last-Modified"][0])
			}

			if headers["Content-Length"] != nil {
				cl, _ = strconv.Atoi(headers["Content-Length"][0])
			}

			tweetsch <- models.Feed{Handle: handle, Status: status, URL: url, Body: body, ContentLength: cl, LastModified: lm}

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
func (c *command) GetFeed(url string) (bool, string, http.Header) {
	client := http.Client{Timeout: time.Second * constants.Two}
	ctx := context.TODO()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if c.config.Get().CommonConfig.DiscloseIdentity {
		req.Header.Set(
			"User-Agent",
			fmt.Sprintf(
				"%s/%s (+%s; @%s)",
				constants.Name,
				constants.Version,
				c.config.Get().CommonConfig.URL,
				c.config.Get().CommonConfig.Nick,
			),
		)
	}

	resp, err := client.Do(req)
	if err != nil {
		return false, "", nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		body = []byte("")
	}

	_ = resp.Body.Close()

	return resp.StatusCode == http.StatusOK, string(body), resp.Header
}

// New creates new Command.
func New(conf config.Service) Command {
	return &command{config: conf}
}
