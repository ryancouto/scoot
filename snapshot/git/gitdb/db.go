package gitdb

import (
	"log"
	"sync"

	"github.com/scootdev/scoot/os/temp"
	"github.com/scootdev/scoot/snapshot/bundlestore"
	"github.com/scootdev/scoot/snapshot/git/repo"
)

func MakeDB(dataRepo *repo.Repository, streams []*CommitStream, tmpDir *temp.TempDir, bundles bundlestore.Store) *DB {
	return &DB{
		dataRepo: dataRepo,
		streams:  streams,
		tmpDir:   tmpDir,
		bundles:  bundles,
	}
}

type DB struct {
	dataRepo *repo.Repository
	streams  []*CommitStream
	tmpDir   *temp.TempDir
	bundles  bundlestore.Store
	mu       sync.Mutex
	inited   bool
	initErr  error
}

func (db *DB) init() error {
	if !db.inited {
		log.Println("snapshot/git/gitdb/db.go: init'ing")
		db.inited = true
		for _, stream := range db.streams {
			if stream.initer != nil {
				db.initErr = stream.initer.InitStream(db.dataRepo)
				if db.initErr != nil {
					log.Println("snapshot/git/gitdb/db.go: error init'ing", db.initErr)
					break
				}
			}
		}
	}

	return db.initErr
}
