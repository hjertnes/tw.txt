package edit

import (
	"testing"
)

func TestTest(t *testing.T) {
	conf, _ := config.New()

	New(conf).Execute("")

	New(conf).Execute("internal-config")

	New(conf).Execute("common-config")
}