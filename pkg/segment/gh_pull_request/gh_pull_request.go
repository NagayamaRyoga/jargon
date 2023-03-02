package gh_pull_request

import (
	"fmt"

	"github.com/NagayamaRyoga/jargon/pkg/ansi"
	"github.com/NagayamaRyoga/jargon/pkg/segment/types"
)

const (
	ghIcon      string = ""
	draftIcon   string = ""
	closedIcon  string = ""
	mergedIcon  string = ""
	commentIcon string = " "
)

var (
	openStyle = types.Style{
		Foreground: ansi.ColorANSIBlack,
		Background: ansi.Color256(214),
	}
	draftStyle = types.Style{
		Foreground: ansi.ColorANSIBlack,
		Background: ansi.Color256(249),
	}
	closedStyle = types.Style{
		Foreground: ansi.ColorANSIBlack,
		Background: ansi.Color256(196),
	}
	mergedStyle = types.Style{
		Foreground: ansi.ColorANSIBlack,
		Background: ansi.Color256(141),
	}
)

func Build(info *types.Info) (*types.Segment, error) {
	if info.GhStatus == nil {
		return nil, nil
	}

	pr := &info.GhStatus.PullRequest

	var content string

	content += fmt.Sprintf("%s #%d", ghIcon, pr.Number)

	var style types.Style
	switch {
	case pr.IsDraft:
		content += fmt.Sprintf(" %s", draftIcon)
		style = draftStyle
	case pr.State == "CLOSED":
		content += fmt.Sprintf(" %s", closedIcon)
		style = closedStyle
	case pr.State == "MERGED":
		content += fmt.Sprintf(" %s", mergedIcon)
		style = mergedStyle
	default:
		content += ""
		style = openStyle
	}

	if pr.Comments > 0 {
		content += fmt.Sprintf(" %s %d", commentIcon, pr.Comments)
	}

	return &types.Segment{
		Style:   style,
		Content: content,
	}, nil
}
