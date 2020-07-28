// Package constants contains constant values
package constants

import "errors"

// ErrTooFewArgs Error for too few args.
var ErrTooFewArgs = errors.New("too few command line arguments")

// ErrConfigDoesNotExist Error for missing config file.
var ErrConfigDoesNotExist = errors.New("config does not exist")

//NameOfTheApp
const Name = "tw.txt"
const Version = "0.2.0"