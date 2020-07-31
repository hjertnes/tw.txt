package tweet

import (
	"fmt"
	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/utils"
	"os"
	"testing"
)

func TestTest(t *testing.T) {
	_ = os.Setenv("TEST", "true")

	config.CreateConfigFiles()

	conf := &config.Config{
		CommonConfig: &config.CommonConfig{
			Nick:             "hjertnes",
			URL:              "https://hjertnes.social/twtxt.txt",
			DiscloseIdentity: true,
			Following: map[string]string{
				"hjertnes":    "https://hjertnes.social/twtxt.txt",
				"nonExisting": "http://example.org/feed.txt",
			},
			File: utils.ReplaceTilde(fmt.Sprintf("%s/twtxt.txt", config.GetConfigDir())),
		},
	}

	New(conf).Execute("@hjertnes test")

	config.DeleteConfigFiles()

	_ = os.Setenv("TEST", "")
}