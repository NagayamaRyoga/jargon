package types

import (
	"github.com/NagayamaRyoga/jargon/pkg/ansi"
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
}
