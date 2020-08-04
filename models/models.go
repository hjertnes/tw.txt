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

type HTMLTweet struct {
	Timestamp time.Time
	Handle    string
	URL       string
	Message   template.HTML
	Classes string
}

// Feed is the status of a request to a feed.
type Feed struct {
	ContentLength int
	LastModified time.Time
	Handle string
	URL    string
	Status bool
	Body   string
}

type FeedHead struct {
	ContentLength int
	LastModified time.Time
	URL    string
}

type HTMLModel struct {
	Timestamp time.Time
	Timeline []HTMLTweet
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
	ConfigFileLocation string
	TemplateFileLocation string
}

// Config Type config contains CommonConfig and InternalConfig.
type Config struct {
	InternalConfig *InternalConfig
	CommonConfig   *CommonConfig
}

type CachedUser struct {
	Handle string
	URL string
	Content string
	NextCheck time.Time
	Expire time.Time // 24h
	LastUpdated time.Time
	ContentLength int
}

type CacheFile struct {
	Users map[string]*CachedUser
}