//go:generate curl -sSL -o async.zsh https://raw.githubusercontent.com/mafredri/zsh-async/main/async.zsh
package initialize

import (
	_ "embed"
	"fmt"
)

type Args struct {
}

//go:embed async.zsh
var asyncZsh string

//go:embed init.zsh
var initZsh string

func Run(*Args) error {
	fmt.Println(asyncZsh)
	fmt.Print(initZsh)
	return nil
}
