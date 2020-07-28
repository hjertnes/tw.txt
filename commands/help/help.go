package help

import "fmt"

type Command interface {
	Execute()
}

type command struct {

}

func (c *command) Execute(){
	fmt.Println("tw.txt is another twtxt client -- https://twtxt.readthedocs.org/en/stable")
	fmt.Println("Usage:")
	fmt.Println("\t tw.txt command")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("\ttimeline - prints last 1000 items in your timeline")
	fmt.Println("\t\t full - prints your entire timeline")
	fmt.Println("\tedit - opens your twtxt in $EDITOR")
	fmt.Println("\t\t twtxt - opens your twtxt in $EDITOR")
	fmt.Println("\t\t common-config - opens common-config in EDITO")
	fmt.Println("\t\t internal-config - opens tw.txt config in $EDITOR")
	fmt.Println("\tsetup - creates config file and opens it in $EDITOR")
}

func New() Command{
	return &command{}
}
