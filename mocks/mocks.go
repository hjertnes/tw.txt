// Package mocks contains interface mocks
package mocks

import (
	"git.sr.ht/~hjertnes/tw.txt/constants"
	"git.sr.ht/~hjertnes/tw.txt/models"
	"github.com/stretchr/testify/mock"
	"time"
)

// LoadFeedsMock is a Mock for the Service loadfeeds.Service
type LoadFeedsMock struct {
	mock.Mock
}
// Execute mock
func (l *LoadFeedsMock) Execute() []models.Feed {
	args := l.Called()

	return args.Get(constants.FirstArgument).([]models.Feed)
}

type ConfigMock struct {
	mock.Mock
}

func (c *ConfigMock) Get() *models.Config{
	args := c.Called()

	return args.Get(constants.FirstArgument).(*models.Config)
}

func (c *ConfigMock) Save() error{
	args := c.Called()

	return args.Error(constants.FirstArgument)
}

type GetFeedsMock struct {
	mock.Mock
}

func (g *GetFeedsMock) Execute(feeds map[string]string) []models.Feed{
	args := g.Called(feeds)

	return args.Get(constants.FirstArgument).([]models.Feed)
}

type HeadFeedsMock struct {
	mock.Mock
}

func (h *HeadFeedsMock) Execute(feeds map[string]string) []models.FeedHead{
	args := h.Called(feeds)

	return args.Get(constants.FirstArgument).([]models.FeedHead)
}

type CacheMock struct {
	mock.Mock
}

type Service interface {
	Get(url string) (*models.CachedUser, error)
	Set(handle string, url string, content string, contentLength int, lastUpdated time.Time)
	Save() error
}

func (c *CacheMock) Get(url string) (*models.CachedUser, error){
	args := c.Called(url)

	return args.Get(constants.FirstArgument).(*models.CachedUser), args.Error(constants.SecondArgument)
}

func (c *CacheMock) Set(handle string, url string, content string, contentLength int, lastUpdated time.Time){
	_ = c.Called(handle, url, content, contentLength, lastUpdated)
}

func (c *CacheMock) Save() error{
	args := c.Called()

	return args.Error(constants.FirstArgument)
}