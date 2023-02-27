package user

import (
	"fmt"
	"os"
	osuser "os/user"

	"github.com/NagayamaRyoga/jargon/pkg/ansi"
	"github.com/NagayamaRyoga/jargon/pkg/segment/types"
)

var (
	style = types.Style{
		Foreground: ansi.ColorANSIWhite,
		Background: ansi.Color256(8),
	}
)

func username() string {
	user, err := osuser.Current()
	if err != nil {
		return "?"
	}
	return user.Username
}

func hostname() string {
	host, err := os.Hostname()
	if err != nil {
		return "?"
	}
	return host
}

func Build(info *types.Info) (*types.Segment, error) {
	user := username()
	host := hostname()

	return &types.Segment{
		Style:   style,
		Content: fmt.Sprintf("%s@%s", user, host),
	}, nil
}
