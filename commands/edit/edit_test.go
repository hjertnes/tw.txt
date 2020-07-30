package edit

import (
	"git.sr.ht/~hjertnes/tw.txt/config"
	"testing"
)

func TestTest(t *testing.T){
	conf, _ := config.New()

	New(conf).Execute("")

	New(conf).Execute("internal-config")

	New(conf).Execute("common-config")
}