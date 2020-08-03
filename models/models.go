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