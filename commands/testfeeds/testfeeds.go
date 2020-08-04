// Package testfeeds tests status code of a feed
package testfeeds

import (
	"fmt"

	"git.sr.ht/~hjertnes/tw.txt/loadfeeds"

	"git.sr.ht/~hjertnes/tw.txt/output"
)

// Command is the publicly exposed interface.
type Command interface {
	Execute()
}

type command struct {
	loadFeeds loadfeeds.Service
}

func (c *command) Execute() {
	statuses := c.loadFeeds.Execute()

	for _, status := range statuses {
		if status.Status {
			fmt.Println(output.Green(fmt.Sprintf("@<%s %s>: OK", status.Handle, status.URL)))
		} else {
			fmt.Println(output.Red(fmt.Sprintf("@<%s %s>: Error", status.Handle, status.URL)))
		}
	}
}

// New creates new Command.
func New(lf loadfeeds.Service) Command {
	return &command{loadFeeds: lf}
}
