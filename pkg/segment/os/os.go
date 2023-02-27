package os

import (
	"runtime"

	"github.com/NagayamaRyoga/jargon/pkg/ansi"
	"github.com/NagayamaRyoga/jargon/pkg/segment/types"
)

const (
	linuxIcon   string = ""
	macIcon     string = ""
	windowsIcon string = ""
)

var (
	style = types.Style{
		Foreground: ansi.ColorANSIWhite,
		Background: ansi.Color256(33),
	}
)

func Build(info *types.Info) (*types.Segment, error) {
	var icon string
	switch runtime.GOOS {
	case "linux":
		icon = linuxIcon
	case "darwin":
		icon = macIcon
	case "windows":
		icon = windowsIcon
	default:
		icon = "?"
	}

	return &types.Segment{
		Style:   style,
		Content: icon,
	}, nil
}
