package loadfeeds

import (
	"fmt"
	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/mocks"
	"git.sr.ht/~hjertnes/tw.txt/models"
	"git.sr.ht/~hjertnes/tw.txt/utils"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

func TestTest(t *testing.T){
	_ = os.Setenv("TEST", "true")

	config.CreateConfigFiles()

	conf := &models.Config{
		CommonConfig: &models.CommonConfig{
			Nick:             "hjertnes",
			URL:              "https://hjertnes.social/twtxt.txt",
			DiscloseIdentity: true,
			Following: map[string]string{
				"hjertnes":    "https://hjertnes.social/twtxt.txt",
			},
			File: utils.ReplaceTilde(fmt.Sprintf("%s/twtxt.txt", config.GetConfigDir())),
		},
	}

	c := &mocks.ConfigMock{}
	g := &mocks.GetFeedsMock{}
	h := &mocks.HeadFeedsMock{}
	cache := &mocks.CacheMock{}

	c.On("Get").Return(conf)
	c.On("Save").Return(nil)
	cache.On("Get", "https://hjertnes.social/twtxt.txt").Return(&models.CachedUser{}, nil)
	cache.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return()
	h.On("Execute", mock.Anything).Return([]models.FeedHead{})
	g.On("Execute", mock.Anything).Return([]models.Feed{})
	cache.On("Save").Return(nil)

		New(c, cache, h, g).Execute()

	config.DeleteConfigFiles()

	_ = os.Setenv("TEST", "")
}

