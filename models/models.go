// Package models contains models
package models

import "time"

// Tweet is the type for a line of a twtxt file.
type Tweet struct {
	Timestamp time.Time
	Handle    string
	URL       string
	Message   string
}

// Feed is the status of a request to a feed.
type Feed struct {
	Handle string
	URL    string
	Status bool
	Body   string
}