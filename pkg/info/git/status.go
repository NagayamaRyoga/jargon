package git

import (
	"context"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type Status struct {
	TopLevel string
	Head     *HeadStatus
	Worktree WorktreeStatus
	Upstream UpstreamStatus
	User     *UserStatus
}

func LoadStatus(ctx context.Context) *Status {
	output, err := exec.CommandContext(ctx, "git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return nil
	}

	s := &Status{
		TopLevel: strings.Trim(string(output), "\r\n"),
	}

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		s.Head = loadHead(ctx)
		wg.Done()
	}()
	go func() {
		s.Worktree, s.Upstream = loadWorktreeAndUpstream(ctx)
		wg.Done()
	}()
	go func() {
		s.User = loadUser(ctx)
		wg.Done()
	}()
	wg.Wait()

	return s
}

type HeadType int

const (
	HeadTypeBranch HeadType = iota + 1
	HeadTypeTag
	HeadTypeCommit
)

type HeadStatus struct {
	Type    HeadType
	RefName string
}

func loadHead(ctx context.Context) *HeadStatus {
	if output, err := exec.CommandContext(ctx, "git", "branch", "--show-current").Output(); err == nil {
		branchName := string(output)
		branchName = strings.Trim(branchName, "\r\n")
		if len(branchName) > 0 {
			return &HeadStatus{
				Type:    HeadTypeBranch,
				RefName: branchName,
			}
		}
	}

	if output, err := exec.CommandContext(ctx, "git", "tag", "--points-at", "HEAD").Output(); err == nil {
		tagName, _, _ := strings.Cut(string(output), "\n")
		tagName = strings.Trim(tagName, "\r\n")
		if len(tagName) > 0 {
			return &HeadStatus{
				Type:    HeadTypeTag,
				RefName: tagName,
			}
		}
	}

	if output, err := exec.CommandContext(ctx, "git", "rev-parse", "--short", "HEAD").Output(); err == nil {
		sha := string(output)
		sha = strings.Trim(sha, "\r\n")
		return &HeadStatus{
			Type:    HeadTypeCommit,
			RefName: sha,
		}
	}

	return nil
}

type WorktreeStatus int

const (
	WorktreeStatusUnstaged WorktreeStatus = 1 << iota

	WorktreeStatusAdded
	WorktreeStatusDeleted
	WorktreeStatusModified
	WorktreeStatusRenamed
	WorktreeStatusConflicted
)

func (w WorktreeStatus) Has(s WorktreeStatus) bool {
	return (w & s) != 0
}

type UpstreamStatus struct {
	Behind int
	Ahead  int
}

func loadWorktreeAndUpstream(ctx context.Context) (worktree WorktreeStatus, upstream UpstreamStatus) {
	output, err := exec.CommandContext(ctx, "git", "status", "-z", "--branch", "--porcelain").Output()
	if err != nil {
		return
	}

	lines := strings.Split(strings.TrimSuffix(string(output), "\x00"), "\x00")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if len(line) < 2 {
			continue
		}

		if line[:2] == "##" {
			matches := upstreamRegex.FindStringSubmatch(line[2:])
			if len(matches) > 0 {
				upstream.Ahead, _ = strconv.Atoi(matches[aheadSubexpIndex])
				upstream.Behind, _ = strconv.Atoi(matches[behindSubexpIndex])
			}
			continue
		}

		if line[1] != ' ' {
			worktree |= WorktreeStatusUnstaged
		}
		if line[0] == 'U' || line[1] == 'U' {
			worktree |= WorktreeStatusConflicted
		}
		if line[0] == 'A' || line[1] == 'A' || line[0] == '?' || line[1] == '?' {
			worktree |= WorktreeStatusAdded
		}
		if line[0] == 'D' || line[1] == 'D' {
			worktree |= WorktreeStatusDeleted
		}
		if line[0] == 'M' || line[1] == 'M' {
			worktree |= WorktreeStatusModified
		}
		if line[0] == 'R' || line[1] == 'R' {
			worktree |= WorktreeStatusRenamed
			i++
		}
	}

	return
}

var upstreamRegex = regexp.MustCompile(`\[(ahead\s+(?P<ahead>\d+))?(,\s*)?(behind\s+(?P<behind>\d+))?\]`)
var aheadSubexpIndex = upstreamRegex.SubexpIndex("ahead")
var behindSubexpIndex = upstreamRegex.SubexpIndex("behind")

type UserStatus struct {
	Name string
}

func loadUser(ctx context.Context) *UserStatus {
	output, err := exec.CommandContext(ctx, "git", "config", "user.name").Output()
	if err != nil {
		return nil
	}

	userName := string(output)
	userName = strings.Trim(userName, "\r\n")

	return &UserStatus{
		Name: userName,
	}
}
