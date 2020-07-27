// Package utils contains various utility functions
package utils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"git.sr.ht/~hjertnes/tw.txt/constants"
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
	return strings.ReplaceAll(input, "~", home)
}

// ParseArgs Parses args.
func ParseArgs(args []string) (string, error) {
	if len(args) < 2 {
		return "", constants.ErrTooFewArgs
	}

	return args[1], nil
}

// ErrorHandler Handles errors with panic.
func ErrorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

// PrettyDuration Pretty print a duration.
func PrettyDuration(duration time.Duration) string {
	s := int(duration.Seconds())
	d := s / 86400

	s %= 86400

	if d >= 365 {
		return fmt.Sprintf("%dy %dw ago", d/365, d%365/7)
	}

	if d >= 14 {
		return fmt.Sprintf("%dw ago", d/7)
	}

	h := s / 3600

	s %= 3600

	if d > 0 {
		str := fmt.Sprintf("%dd", d)
		if h > 0 && d <= 6 {
			str += fmt.Sprintf(" %dh", h)
		}

		return str + " ago"
	}

	m := s / 60

	s %= 60

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
		parts := strings.Split(line, "\t")
		if len(parts) < 2 {
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
