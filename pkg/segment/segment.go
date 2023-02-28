package segment

import (
	"fmt"
	"io"
	"strings"

	"github.com/NagayamaRyoga/jargon/pkg/ansi"
	"github.com/NagayamaRyoga/jargon/pkg/segment/duration"
	"github.com/NagayamaRyoga/jargon/pkg/segment/git_status"
	"github.com/NagayamaRyoga/jargon/pkg/segment/git_user"
	"github.com/NagayamaRyoga/jargon/pkg/segment/os"
	"github.com/NagayamaRyoga/jargon/pkg/segment/path"
	"github.com/NagayamaRyoga/jargon/pkg/segment/status"
	"github.com/NagayamaRyoga/jargon/pkg/segment/time"
	"github.com/NagayamaRyoga/jargon/pkg/segment/types"
	"github.com/NagayamaRyoga/jargon/pkg/segment/user"
	"github.com/mattn/go-runewidth"
)

const (
	controlStart string = "%{"
	controlEnd   string = "%}"

	leftSeparator  string = ""
	rightSeparator string = ""
)

var (
	leftSeparatorWidth  = runewidth.StringWidth(leftSeparator)
	rightSeparatorWidth = runewidth.StringWidth(rightSeparator)
)

type builder func(*types.Info) (*types.Segment, error)

var segments = map[string]builder{
	"os":         os.Build,
	"user":       user.Build,
	"path":       path.Build,
	"status":     status.Build,
	"duration":   duration.Build,
	"time":       time.Build,
	"git_status": git_status.Build,
	"git_user":   git_user.Build,
}

func buildSegment(info *types.Info, segmentName string) (*types.Segment, error) {
	if builder, ok := segments[segmentName]; ok {
		return builder(info)
	}
	return nil, fmt.Errorf("unknown segment %s", segmentName)
}

func buildSegments(info *types.Info, segmentNames []string) ([]*types.Segment, error) {
	segments := make([]*types.Segment, 0, len(segmentNames))
	for _, segName := range segmentNames {
		seg, err := buildSegment(info, segName)
		if err != nil {
			return nil, err
		}

		if seg != nil {
			segments = append(segments, seg)
		}
	}

	return segments, nil
}

func alignRight(w io.Writer, pos int) {
	fmt.Fprintf(w, "\x1b[%dG", pos)
}

func displayLeftSegments(w io.Writer, segments []*types.Segment) int {
	width := 0
	var prevBg ansi.Color
	for i, seg := range segments {
		if i > 0 {
			fmt.Fprintf(
				w,
				"%s%s%s%s%s%s%s%s",
				controlStart,
				seg.Style.Background.Background(),
				prevBg.Foreground(),
				controlEnd,
				leftSeparator,
				controlStart,
				ansi.ANSIReset,
				controlEnd,
			)

			width += leftSeparatorWidth
		}

		fmt.Fprintf(
			w,
			"%s%s%s%s %s %s%s%s",
			controlStart,
			seg.Style.Background.Background(),
			seg.Style.Foreground.Foreground(),
			controlEnd,
			seg.Content,
			controlStart,
			ansi.ANSIReset,
			controlEnd,
		)

		width += runewidth.StringWidth(seg.Content) + 2
		prevBg = seg.Style.Background
	}

	fmt.Fprintf(
		w,
		"%s%s%s%s%s%s%s",
		controlStart,
		prevBg.Foreground(),
		controlEnd,
		leftSeparator,
		controlStart,
		ansi.ANSIReset,
		controlEnd,
	)

	width += leftSeparatorWidth

	return width
}

func displayRightSegments(w io.Writer, segments []*types.Segment) int {
	width := 0
	var prevBg ansi.Color
	for i, seg := range segments {
		if i == 0 {
			fmt.Fprintf(
				w,
				"%s%s%s%s%s%s%s",
				controlStart,
				seg.Style.Background.Foreground(),
				controlEnd,
				rightSeparator,
				controlStart,
				ansi.ANSIReset,
				controlEnd,
			)

			width += rightSeparatorWidth
		} else {
			fmt.Fprintf(
				w,
				"%s%s%s%s%s%s%s%s",
				controlStart,
				prevBg.Background(),
				seg.Style.Background.Foreground(),
				controlEnd,
				rightSeparator,
				controlStart,
				ansi.ANSIReset,
				controlEnd,
			)

			width += rightSeparatorWidth
		}

		fmt.Fprintf(
			w,
			"%s%s%s%s %s %s%s%s",
			controlStart,
			seg.Style.Background.Background(),
			seg.Style.Foreground.Foreground(),
			controlEnd,
			seg.Content,
			controlStart,
			ansi.ANSIReset,
			controlEnd,
		)

		width += runewidth.StringWidth(seg.Content) + 2
		prevBg = seg.Style.Background
	}

	return width
}

func DisplayLine(w io.Writer, info *types.Info, left []string, right []string) error {
	leftSegments, err := buildSegments(info, left)
	if err != nil {
		return err
	}

	rightSegments, err := buildSegments(info, right)
	if err != nil {
		return err
	}

	var leftContent strings.Builder
	leftWidth := displayLeftSegments(&leftContent, leftSegments)

	var rightContent strings.Builder
	rightWidth := displayRightSegments(&rightContent, rightSegments)

	fmt.Fprint(w, leftContent.String())

	if len(rightSegments) > 0 && leftWidth+rightWidth+1 < info.Width {
		alignRight(w, info.Width-rightWidth)
		fmt.Fprint(w, rightContent.String())
	}

	return nil
}

func DisplayRight(w io.Writer, info *types.Info, right []string) error {
	rightSegments, err := buildSegments(info, right)
	if err != nil {
		return err
	}

	_ = displayRightSegments(w, rightSegments)

	return nil
}

func NewLine(w io.Writer) {
	fmt.Fprint(w, "\n")
}

func Finish(w io.Writer) {
	fmt.Fprint(w, " ")
}
