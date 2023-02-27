package initialize

import (
	_ "embed"
	"fmt"
)

type Args struct {
}

//go:embed init.zsh
var initZsh string

func Run(*Args) error {
	fmt.Print(initZsh)
	return nil
}
