package main

import (
	"github.com/NagayamaRyoga/jargon/pkg/cli"
)

var (
	version = "dev"
	commit  = "HEAD"
	date    = "unknown"
)

func main() {
	if err := cli.Run(); err != nil {
		panic(err)
	}
}
