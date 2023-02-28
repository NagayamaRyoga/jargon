package prompt

import (
	"context"

	"github.com/NagayamaRyoga/jargon/pkg/git"
	"github.com/NagayamaRyoga/jargon/pkg/segment"
	"github.com/NagayamaRyoga/jargon/pkg/segment/types"
)

type Args struct {
	Right      bool    `help:"Prints the right prompt"`
	ExitStatus int     `help:"The status code of the last command"`
	Duration   float64 `help:"The command duration of the last command"`
	Jobs       int     `help:"The number of currently running jobs"`
}

func Run(args *Args) error {
	ctx := context.Background()

	info := &types.Info{
		ExitStatus: args.ExitStatus,
		Duration:   args.Duration,
		Jobs:       args.Jobs,
		GitStatus:  git.LoadStatus(ctx),
	}

	segments := []string{
		"os",
		"user",
		"path",
		"git_status",
		"git_user",
		"status",
		"duration",
		"time",
	}

	if err := segment.DisplaySegments(info, segments); err != nil {
		return err
	}

	return nil
}
