package gitdb

import (
	"github.com/scootdev/scoot/snapshot/git/repo"
)

// A CommitStream is what Scoot calls a Branch in Source.
// The idea is that, unlike a Snapshot, this Stream flows forward and we may want to follow it.

type CommitStreamID struct {
	name       string
	branchName string
}

type CommitStreamIniter interface {
	InitStream(*repo.Repository) error
}

type CommitStream struct {
	ID     CommitStreamID
	initer CommitStreamIniter
}

func NewCommitStream(name string, branchName string, initer CommitStreamIniter) *CommitStream {
	return &CommitStream{
		ID: CommitStreamID{
			name:       name,
			branchName: branchName,
		},
		initer: initer,
	}
}
