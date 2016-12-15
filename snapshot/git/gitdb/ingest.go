package gitdb

import (
	"fmt"
	"log"

	"github.com/scootdev/scoot/snapshot/git/repo"
)

func (db *DB) ingestGitCommit(ingestRepo *repo.Repository, commitish string) (SnapID, error) {
	if commitish == "" {
		return nil, fmt.Errorf("must specify a commitish to ingest")
	}
	sha, err := ingestRepo.RunSha("rev-parse", "--verify", fmt.Sprintf("%s^{commit}", commitish))
	if err != nil {
		return nil, fmt.Errorf("not a valid commit: %s, %v", commitish, err)
	}

	log.Printf("Trying to ingest git commit %q->%q from %q", commitish, sha, ingestRepo.Dir())

	if ingestRepo.Dir() != db.dataRepo.Dir() {
		return nil, fmt.Errorf("cross-repo ingesting not yet implemented: %s %s", ingestRepo.Dir(), db.dataRepo.Dir())
	}

	return &localCommitSnapID{sha}, nil
}
