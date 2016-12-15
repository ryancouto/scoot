package gitdb

import (
	"fmt"

	"github.com/scootdev/scoot/snapshot"
)

func (db *DB) gitCheckoutInDataRepo(id SnapID) (path string, err error) {
	// TODO(dbentley): it's messy that this is stateful.
	// We'd rather hand out temporary checkouts. Git allows this with work-tree,
	// but we're stuck on git 2.0 or 2.4. Remove this method and switch to work trees when possible.
	id, err = db.download(id)
	if err != nil {
		return "", err
	}

	cmds := [][]string{
		// -d removes directories. -x ignores gitignore and removes everything.
		// -f is force. -f the second time removes directories even if they're git repos themselves
		{"clean", "-f", "-f", "-d", "-x"},
		{"checkout", id.localSHA()},
	}
	for _, argv := range cmds {
		if _, err := db.dataRepo.Run(argv...); err != nil {
			return "", fmt.Errorf("error checking out %v as %v: %v", id, id.localSHA(), err)
		}
	}

	return db.dataRepo.Dir(), nil
}

func (db *DB) checkout(id string) (co snapshot.Checkout, err error) {
	// TODO(dbentley): lock this repo in some way

	snapID, err := db.ParseID(id)
	if err != nil {
		return nil, err
	}

	path, err := db.gitCheckoutInDataRepo(snapID)
	if err != nil {
		return nil, err
	}

	return &UnmanagedCheckout{id: snapID.ID(), dir: path}, nil
}

// User-owned checkout.
type UnmanagedCheckout struct {
	id  string
	dir string
}

func (c *UnmanagedCheckout) Path() string {
	return c.dir
}

func (c *UnmanagedCheckout) ID() string {
	return c.id
}

func (c *UnmanagedCheckout) Release() error {
	return nil
}
