package fetchfeeds

import (
	"testing"

	"git.sr.ht/~hjertnes/tw.txt/config"
	"github.com/stretchr/testify/assert"
)

func TestFetchFeed(t *testing.T) {
	conf := &config.Config{
		CommonConfig: &config.CommonConfig{
			Nick:             "hjertnes",
			URL:              "https://hjertnes.social/twtxt.txt",
			DiscloseIdentity: true,
			Following: map[string]string{
				"nonExisting": "http://example.org/feed.txt",
				"hjertnes":    "https://hjertnes.social/twtxt.txt",
			},
		},
	}
	result := New(conf).Execute("Test...")
	assert.Len(t, result, 2)

	for _, r := range result {
		if r.Handle == "hjertnes" {
			assert.True(t, r.Status)
		} else {
			assert.False(t, r.Status)
		}
	}
}
