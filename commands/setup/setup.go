// Package setup sets up config files.
package setup

import "C"
import (
	"fmt"
	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/utils"
	"gopkg.in/yaml.v2"
	"os"
)

// Command is the exposed interface.
type Command interface{
	Execute()
}

type command struct {
}

func (c *command) Execute(){
	path := utils.ReplaceTilde(config.GetConfigDir())
	filename := utils.ReplaceTilde(config.GetConfigFilename())

	if !utils.Exist(path){
		err := os.MkdirAll(path, 0755)
		utils.ErrorHandler(err)
	}

	if !utils.Exist(filename){
		f, err := os.Create(filename)
		utils.ErrorHandler(err)

		content, err := yaml.Marshal(&config.InternalConfig{

		})
		utils.ErrorHandler(err)

		_, err = f.Write(content)
		utils.ErrorHandler(err)

		err = f.Close()
		utils.ErrorHandler(err)
	}

	fmt.Println("More about twtxt https://twtxt.readthedocs.org/en/stable/")
	fmt.Println("Information about the config file:")
	fmt.Println("ConfigFileLocation: location of your common twtxt yaml config file")
	fmt.Println("I keep mine next to my twtxt file")
	fmt.Println("Sample: https://git.sr.ht/~hjertnes/tw.txt/tree/master/config.yaml.sample")
}

// New creates new Command.
func New() Command {
	return &command{}
}
