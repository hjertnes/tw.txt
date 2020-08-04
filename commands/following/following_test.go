package following

import (
	"git.sr.ht/~hjertnes/tw.txt/models"
	"testing"
)

func TestTest(t *testing.T) {
	c := &models.Config{
		CommonConfig: &models.CommonConfig{
			Following: map[string]string{},
		},
		InternalConfig: &models.InternalConfig{},
	}
	c.CommonConfig.Following["hjertnes"] = "https://hjertnes.social"
	New(c).Execute()
}