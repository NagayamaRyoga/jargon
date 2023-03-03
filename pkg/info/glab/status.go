package glab

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Status struct {
	MergeRequest MergeRequest
}

type MergeRequest struct {
	Number   int
	State    string // "open", "closed", or "merged"
	Comments int
	IsDraft  bool
}

func LoadStatus(ctx context.Context) *Status {
	cmd := exec.CommandContext(ctx, "glab", "mr", "view")
	cmd.Stdin = strings.NewReader("\n")
	output, err := cmd.Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}

	var status Status
	mr := &status.MergeRequest
loop:
	for _, line := range strings.Split(string(output), "\n") {
		line = strings.Trim(line, "\r\n")
		switch {
		case strings.HasPrefix(line, "number:"):
			if number, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(line, "number:"))); err == nil {
				mr.Number = number
			}
		case strings.HasPrefix(line, "state:"):
			mr.State = strings.TrimSpace(strings.TrimPrefix(line, "state:"))
		case strings.HasPrefix(line, "title:"):
			mr.IsDraft = strings.HasPrefix(line, "Draft:")
		case strings.HasPrefix(line, "comments:"):
			if comments, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(line, "comments:"))); err == nil {
				mr.Comments = comments
			}
		case line == "--":
			break loop
		}
	}

	return &status
}
