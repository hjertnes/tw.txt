// Package utils contains various utility functions
package utils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"git.sr.ht/~hjertnes/tw.txt/models"
	"github.com/goware/urlx"
)

// Exist Checks if file or folder exists.
func Exist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// ReplaceTilde Replaces tilde with value of HOME.
func ReplaceTilde(input string) string {
	home := os.Getenv("HOME")
	return strings.Replace(input, "~", home, 1)
}

const (
	one = 1
	two = 2
)

// ParseArgs Parses args.
func ParseArgs(args []string) (string, string, []string) {
	command := ""
	subCommand := ""
	r := make([]string, 0)

	if len(args) > one {
		command = args[one]
	}

	if len(args) > two {
		subCommand = args[two]
		r = args[two:]
	}

	return command, subCommand, r
}

// ErrorHandler Handles errors with panic.
func ErrorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	daysInAYear     = 365
	daysInAWeek     = 7
	daysInTwoWeeks  = 14
	secondsInADay   = 86400
	secondsInAnHour = 3600
	sixDays         = 6
	minutesInAnHour = 60
)

// PrettyDuration Pretty print a duration.
func PrettyDuration(duration time.Duration) string {
	s := int(duration.Seconds())
	d := s / secondsInADay

	s %= secondsInADay

	if d >= daysInAYear {
		return fmt.Sprintf("%dy %dw ago", d/daysInAYear, d%daysInAYear/daysInAWeek)
	}

	if d >= daysInTwoWeeks {
		return fmt.Sprintf("%dw ago", d/daysInAWeek)
	}

	h := s / secondsInAnHour

	s %= secondsInAnHour

	if d > 0 {
		str := fmt.Sprintf("%dd", d)
		if h > 0 && d <= sixDays {
			str += fmt.Sprintf(" %dh", h)
		}

		return str + " ago"
	}

	m := s / minutesInAnHour

	s %= minutesInAnHour

	if h > 0 || m > 0 {
		str := ""

		hh := ""

		if h > 0 {
			str += fmt.Sprintf("%dh", h)
			hh = " "
		}

		if m > 0 {
			str += fmt.Sprintf("%s%dm", hh, m)
		}

		return str + " ago"
	}

	return fmt.Sprintf("%ds ago", s)
}

// NormalizeURL NormalizeURL.
func NormalizeURL(url string) string {
	if url == "" {
		return ""
	}

	u, err := urlx.Parse(url)
	if err != nil {
		return ""
	}

	if u.Scheme == "https" {
		u.Scheme = "http"
		u.Host = strings.TrimSuffix(u.Host, ":443")
	}

	u.User = nil
	u.Path = strings.TrimSuffix(u.Path, "/")

	norm, err := urlx.Normalize(u)
	if err != nil {
		return ""
	}

	return norm
}

// ParseFile Parses a twtxt line into Tweet.
func ParseFile(handle string, url string, lines []string) []models.Tweet {
	o := make([]models.Tweet, 0)

	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, "\t")
		if len(parts) < two {
			continue
		}

		timestamp, _ := time.Parse(time.RFC3339, parts[0])
		message := parts[1]

		o = append(o, models.Tweet{
			Timestamp: timestamp,
			Handle:    handle,
			Message:   message,
			URL:       url,
		})
	}

	return o
}
