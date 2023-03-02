package types

import (
	"github.com/NagayamaRyoga/jargon/pkg/ansi"
	"github.com/NagayamaRyoga/jargon/pkg/info/gh"
	"github.com/NagayamaRyoga/jargon/pkg/info/git"
)

type Style struct {
	Foreground ansi.Color
	Background ansi.Color
}

type Segment struct {
	Style   Style
	Content string
}

type Info struct {
	ExitStatus int
	Duration   float64
	Jobs       int
	Width      int
	GitStatus  *git.Status
	GhStatus   *gh.Status
}
