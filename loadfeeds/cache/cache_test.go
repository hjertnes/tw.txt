package cache

import (
	"os"
	"testing"
	"time"

	"git.sr.ht/~hjertnes/tw.txt/config"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	_ = os.Setenv("TEST", "true")

	config.CreateConfigFiles()

	cache, _ := New()

	cache.Set("handle", "https://some-url", "content", 0, time.Now())

	d, err := cache.Get("https://some-url")

	assert.Nil(t, err)
	assert.NotNil(t, d)

	err = cache.Save()

	assert.Nil(t, err)
	config.DeleteConfigFiles()

	_ = os.Setenv("TEST", "")
}
