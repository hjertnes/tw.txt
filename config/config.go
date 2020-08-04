// Package config contains everything related to reading config files
package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"git.sr.ht/~hjertnes/tw.txt/models"

	"git.sr.ht/~hjertnes/tw.txt/constants"
	"git.sr.ht/~hjertnes/tw.txt/utils"
	"gopkg.in/yaml.v2"
)

type Service interface {
	Get() *models.Config
	Save() error
}

type service struct {
	config *models.Config
}

func (s *service) Get() *models.Config{
	return s.config
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

func writeInternalConfig(conf *models.InternalConfig) error {
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

func writeCommonConfig(conf *models.Config) error {
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

func readInternalConfig() (*models.InternalConfig, error) {
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

	config := &models.InternalConfig{}

	err = yaml.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func readCommonConfig(filename string) (*models.CommonConfig, error) {
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

	config := &models.CommonConfig{}

	err = yaml.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// New builds configs.
func New() (Service, error) {
	internal, err := readInternalConfig()
	if err != nil {
		return nil, err
	}

	common, err := readCommonConfig(internal.ConfigFileLocation)
	if err != nil {
		return nil, err
	}

	return &service{
		config: &models.Config{
			InternalConfig: internal,
			CommonConfig:   common,
		},
	}, nil
}

// Save Write back config files.
func (s *service) Save() error {
	err := writeInternalConfig(s.config.InternalConfig)
	if err != nil{
		return err
	}

	err = writeCommonConfig(s.config)

	return err
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

	content, err := yaml.Marshal(&models.InternalConfig{
		ConfigFileLocation: filename2,
	})
	utils.ErrorHandler(err)

	_, err = f.Write(content)
	utils.ErrorHandler(err)

	err = f.Close()
	utils.ErrorHandler(err)

	f, err = os.Create(filename2)
	utils.ErrorHandler(err)

	content2, err := yaml.Marshal(&models.CommonConfig{
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