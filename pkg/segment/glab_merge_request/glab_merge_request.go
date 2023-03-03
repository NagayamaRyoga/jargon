package glab_merge_request

import (
	"fmt"

	"github.com/NagayamaRyoga/jargon/pkg/ansi"
	"github.com/NagayamaRyoga/jargon/pkg/segment/types"
)

const (
	glabIcon    string = ""
	draftIcon   string = ""
	closedIcon  string = ""
	mergedIcon  string = ""
	commentIcon string = ""
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
	if info.GlabStatus == nil {
		return nil, nil
	}

	mr := &info.GlabStatus.MergeRequest

	var content string

	content += fmt.Sprintf("%s !%d", glabIcon, mr.Number)

	var style types.Style
	switch {
	case mr.IsDraft:
		content += fmt.Sprintf(" %s", draftIcon)
		style = draftStyle
	case mr.State == "closed":
		content += fmt.Sprintf(" %s", closedIcon)
		style = closedStyle
	case mr.State == "merged":
		content += fmt.Sprintf(" %s", mergedIcon)
		style = mergedStyle
	default:
		content += ""
		style = openStyle
	}

	if mr.Comments > 0 {
		content += fmt.Sprintf(" %s %d", commentIcon, mr.Comments)
	}

	return &types.Segment{
		Style:   style,
		Content: content,
	}, nil
}
