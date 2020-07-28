// Package help exposes a command for showing help message
package help

import "fmt"

// Command is the exposed interface of this package.
type Command interface {
	Execute()
}

type command struct {
}

func (c *command) Execute() {
	fmt.Println("tw.txt is another twtxt client -- https://twtxt.readthedocs.org/en/stable")
	fmt.Println("Usage:")
	fmt.Println("\t tw.txt command")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("\ttimeline - prints last 1000 items in your timeline")
	fmt.Println("\t\t full - prints your entire timeline")
	fmt.Println("\tedit - opens your twtxt in $EDITOR")
	fmt.Println("\t\t twtxt - opens your twtxt in $EDITOR")
	fmt.Println("\t\t common-config - opens common-config in EDITOR")
	fmt.Println("\t\t internal-config - opens tw.txt config in $EDITOR")
	fmt.Println("\tsetup - creates config file and opens it in $EDITOR")
	fmt.Println("\tfollow [nick] [url] - follows the specified feed")
	fmt.Println("\tunfollow [nick] - unfollows the specified feed")
	fmt.Println("\ttweet [message] - posts a new row to your twtxt; will replace any @handle with @<handle url>")
	fmt.Println("\treplace-mentions - will replace all @handle with @<handle url>")
}

// New is the constructor.
func New() Command {
	return &command{}
}
