package gh

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type Status struct {
	PullRequest PullRequest
}

type PullRequest struct {
	Number   int
	State    string // "OPEN", "CLOSED", or "MERGED"
	Comments int
	IsDraft  bool
}

func LoadStatus(ctx context.Context) *Status {
	output, err := exec.CommandContext(ctx, "gh", "pr", "view", "--json=number,state,comments,reviews,isDraft").Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}

	type prResponse struct {
		Number   int        `json:"number"`
		State    string     `json:"state"`
		Comments []struct{} `json:"comments"`
		Reviews  []struct{} `json:"reviews"`
		IsDraft  bool       `json:"isDraft"`
	}

	var pr prResponse
	if err := json.Unmarshal(output, &pr); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}

	return &Status{
		PullRequest: PullRequest{
			Number:   pr.Number,
			State:    pr.State,
			Comments: len(pr.Comments) + len(pr.Reviews),
			IsDraft:  pr.IsDraft,
		},
	}
}
