package cli

import (
	initialize "github.com/NagayamaRyoga/jargon/pkg/cmd/init"
	"github.com/NagayamaRyoga/jargon/pkg/cmd/prepare"
	"github.com/NagayamaRyoga/jargon/pkg/cmd/prompt"
	"github.com/alecthomas/kong"
)

type command struct {
	Init    initialize.Args `cmd:"" help:"Prints the initialization scripts"`
	Prompt  prompt.Args     `cmd:"" help:"Prints the prompt"`
	Prepare prepare.Args    `cmd:"" help:"Serialize info for lazy segments"`
}

func Run() error {
	var cmd command
	switch ctx := kong.Parse(&cmd); ctx.Command() {
	case "init":
		return initialize.Run(&cmd.Init)
	case "prompt":
		return prompt.Run(&cmd.Prompt)
	case "prepare":
		return prepare.Run(&cmd.Prepare)
	default:
		panic(ctx.Command())
	}
}
