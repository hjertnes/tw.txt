package edit

import (
	"git.sr.ht/~hjertnes/tw.txt/config"
	"os"
	"testing"
)

func TestTest(t *testing.T) {
	conf, _ := config.New()

	e := os.Getenv("EDITOR")

	_ = os.Setenv("EDITOR", "echo")

	New(conf).Execute("")

	New(conf).Execute("internal-config")

	New(conf).Execute("common-config")

	_ = os.Setenv("EDITOR", e)
}