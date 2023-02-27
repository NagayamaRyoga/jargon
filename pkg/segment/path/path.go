package path

import (
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/NagayamaRyoga/jargon/pkg/ansi"
	"github.com/NagayamaRyoga/jargon/pkg/segment/types"
)

const (
	homeIcon      string = "~"
	pathSeparator string = "/"
)

var (
	style = types.Style{
		Foreground: ansi.ColorANSIBlack,
		Background: ansi.ColorANSIBlue,
	}
)

func takeFirstNChars(s string, n int) string {
	runes := []rune(s)
	if len(runes) <= n {
		return s
	}
	return string(runes[:n])
}

func shrink(cwd string, home string) string {
	pathSegments := []string{}

	p := cwd
	for {
		if p == home {
			pathSegments = append(pathSegments, homeIcon)
			break
		}

		parent := path.Dir(p)
		if parent == p {
			if runtime.GOOS == "windows" || len(pathSegments) == 0 {
				pathSegments = append(pathSegments, parent)
			} else {
				pathSegments = append(pathSegments, "")
			}
			break
		}

		base := path.Base(p)
		if len(pathSegments) == 0 {
			pathSegments = append(pathSegments, base)
		} else if strings.HasPrefix(base, ".") {
			pathSegments = append(pathSegments, takeFirstNChars(base, 2))
		} else {
			pathSegments = append(pathSegments, takeFirstNChars(base, 1))
		}

		p = parent
	}

	var result string
	for i := len(pathSegments) - 1; i >= 0; i-- {
		result += pathSegments[i]
		if i > 0 {
			result += pathSeparator
		}
	}

	return result
}

func Build(info *types.Info) (*types.Segment, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		home = ""
	}

	path := shrink(cwd, home)

	return &types.Segment{
		Style:   style,
		Content: path,
	}, nil
}
