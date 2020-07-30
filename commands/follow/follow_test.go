package follow

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

	New(c).Execute("a", "b")
	assert.Equal(t, "b", c.CommonConfig.Following["a"])

	config.DeleteConfigFiles()

	_ = os.Setenv("TEST", "")
}