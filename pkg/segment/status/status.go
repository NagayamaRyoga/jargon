package status

import (
	"fmt"

	"github.com/NagayamaRyoga/jargon/pkg/ansi"
	"github.com/NagayamaRyoga/jargon/pkg/segment/types"
)

const (
	successIcon string = "✓"
	errorIcon   string = ""
	jobsIcon    string = ""
)

var (
	successStyle = types.Style{
		Foreground: ansi.ColorANSIGreen,
		Background: ansi.ColorANSIWhite,
	}
	errorStyle = types.Style{
		Foreground: ansi.ColorANSIWhite,
		Background: ansi.ColorANSIRed,
	}
)

func Build(info *types.Info) (*types.Segment, error) {
	content := ""
	if info.ExitStatus == 0 {
		content += successIcon
	} else {
		content += fmt.Sprintf("%s %d", errorIcon, info.ExitStatus)
	}

	if info.Jobs > 0 {
		content += " "
		content += jobsIcon
	}

	var style types.Style
	if info.ExitStatus == 0 {
		style = successStyle
	} else {
		style = errorStyle
	}

	return &types.Segment{
		Style:   style,
		Content: content,
	}, nil
}
