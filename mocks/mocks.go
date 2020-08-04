// Package mocks contains interface mocks
package mocks

import (
	"time"

	"git.sr.ht/~hjertnes/tw.txt/constants"
	"git.sr.ht/~hjertnes/tw.txt/models"
	"github.com/stretchr/testify/mock"
)

// LoadFeedsMock is a Mock for the Service loadfeeds.Service.
type LoadFeedsMock struct {
	mock.Mock
}

// Execute mock.
func (l *LoadFeedsMock) Execute() []models.Feed {
	args := l.Called()

	return args.Get(constants.FirstArgument).([]models.Feed)
}

// ConfigMock is a mock for config Service.
type ConfigMock struct {
	mock.Mock
}

// Get is a mock.
func (c *ConfigMock) Get() *models.Config {
	args := c.Called()

	return args.Get(constants.FirstArgument).(*models.Config)
}

// Save is a mock.
func (c *ConfigMock) Save() error {
	args := c.Called()

	return args.Error(constants.FirstArgument)
}

// GetFeedsMock is mock for headfeed package.
type GetFeedsMock struct {
	mock.Mock
}

// Execute is a mock.
func (g *GetFeedsMock) Execute(feeds map[string]string) []models.Feed {
	args := g.Called(feeds)

	return args.Get(constants.FirstArgument).([]models.Feed)
}

// HeadFeedsMock is mock for headfeed package.
type HeadFeedsMock struct {
	mock.Mock
}

// Execute is a mock.
func (h *HeadFeedsMock) Execute(feeds map[string]string) []models.FeedHead {
	args := h.Called(feeds)

	return args.Get(constants.FirstArgument).([]models.FeedHead)
}

// CacheMock is a mock for Cache.Service.
type CacheMock struct {
	mock.Mock
}

// Get is a mock.
func (c *CacheMock) Get(url string) (*models.CachedUser, error) {
	args := c.Called(url)

	return args.Get(constants.FirstArgument).(*models.CachedUser), args.Error(constants.SecondArgument)
}

// Set is a mock.
func (c *CacheMock) Set(handle string, url string, content string, contentLength int, lastUpdated time.Time) {
	_ = c.Called(handle, url, content, contentLength, lastUpdated)
}

// Save is a mock.
func (c *CacheMock) Save() error {
	args := c.Called()

	return args.Error(constants.FirstArgument)
}
