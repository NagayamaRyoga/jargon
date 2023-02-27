package prompt

import (
	"fmt"
)

type Args struct {
	Right bool `cmd:"" help:"Prints the right prompt"`
}

func Run(args *Args) error {
	fmt.Println(args)
	return nil
}
