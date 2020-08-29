// Package loadfeeds is the package that loads feeds and deals with caching, re-evaluation etc
package loadfeeds

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"

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

func buildLogger(writer io.WriteCloser) *log.Logger{
	l := log.New()
	l.SetReportCaller(true)
	l.SetOutput(writer)
	return l
}

type service struct {
	logger *log.Logger
	writer io.WriteCloser
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
		_, err := s.cache.Get(url)

		if err != nil {
			if errors.Is(err, constants.ErrExpired) || errors.Is(err, constants.ErrNotInCache) {
				s.logger.Info(fmt.Sprintf("%s is expired", url))
				feedsToGet[handle] = url
			}

			if errors.Is(err, constants.ErrFetchHead) {
				s.logger.Info(fmt.Sprintf("%s is should be re-evaluated", url))
				feedsToHead[handle] = url
			}
		} else {
			panic(err)
		}
	}

	for _, headData := range s.headFeeds.Execute(feedsToHead) {
		s.logger.Info(fmt.Sprintf("HEAD %s", headData.URL))
		d, _ := s.cache.Get(headData.URL)

		if d.ContentLength != headData.ContentLength || headData.LastModified.After(d.LastUpdated) {
			s.logger.Info(fmt.Sprintf("Should GET %s", headData.URL))
			feedsToGet[d.Handle] = d.URL
		} else {
			s.logger.Info(fmt.Sprintf("Keeping cache for %s", headData.URL))
			s.cache.Update(d.Handle, d.URL, d.Content, d.ContentLength, d.LastUpdated, d.Expire)
			data = append(data, FromCachedUser(d))
		}
	}

	for _, getData := range s.getFeeds.Execute(feedsToGet) {
		s.logger.Info(fmt.Sprintf("GET %s", getData.URL))
		s.cache.Set(getData.Handle, getData.URL, getData.Body, getData.ContentLength, getData.LastModified)
		data = append(data, getData)
	}

	err := s.cache.Save()
	if err != nil {
		utils.ErrorHandler(err)
	}

	_ = s.writer.Close()
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
func New(
	writer io.WriteCloser,
	config config.Service,
	cache cache.Service,
	headFeeds headfeeds.Command,
	getFeeds getfeeds.Command,
	) Service {

	return &service{
		buildLogger(writer),
		writer,
		config,
		cache,
		headFeeds,
		getFeeds,
	}
}
