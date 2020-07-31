package following

import (
	"git.sr.ht/~hjertnes/tw.txt/config"
	"testing"
)

func TestTest(t *testing.T) {
	c := &config.Config{
		CommonConfig: &config.CommonConfig{
			Following: map[string]string{},
		},
		InternalConfig: &config.InternalConfig{},
	}
	c.CommonConfig.Following["hjertnes"] = "https://hjertnes.social"
	New(c).Execute()
}