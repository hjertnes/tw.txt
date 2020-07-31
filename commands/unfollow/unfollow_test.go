package unfollow

import (
	"git.sr.ht/~hjertnes/tw.txt/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTest(t *testing.T) {
	_ = os.Setenv("TEST", "true")

	config.CreateConfigFiles()

	c, _ := config.New()

	c.CommonConfig.Following["a"] = "b"

	New(c).Execute("a")
	assert.Equal(t, "", c.CommonConfig.Following["a"])

	config.DeleteConfigFiles()

	_ = os.Setenv("TEST", "")
}