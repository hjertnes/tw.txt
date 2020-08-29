// Package cache is the feed caching package.
package cache

import (
	"io/ioutil"
	"os"
	"time"

	"git.sr.ht/~hjertnes/tw.txt/constants"
	"git.sr.ht/~hjertnes/tw.txt/models"
	"git.sr.ht/~hjertnes/tw.txt/utils"
	"gopkg.in/yaml.v3"
)

// Service is the exposed interface.
type Service interface {
	Get(url string) (*models.CachedUser, error)
	Set(handle string, url string, content string, contentLength int, lastUpdated time.Time)
	Update(handle string, url string, content string, contentLength int, lastUpdated time.Time, expire time.Time)
	Save() error
}

type service struct {
	data *models.CacheFile
}

func (s *service) Get(url string) (*models.CachedUser, error) {
	user := s.data.Users[url]
	if user == nil {
		return nil, constants.ErrNotInCache
	}

	if user.Expire.Before(time.Now()) {
		return nil, constants.ErrExpired
	}

	if user.NextCheck.Before(time.Now()) {
		return user, constants.ErrFetchHead
	}

	return user, nil
}

func (s *service) Set(handle string, url string, content string, contentLength int, lastUpdated time.Time) {
	s.data.Users[url] = &models.CachedUser{
		Handle:        handle,
		URL:           url,
		Content:       content,
		ContentLength: contentLength,
		LastUpdated:   lastUpdated,
		NextCheck:     time.Now().Add(time.Minute * constants.Two),
		Expire:        time.Now().Add(constants.OneDay),
	}
}

func (s *service) Update(handle string, url string, content string, contentLength int, lastUpdated time.Time, expire time.Time) {
	s.data.Users[url] = &models.CachedUser{
		Handle:        handle,
		URL:           url,
		Content:       content,
		ContentLength: contentLength,
		LastUpdated:   lastUpdated,
		NextCheck:     time.Now().Add(time.Minute * constants.Two),
		Expire:        expire,
	}
}

func (s *service) Save() error {
	filename := "~/.tw.txt/cache.yaml"
	if os.Getenv("TEST") != "" {
		filename = "~/.tw.txt-test/cache.yaml"
	}

	cc, err := yaml.Marshal(s.data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(utils.ReplaceTilde(filename), cc, 0)

	return err
}

// New is the constructor.
func New() (Service, error) {
	filename := "~/.tw.txt/cache.yaml"
	if os.Getenv("TEST") != "" {
		filename = "~/.tw.txt-test/cache.yaml"
	}

	if !utils.Exist(utils.ReplaceTilde(filename)) {
		f, err := os.Create(utils.ReplaceTilde(filename))
		if err != nil {
			return nil, err
		}

		c := &models.CacheFile{
			Users: make(map[string]*models.CachedUser),
		}

		cc, err := yaml.Marshal(c)
		if err != nil {
			return nil, err
		}

		_, err = f.Write(cc)
		if err != nil {
			return nil, err
		}

		err = f.Close()
		if err != nil {
			return nil, err
		}
	}

	content, err := ioutil.ReadFile(utils.ReplaceTilde(filename))
	if err != nil {
		return nil, err
	}

	c := &models.CacheFile{}

	err = yaml.Unmarshal(content, c)
	if err != nil {
		return nil, err
	}

	return &service{data: c}, nil
}
