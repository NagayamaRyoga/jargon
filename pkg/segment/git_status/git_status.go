package git_status

import (
	"fmt"

	"github.com/NagayamaRyoga/jargon/pkg/ansi"
	"github.com/NagayamaRyoga/jargon/pkg/git"
	"github.com/NagayamaRyoga/jargon/pkg/segment/types"
)

const (
	branchIcon string = ""
	tagIcon    string = ""
	commitIcon string = ""

	addedIcon      string = "+"
	deletedIcon    string = "-"
	modifiedIcon   string = "…"
	renamedIcon    string = "→"
	conflictedIcon string = ""

	behindIcon string = ""
	aheadIcon  string = ""
)

var (
	cleanStyle = types.Style{
		Foreground: ansi.ColorANSIBlack,
		Background: ansi.ColorANSIGreen,
	}
	dirtyStyle = types.Style{
		Foreground: ansi.ColorANSIBlack,
		Background: ansi.ColorANSIYellow,
	}
	conflictStyle = types.Style{
		Foreground: ansi.ColorANSIBlack,
		Background: ansi.ColorANSIRed,
	}
)

func Build(info *types.Info) (*types.Segment, error) {
	if info.GitStatus == nil {
		return nil, nil
	}

	style := &cleanStyle
	var content string

	if head := info.GitStatus.Head; head != nil {
		var icon string
		switch head.Type {
		case git.HeadTypeBranch:
			icon = branchIcon
		case git.HeadTypeTag:
			icon = tagIcon
		case git.HeadTypeCommit:
			icon = commitIcon
		default:
			icon = "?"
		}
		content += fmt.Sprintf("%s %s", icon, head.RefName)
	}

	if worktree := info.GitStatus.Worktree; worktree != 0 {
		if worktree.Has(git.WorktreeStatusConflicted) {
			style = &conflictStyle
		} else if worktree.Has(git.WorktreeStatusUnstaged) {
			style = &dirtyStyle
		}

		content += " "
		if worktree.Has(git.WorktreeStatusAdded) {
			content += addedIcon
		}
		if worktree.Has(git.WorktreeStatusDeleted) {
			content += deletedIcon
		}
		if worktree.Has(git.WorktreeStatusModified) {
			content += modifiedIcon
		}
		if worktree.Has(git.WorktreeStatusRenamed) {
			content += renamedIcon
		}
		if worktree.Has(git.WorktreeStatusConflicted) {
			content += conflictedIcon
		}
	}

	if upstream := info.GitStatus.Upstream; upstream.Ahead > 0 || upstream.Behind > 0 {
		content += " "
		if upstream.Behind > 0 {
			content += fmt.Sprintf("%s%d", behindIcon, upstream.Behind)
		}
		if upstream.Ahead > 0 {
			content += fmt.Sprintf("%s%d", aheadIcon, upstream.Ahead)
		}
	}

	return &types.Segment{
		Style:   *style,
		Content: content,
	}, nil
}
