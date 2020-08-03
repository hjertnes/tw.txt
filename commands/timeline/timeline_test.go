package timeline

import (
	"testing"

	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/loadfeeds/headfeeds"
)

func TestTest(t *testing.T) {
	conf := &config.Config{
		CommonConfig: &config.CommonConfig{
			Nick:             "hjertnes",
			URL:              "https://hjertnes.social/twtxt.txt",
			DiscloseIdentity: true,
			Following: map[string]string{
				"hjertnes": "https://hjertnes.social/twtxt.txt",
			},
		},
	}
	ff := headfeeds.New(conf)

	New(conf, ff).Execute("")
	New(conf, ff).Execute("full")
}