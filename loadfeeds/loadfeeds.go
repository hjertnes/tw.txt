// Package loadfeeds is the package that loads feeds and deals with caching, re-evaluation etc
package loadfeeds

import (
	"errors"

	"git.sr.ht/~hjertnes/tw.txt/config"

	"git.sr.ht/~hjertnes/tw.txt/constants"
	"git.sr.ht/~hjertnes/tw.txt/loadfeeds/cache"
	"git.sr.ht/~hjertnes/tw.txt/loadfeeds/getfeeds"
	"git.sr.ht/~hjertnes/tw.txt/loadfeeds/headfeeds"
	"git.sr.ht/~hjertnes/tw.txt/models"
	"git.sr.ht/~hjertnes/tw.txt/utils"
)

// Service is the exposed interface.
type Service interface {
	Execute() []models.Feed
}

type service struct {
	config    config.Service
	cache     cache.Service
	headFeeds headfeeds.Command
	getFeeds  getfeeds.Command
}

func (s *service) Execute() []models.Feed {
	feeds := s.config.Get().CommonConfig.Following
	feeds[s.config.Get().CommonConfig.Nick] = s.config.Get().CommonConfig.URL

	feedsToGet := make(map[string]string)
	feedsToHead := make(map[string]string)

	data := make([]models.Feed, 0)

	for handle, url := range feeds {
		d, err := s.cache.Get(url)

		if err != nil {
			if errors.Is(err, constants.ErrExpired) || errors.Is(err, constants.ErrNotInCache) {
				feedsToGet[handle] = url
			}

			if errors.Is(err, constants.ErrFetchHead) {
				feedsToHead[handle] = url
			}
		} else {
			s.cache.Set(d.Handle, d.URL, d.Content, d.ContentLength, d.LastUpdated)
			data = append(data, FromCachedUser(d))
		}
	}

	for _, headData := range s.headFeeds.Execute(feedsToHead) {
		d, _ := s.cache.Get(headData.URL)

		if d.ContentLength != headData.ContentLength && headData.LastModified.After(d.LastUpdated) {
			feedsToGet[d.Handle] = d.URL
		} else {
			s.cache.Set(d.Handle, d.URL, d.Content, d.ContentLength, d.LastUpdated)
			data = append(data, FromCachedUser(d))
		}
	}

	for _, getData := range s.getFeeds.Execute(feedsToGet) {
		s.cache.Set(getData.Handle, getData.URL, getData.Body, getData.ContentLength, getData.LastModified)
		data = append(data, getData)
	}

	err := s.cache.Save()
	if err != nil {
		utils.ErrorHandler(err)
	}

	return data
}

// FromCachedUser Builds Feed from CachedUser.
func FromCachedUser(d *models.CachedUser) models.Feed {
	return models.Feed{
		Handle:        d.Handle,
		URL:           d.URL,
		Status:        true,
		Body:          d.Content,
		LastModified:  d.LastUpdated,
		ContentLength: d.ContentLength,
	}
}

// New is the constructor.
func New(config config.Service, cache cache.Service, headFeeds headfeeds.Command, getFeeds getfeeds.Command) Service {
	return &service{
		config, cache, headFeeds, getFeeds,
	}
}
