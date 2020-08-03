// Package config contains everything related to reading config files
package config

import (
	"fmt"
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

// InternalConfig config file used by this client: located at ~/.tw.txt/config.yaml.
type InternalConfig struct {
	ConfigFileLocation string
	TemplateFileLocation string
}

// Config Type config contains CommonConfig and InternalConfig.
type Config struct {
	InternalConfig *InternalConfig
	CommonConfig   *CommonConfig
}

// GetConfigDir Get Config Dir.
func GetConfigDir() string {
	if os.Getenv("TEST") != "" {
		return "~/.tw.txt-test"
	}

	return "~/.tw.txt"
}

// GetConfigFilename Get Config Filename.
func GetConfigFilename() string {
	if os.Getenv("TEST") != "" {
		return "~/.tw.txt-test/config.yaml"
	}

	return "~/.tw.txt/config.yaml"
}

func writeInternalConfig(conf *InternalConfig) error {
	configFilename := utils.ReplaceTilde(GetConfigFilename())
	if !utils.Exist(configFilename) {
		return constants.ErrConfigDoesNotExist
	}

	content, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configFilename, content, 0)
	if err != nil {
		return err
	}

	return nil
}

func writeCommonConfig(conf *Config) error {
	configFilename := utils.ReplaceTilde(conf.InternalConfig.ConfigFileLocation)
	if !utils.Exist(configFilename) {
		return constants.ErrConfigDoesNotExist
	}

	content, err := yaml.Marshal(conf.CommonConfig)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configFilename, content, 0)
	if err != nil {
		return err
	}

	return nil
}

func readInternalConfig() (*InternalConfig, error) {
	configFilename := utils.ReplaceTilde(GetConfigFilename())
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

	f, err := os.OpenFile(configFilename, os.O_RDONLY, 0600)
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

// Save Write back config files.
func Save(conf *Config) {
	err := writeInternalConfig(conf.InternalConfig)
	utils.ErrorHandler(err)

	err = writeCommonConfig(conf)
	utils.ErrorHandler(err)
}

// CreateConfigFiles Creates config files for tests.
func CreateConfigFiles() {
	path := utils.ReplaceTilde(GetConfigDir())
	filename := utils.ReplaceTilde(GetConfigFilename())
	filename2 := fmt.Sprintf("%s/config2.yaml", path)
	filename3 := fmt.Sprintf("%s/twtxt.txt", path)

	err := os.MkdirAll(path, 0700)
	utils.ErrorHandler(err)

	f, err := os.Create(filename)
	utils.ErrorHandler(err)

	content, err := yaml.Marshal(&InternalConfig{
		ConfigFileLocation: filename2,
	})
	utils.ErrorHandler(err)

	_, err = f.Write(content)
	utils.ErrorHandler(err)

	err = f.Close()
	utils.ErrorHandler(err)

	f, err = os.Create(filename2)
	utils.ErrorHandler(err)

	content2, err := yaml.Marshal(&CommonConfig{
		File:      utils.ReplaceTilde(fmt.Sprintf("%s/twtxt.txt", GetConfigDir())),
		Following: map[string]string{"hjertnes": "https://hjertnes.social/twtxt.txt"},
	})
	utils.ErrorHandler(err)

	_, err = f.Write(content2)
	utils.ErrorHandler(err)

	err = f.Close()
	utils.ErrorHandler(err)

	f, err = os.Create(filename3)
	utils.ErrorHandler(err)

	_, err = f.Write([]byte(""))
	utils.ErrorHandler(err)

	err = f.Close()
	utils.ErrorHandler(err)
}

// DeleteConfigFiles deletes config files for tests.
func DeleteConfigFiles() {
	_ = os.RemoveAll(utils.ReplaceTilde(GetConfigDir()))
}