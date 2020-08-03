package loadfeeds

import (
	"errors"
	"fmt"
	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/loadfeeds/cache"
	"git.sr.ht/~hjertnes/tw.txt/loadfeeds/getfeeds"
	"git.sr.ht/~hjertnes/tw.txt/loadfeeds/headfeeds"
	"git.sr.ht/~hjertnes/tw.txt/models"
	"git.sr.ht/~hjertnes/tw.txt/utils"
)

type Service interface {
	Execute() []models.Feed
}

type service struct {
	config *config.Config
	cache cache.Service
	headFeeds headfeeds.Command
	getFeeds getfeeds.Command
}

func (s *service) Execute() []models.Feed {
	feeds := s.config.CommonConfig.Following
	feeds[s.config.CommonConfig.Nick] = s.config.CommonConfig.URL

	feedsToGet := make(map[string]string)
	feedsToHead := make(map[string]string)


	data := make([]models.Feed, 0)

	for handle, url := range feeds{
		fmt.Println(s.cache)
		d, err := s.cache.Get(url)

		if err != nil{
			if errors.Is(err, cache.ErrorExpired) || errors.Is(err, cache.ErrorNotInCache){
				feedsToGet[handle] = url
			}

			if errors.Is(err, cache.ErrorFetchHead){
				feedsToHead[handle] = url
			}
		} else {
			data = append(data, models.Feed{
				Handle: handle,
				URL: url,
				Status: true,
				Body: d.Content,
				LastModified: d.LastUpdated,
				ContentLength: d.ContentLength,
			})
		}
	}

	for _, headData := range s.headFeeds.Execute(feedsToHead){
		d, _ := s.cache.Get(headData.URL)

		if d.ContentLength != headData.ContentLength && headData.LastModified.After(d.LastUpdated){
			feedsToHead[d.Handle] = d.URL
		}
	}

	for _, getData := range s.getFeeds.Execute(feedsToGet){
		s.cache.Set(getData.Handle, getData.URL, getData.Body, getData.ContentLength, getData.LastModified)
		data = append(data, getData)
	}

	err := s.cache.Save()
	if err != nil{
		utils.ErrorHandler(err)
	}

	return data
}


func New(config *config.Config, cache cache.Service, headFeeds headfeeds.Command, getFeeds getfeeds.Command) Service{
	return &service{
		config, cache, headFeeds, getFeeds,
	}
}