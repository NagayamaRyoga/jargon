package main

import (
	"github.com/NagayamaRyoga/jargon/pkg/cli"
)

func main() {
	if err := cli.Run(); err != nil {
		panic(err)
	}
}
