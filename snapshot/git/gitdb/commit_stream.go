package gitdb

import (
	"github.com/scootdev/scoot/snapshot/git/repo"
)

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
