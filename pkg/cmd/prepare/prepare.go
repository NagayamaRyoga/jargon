package prepare

import (
	"context"
	"fmt"

	"github.com/NagayamaRyoga/jargon/pkg/info"
	"github.com/NagayamaRyoga/jargon/pkg/info/gh"
	"github.com/NagayamaRyoga/jargon/pkg/info/git"
)

type Args struct {
	Source string `help:"git, gh, or glab"`
}

func prepareGit(ctx context.Context) error {
	status := git.LoadStatus(ctx)
	if status == nil {
		return nil
	}

	encoded, err := info.Encode(status)
	if err != nil {
		return err
	}

	fmt.Print(encoded)
	return nil
}

func prepareGh(ctx context.Context) error {
	status := gh.LoadStatus(ctx)
	if status == nil {
		return nil
	}

	encoded, err := info.Encode(status)
	if err != nil {
		return err
	}

	fmt.Print(encoded)
	return nil
}

func Run(args *Args) error {
	ctx := context.Background()

	switch args.Source {
	case "git":
		return prepareGit(ctx)
	case "gh":
		return prepareGh(ctx)
	case "glab":
		panic(0)
	default:
		panic(args.Source)
	}
}
