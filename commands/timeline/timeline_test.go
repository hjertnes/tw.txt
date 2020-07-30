package timeline

import (
	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/services/fetchfeeds"
	"testing"
)

func TestTest(t *testing.T){
	conf := &config.Config{
		CommonConfig: &config.CommonConfig{
			Nick:             "hjertnes",
			URL:              "https://hjertnes.social/twtxt.txt",
			DiscloseIdentity: true,
			Following: map[string]string{
				"hjertnes":    "https://hjertnes.social/twtxt.txt",
			},
		},
	}
	ff := fetchfeeds.New(conf)

	New(conf, ff).Execute("")
	New(conf, ff).Execute("full")

}