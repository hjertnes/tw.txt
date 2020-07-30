package config

import (
	"git.sr.ht/~hjertnes/tw.txt/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetConfigDir(t *testing.T) {
	assert.Equal(t, "~/.tw.txt", GetConfigDir())
	_ = os.Setenv("TEST", "true")
	assert.Equal(t, "~/.tw.txt-test", GetConfigDir())
	_ = os.Setenv("TEST", "")
}

func TestGetConfigFile(t *testing.T) {
	assert.Equal(t, "~/.tw.txt/config.yaml", GetConfigFilename())
	_ = os.Setenv("TEST", "true")
	assert.Equal(t, "~/.tw.txt-test/config.yaml", GetConfigFilename())
	_ = os.Setenv("TEST", "")
}

func TestNew(t *testing.T){
	c, _ := New()

	Save(c)

}

func TestCreateConfigFiles(t *testing.T) {
	_ = os.Setenv("TEST", "true")
	assert.False(t, utils.Exist(GetConfigDir()))
	CreateConfigFiles()
	assert.True(t, utils.Exist(GetConfigDir()))
	DeleteConfigFiles()
	assert.False(t, utils.Exist(GetConfigDir()))
	_ = os.Setenv("TEST", "")
}