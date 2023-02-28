package prompt

import (
	"os"

	"github.com/NagayamaRyoga/jargon/pkg/info"
	"github.com/NagayamaRyoga/jargon/pkg/info/git"
	"github.com/NagayamaRyoga/jargon/pkg/segment"
	"github.com/NagayamaRyoga/jargon/pkg/segment/types"
)

type Args struct {
	Right      bool    `help:"Prints the right prompt"`
	ExitStatus int     `help:"The status code of the last command"`
	Duration   float64 `help:"The command duration of the last command"`
	Jobs       int     `help:"The number of currently running jobs"`
	Width      int     `help:"Column width"`
	DataGit    string  `help:"Output of 'jargon prepare --source=git'"`
}

type line struct {
	left  []string
	right []string
}

func Run(args *Args) error {
	gitStatus, err := info.Decode[git.Status](args.DataGit)
	if err != nil {
		return err
	}

	info := &types.Info{
		ExitStatus: args.ExitStatus,
		Duration:   args.Duration,
		Jobs:       args.Jobs,
		Width:      args.Width,
		GitStatus:  gitStatus,
	}

	segmentLines := []line{
		{
			left: []string{
				"os",
				"user",
				"path",
				"git_status",
				"git_user",
			},
			right: []string{
				"time",
			},
		},
		{
			left: []string{
				"status",
			},
			right: []string{
				"duration",
			},
		},
	}

	w := os.Stdout

	if !args.Right {
		for i, line := range segmentLines {
			if i > 0 {
				segment.NewLine(w)
			}

			var left, right []string
			left = line.left
			if i != len(segmentLines)-1 {
				right = line.right
			}

			if err := segment.DisplayLine(w, info, left, right); err != nil {
				return err
			}
		}

		segment.Finish(w)
	} else {
		right := segmentLines[len(segmentLines)-1].right

		if err := segment.DisplayRight(w, info, right); err != nil {
			return err
		}
	}

	return nil
}
