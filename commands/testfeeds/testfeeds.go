// Package testfeeds tests status code of a feed
package testfeeds

import (
	"fmt"

	"git.sr.ht/~hjertnes/tw.txt/output"
	"git.sr.ht/~hjertnes/tw.txt/services/fetchfeeds"
)

// Command is the publicly exposed interface.
type Command interface {
	Execute()
}

type command struct {
	fetchFeeds fetchfeeds.Command
}

func (c *command) Execute() {
	statuses := c.fetchFeeds.Execute("Testing feeds...")

	for _, status := range statuses {
		if status.Status {
			fmt.Println(output.Green(fmt.Sprintf("@<%s %s>: OK", status.Handle, status.URL)))
		} else {
			fmt.Println(output.Red(fmt.Sprintf("@<%s %s>: Error", status.Handle, status.URL)))
		}
	}
}

// New creates new Command.
func New(ff fetchfeeds.Command) Command {
	return &command{fetchFeeds: ff}
}
