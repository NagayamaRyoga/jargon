package prepare

import (
	"context"
	"fmt"

	"github.com/NagayamaRyoga/jargon/pkg/info"
	"github.com/NagayamaRyoga/jargon/pkg/info/gh"
	"github.com/NagayamaRyoga/jargon/pkg/info/git"
	"github.com/NagayamaRyoga/jargon/pkg/info/glab"
)

type Args struct {
	Source string `help:"git, gh, or glab"`
}

func prepare[T any](ctx context.Context, loadStatus func(context.Context) *T) error {
	status := loadStatus(ctx)
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
		return prepare(ctx, git.LoadStatus)
	case "gh":
		return prepare(ctx, gh.LoadStatus)
	case "glab":
		return prepare(ctx, glab.LoadStatus)
	default:
		panic(args.Source)
	}
}
