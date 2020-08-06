package html

import (
	"fmt"
	"testing"
	"time"

	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/mocks"
	"git.sr.ht/~hjertnes/tw.txt/models"
	"git.sr.ht/~hjertnes/tw.txt/utils"
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
		InternalConfig: &models.InternalConfig{
			TemplateFileLocation: "~/Code/tw.txt/template.html",
		},
	}

	c := &mocks.ConfigMock{}

	c.On("Get").Return(conf)

	lf := &mocks.LoadFeedsMock{}

	lf.On("Execute").Return([]models.Feed{
		{
			URL: "https://hjertnes.social/twtxt.txt",
			Handle: "hjertnes",
			Body: fmt.Sprintf("%s\t@<hjertnes https://hjertnes.social> test", time.Now().Format(time.RFC3339)),
		},
		{
			URL: "https://hjertnes.social/twtxt.txt",
			Handle: "hjertnes",
			Body: fmt.Sprintf("%s\t@<hjertnes https://hjertnes.social> test", time.Now().Format(time.RFC3339)),
		},
	})

	New(c, lf).Execute()
}
