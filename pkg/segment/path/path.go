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

func shrinkPathComponent(s string, n int) string {
	if strings.HasPrefix(s, ".") {
		return takeFirstNChars(s, n+1)
	}
	return takeFirstNChars(s, n)
}

func shrink(cwd string, home string, projectRoot string) string {
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

		isCurrent := len(pathSegments) == 0
		isProjectRoot := p == projectRoot
		if isCurrent || isProjectRoot {
			pathSegments = append(pathSegments, base)
		} else {
			pathSegments = append(pathSegments, shrinkPathComponent(base, 1))
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

	var path string
	if info.GitStatus != nil {
		path = shrink(cwd, home, info.GitStatus.TopLevel)
	} else {
		path = shrink(cwd, home, "")
	}

	return &types.Segment{
		Style:   style,
		Content: path,
	}, nil
}
