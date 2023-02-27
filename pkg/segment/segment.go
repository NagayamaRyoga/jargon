package segment

import (
	"fmt"

	"github.com/NagayamaRyoga/jargon/pkg/segment/duration"
	"github.com/NagayamaRyoga/jargon/pkg/segment/os"
	"github.com/NagayamaRyoga/jargon/pkg/segment/path"
	"github.com/NagayamaRyoga/jargon/pkg/segment/status"
	"github.com/NagayamaRyoga/jargon/pkg/segment/types"
	"github.com/NagayamaRyoga/jargon/pkg/segment/user"
)

type builder func(*types.Info) (*types.Segment, error)

var segments = map[string]builder{
	"os":       os.Build,
	"user":     user.Build,
	"path":     path.Build,
	"status":   status.Build,
	"duration": duration.Build,
}

func buildSegment(info *types.Info, segmentName string) (*types.Segment, error) {
	if builder, ok := segments[segmentName]; ok {
		return builder(info)
	}
	return nil, fmt.Errorf("unknown segment %s", segmentName)
}

func DisplaySegments(info *types.Info, segmentNames []string) error {
	for _, name := range segmentNames {
		seg, err := buildSegment(info, name)
		if err != nil {
			return err
		}

		if seg == nil {
			continue
		}

		fmt.Printf("[ %s ]", seg.Content)
	}
	fmt.Print(" ")

	return nil
}
