package utils

import (
	"fmt"
	"git.sr.ht/~hjertnes/tw.txt/constants"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"

)

func TestExist(t *testing.T) {
	t.Run("Exists", func(t *testing.T) {
		assert.True(t, Exist("/tmp"))
	})

	t.Run("Does not exists", func(t *testing.T) {
		assert.False(t, Exist("/mp"))
	})
}

func TestReplaceTilde(t *testing.T) {
	testString := "~/~/~"
	expectedString := fmt.Sprintf("%s/~/~", os.Getenv("HOME"))
	assert.Equal(t, expectedString, ReplaceTilde(testString))
}

func TestParseArgs(t *testing.T) {
	t.Run("one arg", func(t *testing.T) {
		cmd, subCmd, args := ParseArgs([]string{"one"})
		assert.Equal(t, cmd, "")
		assert.Equal(t, subCmd, "")
		assert.Len(t, args, 0)
	})
	t.Run("two args", func(t *testing.T) {
		cmd, subCmd, args := ParseArgs([]string{"one", "two"})
		assert.Equal(t, cmd, "two")
		assert.Equal(t, subCmd, "")
		assert.Len(t, args, 0)
	})
	t.Run("three args", func(t *testing.T) {
		cmd, subCmd, args := ParseArgs([]string{"one", "two", "three"})
		assert.Equal(t, cmd, "two")
		assert.Equal(t, subCmd, "three")
		assert.Len(t, args, 1)
		assert.Equal(t, args[0], "three")
	})
	t.Run("four args", func(t *testing.T) {
		cmd, subCmd, args := ParseArgs([]string{"one", "two", "three", "four"})
		assert.Equal(t, cmd, "two")
		assert.Equal(t, subCmd, "three")
		assert.Len(t, args, 2)
		assert.Equal(t, args[0], "three")
		assert.Equal(t, args[1], "four")
	})
	t.Run("five args", func(t *testing.T) {
		cmd, subCmd, args := ParseArgs([]string{"one", "two", "three", "four", "five"})
		assert.Equal(t, cmd, "two")
		assert.Equal(t, subCmd, "three")
		assert.Len(t, args, 3)
		assert.Equal(t, args[0], "three")
		assert.Equal(t, args[1], "four")
		assert.Equal(t, args[2], "five")
	})
}

func TestErrorHandler(t *testing.T) {
	t.Run("Is null", func(t *testing.T) {
		ErrorHandler(nil)
	})
	t.Run("Is not null", func(t *testing.T) {
		assert.Panics(t, func() {
			ErrorHandler(constants.ErrTooFewArgs)
		})
	})
}

func TestParseFile(t *testing.T) {
	testData := []string{
		"#",
		"2020-07-30T11:51:36+02:00	time to write some unit tests for tw.txt",
		"",
	}

	result := ParseFile("hjertnes", "https://hjertnes.social/twtxt.txt", testData)
	assert.Len(t, result, 1)
}

func TestNormalizeURL(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, "", NormalizeURL(""))
	})
	t.Run("invalid url", func(t *testing.T) {
		assert.Equal(t, "", NormalizeURL("not#a#!%url"))
	})
	t.Run("without prefix", func(t *testing.T) {
		assert.Equal(t, "http://url", NormalizeURL("url"))
	})
	t.Run("with http prefix", func(t *testing.T) {
		assert.Equal(t, "http://url", NormalizeURL("http://url"))
	})
	t.Run("with https prefix", func(t *testing.T) {
		assert.Equal(t, "http://url", NormalizeURL("https://url"))
	})
}

func TestPrettyDuration(t *testing.T) {
	t.Run("Second", func(t *testing.T) {
		assert.Equal(t, "1s ago", PrettyDuration(time.Second))
	})

	t.Run("1h 59m", func(t *testing.T) {
		assert.Equal(t, "1h 59m ago", PrettyDuration(time.Second*secondsInAnHour+(time.Second*secondsInAnHour-30)))
	})

	t.Run("15 days ago", func(t *testing.T) {
		assert.Equal(t, "2w ago", PrettyDuration(time.Second*secondsInADay*daysInTwoWeeks+1))
	})

	t.Run("6 days ago", func(t *testing.T) {
		assert.Equal(t, "6d 23h ago", PrettyDuration((time.Second*secondsInADay*daysInAWeek)-(time.Second*secondsInAnHour)))
	})

	t.Run("Two years ago", func(t *testing.T) {
		assert.Equal(t, "2y 0w ago", PrettyDuration(time.Second*secondsInADay*daysInAYear*2))
	})
}
