package cli

import (
	"github.com/alecthomas/kong"
)

type command struct {
	Init    struct{} `cmd:"" help:"Prints the initialization scripts"`
	Prompt  struct{} `cmd:"" help:"Prints the prompt"`
	Prepare struct{} `cmd:"" help:"Serialize info for lazy segments"`
}

func Run() {
	var cmd command
	switch ctx := kong.Parse(&cmd); ctx.Command() {
	default:
		panic(ctx.Command())
	}
}
