// Package constants contains constant values
package constants

import (
	"errors"
)

// ErrTooFewArgs Error for too few args.
var ErrTooFewArgs = errors.New("too few command line arguments")

// ErrConfigDoesNotExist Error for missing config file.
var ErrConfigDoesNotExist = errors.New("config does not exist")

// Name Of The thing.
const Name = "tw.txt"

// Version of the app.
const Version = "0.5.4"

// Two Constant for the number two.
const Two = 2
