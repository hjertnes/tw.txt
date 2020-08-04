package unfollow

import (
	"fmt"
	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/mocks"
	"git.sr.ht/~hjertnes/tw.txt/models"
	"git.sr.ht/~hjertnes/tw.txt/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTest(t *testing.T) {
	_ = os.Setenv("TEST", "true")

	config.CreateConfigFiles()

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

	c.On("Get").Return(conf)
	c.On("Save").Return(nil)

	conf.CommonConfig.Following["a"] = "b"

	New(c).Execute("a")
	assert.Equal(t, "", conf.CommonConfig.Following["a"])

	config.DeleteConfigFiles()

	_ = os.Setenv("TEST", "")
}