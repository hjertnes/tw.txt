package loadfeeds

import (
	"fmt"
	"git.sr.ht/~hjertnes/tw.txt/constants"
	"os"
	"testing"
	"time"

	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/mocks"
	"git.sr.ht/~hjertnes/tw.txt/models"
	"git.sr.ht/~hjertnes/tw.txt/utils"
	"github.com/stretchr/testify/mock"
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
				"not-in-cache": "https://not-in-cache",
				"in-cache": "https://in-cache",
				"in-cache-re-chech": "https://in-cache-re-check",
				"in-cache-re-chech1": "https://in-cache-re-check1",
				"in-cache-invalidated": "https://in-cache-invalidated",
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
	dt := time.Now()

	cache.On("Get", "https://hjertnes.social/twtxt.txt").Return(&models.CachedUser{
		URL: "https://hjertnes.social/twtxt.txt",
	}, constants.ErrNotInCache)
	cache.On("Get", "https://not-in-cache").Return(&models.CachedUser{
		URL: "https://not-in-cache",
	}, constants.ErrNotInCache)
	cache.On("Get", "https://in-cache-invalidated").Return(&models.CachedUser{
		URL: "https://in-cache-invalidated",
	}, constants.ErrExpired)
	cache.On("Get", "https://in-cache-re-check").Return(&models.CachedUser{
		URL: "https://in-cache-re-check",
		ContentLength: 1,
		LastUpdated: dt,
	}, constants.ErrFetchHead)
	cache.On("Get", "https://in-cache-re-check1").Return(&models.CachedUser{
		URL: "https://in-cache-re-check1",
		ContentLength: 2,
		LastUpdated: dt,
		Handle: "Yolo",
	}, constants.ErrFetchHead)
	cache.On("Get", "https://in-cache").Return(&models.CachedUser{
		URL: "https://in-cache",
	}, nil)

	cache.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return()
	h.On("Execute", map[string]string{"in-cache-re-chech": "https://in-cache-re-check", "in-cache-re-chech1": "https://in-cache-re-check1"}).Return([]models.FeedHead{models.FeedHead{
		URL: "https://in-cache-re-check",
		ContentLength: 1,
		LastModified: dt,
	}, {
		URL: "https://in-cache-re-check1",
		ContentLength: 1,
		LastModified: time.Now(),
	}})
	g.On("Execute", mock.Anything).Return([]models.Feed{{},
	})
	cache.On("Save").Return(nil)
	New(c, cache, h, g).Execute()

	config.DeleteConfigFiles()

	_ = os.Setenv("TEST", "")
}
