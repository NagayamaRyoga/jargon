package duration

import (
	"fmt"
	"time"

	"github.com/NagayamaRyoga/jargon/pkg/ansi"
	"github.com/NagayamaRyoga/jargon/pkg/segment/types"
	"github.com/hako/durafmt"
)

const (
	icon string = "祥"
)

var (
	style = types.Style{
		Foreground: ansi.ColorANSIWhite,
		Background: ansi.Color256(242),
	}
)

var (
	units, _ = durafmt.DefaultUnitsCoder.Decode("y:y,w:w,d:d,h:h,m:m,s:s,ms:ms,μs:μs")
)

func Build(info *types.Info) (*types.Segment, error) {
	if info.Duration <= 0 {
		return nil, nil
	}

	duration := durafmt.Parse(time.Duration(info.Duration * float64(time.Second))).LimitFirstN(2)

	return &types.Segment{
		Style:   style,
		Content: fmt.Sprintf("%s %s", icon, duration.Format(units)),
	}, nil
}
