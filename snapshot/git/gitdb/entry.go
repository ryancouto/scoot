package gitdb

import (
	"fmt"

	"github.com/scootdev/scoot/snapshot"
	"github.com/scootdev/scoot/snapshot/git/repo"
)

func (db *DB) Init() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	return db.init()
}

func (db *DB) Download(id SnapID) (SnapID, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if err := db.init(); err != nil {
		return nil, err
	}

	return db.download(id)
}

func (db *DB) IngestGitCommit(ingestRepo *repo.Repository, commitish string) (SnapID, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if err := db.init(); err != nil {
		return nil, err
	}

	return db.ingestGitCommit(ingestRepo, commitish)
}

func (db *DB) IngestGitCommitAndUpload(ingestRepo *repo.Repository, commitish string) (SnapID, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if err := db.init(); err != nil {
		return nil, err
	}

	return db.ingestGitCommitAndUpload(ingestRepo, commitish)
}

func (db *DB) GitCheckoutInDataRepo(id SnapID) (path string, err error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if err := db.init(); err != nil {
		return "", err
	}

	return db.gitCheckoutInDataRepo(id)
}

func (db *DB) Upload(id SnapID) (SnapID, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if err := db.init(); err != nil {
		return nil, err
	}

	return db.upload(id)
}

func (db *DB) Checkout(id string) (snapshot.Checkout, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if err := db.init(); err != nil {
		return nil, err
	}

	return db.checkout(id)
}

func (db *DB) CheckoutAt(id string, dir string) (snapshot.Checkout, error) {
	return nil, fmt.Errorf("not yet implemented")
}
