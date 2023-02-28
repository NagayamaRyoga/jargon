package git_user

import (
	"fmt"

	"github.com/NagayamaRyoga/jargon/pkg/ansi"
	"github.com/NagayamaRyoga/jargon/pkg/segment/types"
)

const (
	icon string = ""
)

var (
	style = types.Style{
		Foreground: ansi.ColorANSIBlack,
		Background: ansi.Color256(117),
	}
)

func Build(info *types.Info) (*types.Segment, error) {
	if info.GitStatus == nil || info.GitStatus.User == nil || len(info.GitStatus.User.Name) == 0 {
		return nil, nil
	}

	userName := info.GitStatus.User.Name

	return &types.Segment{
		Style:   style,
		Content: fmt.Sprintf("%s %s", icon, userName),
	}, nil
}
