package cache

import (
	"errors"
	"git.sr.ht/~hjertnes/tw.txt/utils"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"time"
)

type CachedUser struct {
	Handle string
	URL string
	Content string
	NextCheck time.Time
	Expire time.Time // 24h
	LastUpdated time.Time
	ContentLength int
}

var ErrorNotInCache = errors.New("not in cache")
var ErrorExpired = errors.New("in cache but expired")
var ErrorFetchHead = errors.New("in cache but should fetch head and re-validate")

type CacheFile struct {
	Users map[string]*CachedUser
}

type Service interface {
	Get(url string) (*CachedUser, error)
	Set(handle string, url string, content string, contentLength int, lastUpdated time.Time)
	Save() error

}

type service struct {
	data *CacheFile
}

func(s *service) Get(url string) (*CachedUser, error){
	user := s.data.Users[url]
	if user == nil {
		return nil, ErrorNotInCache
	}

	if user.Expire.Before(time.Now()){
		return nil, ErrorExpired
	}

	if user.NextCheck.Before(time.Now()){
		return user, ErrorFetchHead
	}

	return user, nil
}

func (s *service) Set(handle string, url string, content string, contentLength int, lastUpdated time.Time){
	s.data.Users[url] = &CachedUser{
		Handle: handle,
		URL: url,
		Content: content,
		ContentLength: contentLength,
		LastUpdated: lastUpdated,
		NextCheck: time.Now().Add(time.Minute * 2),
		Expire: time.Now().Add(time.Hour * 24),
	}
}

func (s *service) Save() error{
	filename := "~/.tw.txt/cache.yaml"
	cc, err := yaml.Marshal(s.data)
	if err != nil{
		return err
	}

	err = ioutil.WriteFile(utils.ReplaceTilde(filename), cc, 0)
	return err
}

func New() (Service, error){
	filename := "~/.tw.txt/cache.yaml"
	if !utils.Exist(utils.ReplaceTilde(filename)){
		f, err :=  os.Create(utils.ReplaceTilde(filename))
		if err != nil{
			return nil, err
		}

		c := &CacheFile{
			Users: make(map[string]*CachedUser),
		}

		cc, err := yaml.Marshal(c)
		if err != nil{
			return nil, err
		}

		_, err = f.Write(cc)
		if err != nil{
			return nil, err
		}

		err = f.Close()
		if err != nil{
			return nil, err
		}
	}

	content, err := ioutil.ReadFile(utils.ReplaceTilde(filename))
	if err != nil {
		return nil, err
	}

	c := &CacheFile{}

	err = yaml.Unmarshal(content, c)
	if err != nil {
		return nil, err
	}

	return &service{data: c}, nil
}