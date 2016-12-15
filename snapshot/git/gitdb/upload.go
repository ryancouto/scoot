package gitdb

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/scootdev/scoot/snapshot/git/repo"
)

func (db *DB) upload(id SnapID) (SnapID, error) {
	switch id := id.(type) {
	case *streamCommitSnapID:
		return id, nil
	case *bundlestoreSnapID:
		return id, nil
	case *localCommitSnapID:
		return db.uploadLocalCommit(id)
	default:
		return nil, fmt.Errorf("unexpected type %T for %v in creators.go:distribute", id, id)
	}
}

const bundleStoreTempRef = "refs/heads/scoot/bundlestore/__temp_for_writing"

func (db *DB) uploadLocalCommit(id *localCommitSnapID) (SnapID, error) {
	if err := db.snapPresent(id); err != nil {
		return nil, fmt.Errorf("cannot upload %v: sha %q not present %v", id, id.localSHA(), err)
	}

	bundleID := &bundlestoreSnapID{id.sha}

	exists, err := db.bundles.Exists(bundleID.baseName())
	if err != nil {
		return nil, err
	}
	if exists {
		return bundleID, nil
	}

	refName := bundleStoreTempRef
	_, err = db.dataRepo.Run("update-ref", refName, id.localSHA())
	if err != nil {
		return nil, err
	}

	var revList string
	if len(db.streams) > 0 {
		// TODO(dbentley): try all streams to find the smallest merge base
		stream := db.streams[0]
		mergeBaseSHA, err := db.dataRepo.RunSha("merge-base", id.localSHA(), stream.ID.branchName)
		if err != nil {
			return nil, err
		}
		log.Println("Using merge-base", mergeBaseSHA)
		revList = fmt.Sprintf("%s..%s", mergeBaseSHA, refName)
	} else {
		log.Println("Bundling without merge-base")
		revList = refName
	}

	d, err := db.tmpDir.TempDir("bundle-")
	if err != nil {
		return nil, err
	}
	bundleFilename := path.Join(d.Dir, fmt.Sprintf("commit-%s.bundle", id.localSHA()))

	if _, err := db.dataRepo.Run("bundle", "create", bundleFilename, revList); err != nil {
		return nil, err
	}

	f, err := os.Open(bundleFilename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if err := db.bundles.Write(bundleID.baseName(), f); err != nil {
		return nil, err
	}

	return bundleID, nil
}

func (db *DB) ingestGitCommitAndUpload(ingestRepo *repo.Repository, commitish string) (SnapID, error) {
	id, err := db.ingestGitCommit(ingestRepo, commitish)
	if err != nil {
		return nil, err
	}

	return db.upload(id)
}
