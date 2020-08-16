// Package constants contains constant values
package constants

import (
	"errors"
	"time"
)

// ErrTooFewArgs Error for too few args.
var ErrTooFewArgs = errors.New("too few command line arguments")

// ErrConfigDoesNotExist Error for missing config file.
var ErrConfigDoesNotExist = errors.New("config does not exist")

var (
	// ErrNotInCache error for when a key is not in cache
	ErrNotInCache = errors.New("not in cache")

	// ErrExpired error for when a key has expired
	ErrExpired = errors.New("in cache but expired")

	// ErrFetchHead error for when it is in cache, has not expired, but should be re-evaluated
	ErrFetchHead = errors.New("in cache but should fetch head and re-validate")
)

// Name Of The thing.
const Name = "tw.txt"

// Version of the app.
const Version = "0.5.4"

// Two Constant for the number two.
const Two = 2

const twentyFour = 24
const five = 5

// HttpClientTimeout is the timeout used on all http clients
const HttpClientTimeout = time.Second * five

// OneDay Duration for a day.
const OneDay = time.Hour * twentyFour

const (
	// FirstArgument is the first arg when you need to reference then as a zero indexed int.
	FirstArgument = 0

	// SecondArgument is the first arg when you need to reference then as a zero indexed int.
	SecondArgument = 1
)
