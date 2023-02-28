package time

import (
	"fmt"
	"time"

	"github.com/NagayamaRyoga/jargon/pkg/ansi"
	"github.com/NagayamaRyoga/jargon/pkg/segment/types"
)

const (
	icon string = ""
)

var (
	style = types.Style{
		Foreground: ansi.ColorANSIWhite,
		Background: ansi.Color256(8),
	}
)

func Build(info *types.Info) (*types.Segment, error) {
	now := time.Now()

	return &types.Segment{
		Style:   style,
		Content: fmt.Sprintf("%s %s", icon, now.Format("2006/01/02 15:04:05")),
	}, nil
}
