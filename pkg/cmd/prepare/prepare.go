package prepare

import (
	"fmt"
)

type Args struct {
}

func Run(args *Args) error {
	fmt.Println(args)
	return nil
}
