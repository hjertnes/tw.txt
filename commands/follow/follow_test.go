package follow

import (
	"fmt"
	"os"
	"testing"
)

func TestTest(t *testing.T) {
	_ = os.Setenv("TEST", "true")

	config.CreateConfigFiles()

	c, err := config.New()
	fmt.Println(err)

	New(c).Execute("a", "b")
	assert.Equal(t, "b", c.CommonConfig.Following["a"])

	config.DeleteConfigFiles()

	_ = os.Setenv("TEST", "")
}