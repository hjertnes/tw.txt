// Package setup sets up config files.
package setup

import "C"
import (
	"fmt"
	"git.sr.ht/~hjertnes/tw.txt/config"
	"git.sr.ht/~hjertnes/tw.txt/utils"
	"gopkg.in/yaml.v2"
	"os"
	"os/exec"
)

// Command is the exposed interface.
type Command interface{
	Execute()
}

type command struct {
}

func (c *command) Execute(){
	path := utils.ReplaceTilde("~/.tw.txt")
	filename := fmt.Sprintf("%s/config.yaml", path)

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
		fmt.Println("Information about the config file:")
		fmt.Println("TwtxtLocation: where you twtxt file is")
		fmt.Println("More about twtxt https://twtxt.readthedocs.org/en/stable/")
		fmt.Println("ConfigFileLocation: location of your common twtxt yaml config file")
		fmt.Println("I keep mine next to my twtxt file")
		fmt.Println("Sample: https://git.sr.ht/~hjertnes/tw.txt/tree/master/config.yaml.sample")
		cmd := exec.Command(os.Getenv("EDITOR"))
		cmd.Args = append(cmd.Args, filename)
		err = cmd.Run()
		utils.ErrorHandler(err)
	}
}

// New creates new Command.
func New() Command {
	return &command{}
}
