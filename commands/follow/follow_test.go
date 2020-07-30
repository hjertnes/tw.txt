package follow

import (
	"fmt"
	"git.sr.ht/~hjertnes/tw.txt/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTest(t *testing.T){
	_ = os.Setenv("TEST", "true")

	config.CreateConfigFiles()

	c, err := config.New()
	fmt.Println(err)

	New(c).Execute("a", "b")
	assert.Equal(t, "b", c.CommonConfig.Following["a"])

	config.DeleteConfigFiles()
	_ = os.Setenv("TEST", "")
}