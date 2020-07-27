// Package config contains everything related to reading config files
package config

import (
	"io/ioutil"
	"os"

	"git.sr.ht/~hjertnes/tw.txt/constants"
	"git.sr.ht/~hjertnes/tw.txt/utils"
	"gopkg.in/yaml.v2"
)

// CommonConfig is a shared config intended to be supported by all twtxt clients.
type CommonConfig struct {
	Nick             string
	URL              string
	File             string
	Following        map[string]string
	DiscloseIdentity bool
}

// InternalConfig Config file used by this client: located at ~/.tw.txt/config.yaml.
type InternalConfig struct {
	TwtxtLocation      string
	ConfigFileLocation string
}

// Config Type Config contains CommonConfig and InternalConfig.
type Config struct {
	InternalConfig *InternalConfig
	CommonConfig   *CommonConfig
}

func readInternalConfig() (*InternalConfig, error) {
	configFilename := utils.ReplaceTilde("~/.tw.txt/config.yaml")
	if !utils.Exist(configFilename) {
		return nil, constants.ErrConfigDoesNotExist
	}

	f, err := os.OpenFile(configFilename, os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	config := &InternalConfig{}

	err = yaml.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func readCommonConfig(filename string) (*CommonConfig, error) {
	configFilename := utils.ReplaceTilde(filename)
	if !utils.Exist(configFilename) {
		return nil, constants.ErrConfigDoesNotExist
	}

	f, err := os.OpenFile(configFilename, os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	config := &CommonConfig{}

	err = yaml.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// New builds configs.
func New() (*Config, error) {
	internal, err := readInternalConfig()
	if err != nil {
		return nil, err
	}

	common, err := readCommonConfig(internal.ConfigFileLocation)
	if err != nil {
		return nil, err
	}

	return &Config{
		InternalConfig: internal,
		CommonConfig:   common,
	}, nil
}
