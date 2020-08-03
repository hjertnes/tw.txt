// Package headfeeds fetches feeds
package headfeeds

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"git.sr.ht/~hjertnes/tw.txt/models"

	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/constants"
	"github.com/schollz/progressbar/v3"
)

// Command is the publicly exposed interface.
type Command interface {
	Execute(feeds map[string]string) []models.FeedHead
}

type command struct {
	config *config.Config
}

const maxfetchers = 50

func (c *command) Execute(feeds map[string]string) []models.FeedHead {
	bar := progressbar.Default(int64(len(feeds)), "Loading...")

	tweetsch := make(chan *http.Response, len(feeds))

	var wg sync.WaitGroup
	// max parallel http fetchers
	fetchers := make(chan struct{}, maxfetchers)

	for handle, url := range feeds {
		wg.Add(1)
		fetchers <- struct{}{}

		go func(handle string, url string) {
			resp := c.GetFeed(url)

			tweetsch <- resp

			<-fetchers

			_ = bar.Add(1)

			wg.Done()
		}(handle, url)
	}

	go func() {
		wg.Wait()
		close(tweetsch)
	}()

	result := make([]models.FeedHead, 0)

	for feed := range tweetsch {
		if feed == nil{
			continue
		}

		lm := time.Now()
		cl := 0

		if feed.Header["Last-Modified"] != nil{
			lm, _ = time.Parse(time.RFC1123, feed.Header["Last-Modified"][0])
		}

		if feed.Header["Content-Length"] != nil{
			cl, _ = strconv.Atoi(feed.Header["Content-Length"][0])
		}


		result = append(result, models.FeedHead{
			URL: feed.Request.URL.String(),
			LastModified: lm,
			ContentLength: cl,
		})
	}

	return result
}

// GetFeed Fetches a feed.
func (c *command) GetFeed(url string) *http.Response {
	client := http.Client{Timeout: time.Second * constants.Two}
	ctx := context.TODO()
	req, _ := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)

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
		return nil
	}

	_ = resp.Body.Close()

	return resp
}

// New creates new Command.
func New(conf *config.Config) Command {
	return &command{config: conf}
}
