package following

import (
	"testing"

	"git.sr.ht/~hjertnes/tw.txt/mocks"
	"git.sr.ht/~hjertnes/tw.txt/models"
)

func TestTest(t *testing.T) {
	conf := &models.Config{
		CommonConfig: &models.CommonConfig{
			Following: map[string]string{},
		},
		InternalConfig: &models.InternalConfig{},
	}
	conf.CommonConfig.Following["hjertnes"] = "https://hjertnes.social"

	c := &mocks.ConfigMock{}

	c.On("Get").Return(conf)

	New(c).Execute()
}
