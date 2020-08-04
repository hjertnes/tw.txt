package timeline

import (
	"fmt"
	"testing"

	"git.sr.ht/~hjertnes/tw.txt/mocks"
	"git.sr.ht/~hjertnes/tw.txt/models"
	"git.sr.ht/~hjertnes/tw.txt/utils"

	"git.sr.ht/~hjertnes/tw.txt/config"
)

func TestTest(t *testing.T) {
	conf := &models.Config{
		CommonConfig: &models.CommonConfig{
			Nick:             "hjertnes",
			URL:              "https://hjertnes.social/twtxt.txt",
			DiscloseIdentity: true,
			Following: map[string]string{
				"hjertnes":    "https://hjertnes.social/twtxt.txt",
				"nonExisting": "http://example.org/feed.txt",
			},
			File: utils.ReplaceTilde(fmt.Sprintf("%s/twtxt.txt", config.GetConfigDir())),
		},
	}

	c := &mocks.ConfigMock{}
	lf := &mocks.LoadFeedsMock{}

	c.On("Get").Return(conf)

	lf.On("Execute").Return([]models.Feed{
		{},
	})

	New(c, lf).Execute("")
	New(c, lf).Execute("full")
}
