package main

import (
	"os"
	"testing"

	"git.sr.ht/~hjertnes/tw.txt/config"
)

func TestTest(t *testing.T) {
	t.Run("", func(t *testing.T) {
		_ = os.Setenv("TEST", "true")

		runProgram([]string{""})

		runProgram([]string{"", "setup"})

		config.DeleteConfigFiles()

		config.CreateConfigFiles()

		runProgram([]string{"", "timeline"})

		runProgram([]string{"", "edit"})

		runProgram([]string{"", "follow"})

		runProgram([]string{"", "follow", "a", "b"})

		runProgram([]string{"", "unfollow"})

		runProgram([]string{"", "unfollow", "a", "b"})

		runProgram([]string{"", "tweet"})

		runProgram([]string{"", "tweet", "a"})

		runProgram([]string{"", "replace-mentions"})

		runProgram([]string{"", "following"})

		runProgram([]string{"", "test-feeds"})

		config.DeleteConfigFiles()

		_ = os.Setenv("TEST", "")
	})
}
