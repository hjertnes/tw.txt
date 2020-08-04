// Package models contains models
package models

import (
	"html/template"
	"time"
)

// Tweet is the type for a line of a twtxt file.
type Tweet struct {
	Timestamp time.Time
	Handle    string
	URL       string
	Message   string
}

// HTMLTweet is like Tweet but for generating timeline as HTML.
type HTMLTweet struct {
	Timestamp time.Time
	Handle    string
	URL       string
	Message   template.HTML
	Classes   string
}

// Feed is the status of a request to a feed.
type Feed struct {
	ContentLength int
	LastModified  time.Time
	Handle        string
	URL           string
	Status        bool
	Body          string
}

// FeedHead is the type returned after a HEAD request to a feed.
type FeedHead struct {
	ContentLength int
	LastModified  time.Time
	URL           string
}

// HTMLModel is a the model used when generating a timelien to html.
type HTMLModel struct {
	Timestamp time.Time
	Timeline  []HTMLTweet
}

// CommonConfig is a shared config intended to be supported by all twtxt clients.
type CommonConfig struct {
	Nick             string
	URL              string
	File             string
	Following        map[string]string
	DiscloseIdentity bool
}

// InternalConfig config file used by this client: located at ~/.tw.txt/config.yaml.
type InternalConfig struct {
	ConfigFileLocation   string
	TemplateFileLocation string
}

// Config Type config contains CommonConfig and InternalConfig.
type Config struct {
	InternalConfig *InternalConfig
	CommonConfig   *CommonConfig
}

// CachedUser is the model for a cached feed.
type CachedUser struct {
	Handle        string
	URL           string
	Content       string
	NextCheck     time.Time
	Expire        time.Time
	LastUpdated   time.Time
	ContentLength int
}

// CacheFile is the cache file model.
type CacheFile struct {
	Users map[string]*CachedUser
}
