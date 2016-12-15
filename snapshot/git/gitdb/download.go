package gitdb

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
)

func (db *DB) download(id SnapID) (SnapID, error) {
	switch id := id.(type) {
	case *localCommitSnapID:
		return id, nil
	case *streamCommitSnapID:
		return db.downloadStreamCommit(id)
	case *bundlestoreSnapID:
		return db.downloadBundlestore(id)
	default:
		return nil, fmt.Errorf("unexpected type %T for %v in creators.go:distribute", id, id)
	}
}

func (db *DB) downloadStreamCommit(id *streamCommitSnapID) (SnapID, error) {
	if err := db.snapPresent(id); err == nil {
		return id, nil
	}

	if _, err := db.dataRepo.Run("fetch"); err != nil {
		return nil, err
	}

	if err := db.snapPresent(id); err != nil {
		return nil, err
	}

	return id, nil
}

func (db *DB) downloadBundlestore(id *bundlestoreSnapID) (SnapID, error) {
	if err := db.snapPresent(id); err == nil {
		return id, nil
	}

	filename, err := db.downloadBundle(id)
	if err != nil {
		return nil, err
	}

	if err := db.downloadBundlePrereqs(filename); err != nil {
		return nil, err
	}

	// unbundle
	if _, err := db.dataRepo.Run("bundle", "unbundle", filename); err != nil {
		return nil, err
	}

	if err := db.snapPresent(id); err != nil {
		return nil, err
	}

	return id, nil
}

func (db *DB) downloadBundle(id *bundlestoreSnapID) (filename string, err error) {
	d, err := db.tmpDir.TempDir("bundle-")
	if err != nil {
		return "", err
	}
	bundleFilename := path.Join(d.Dir, fmt.Sprintf("commit-%s.bundle", id.localSHA()))
	f, err := os.Create(bundleFilename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r, err := db.bundles.OpenForRead(id.baseName())
	if err != nil {
		return "", err
	}
	if _, err := io.Copy(f, r); err != nil {
		return "", err
	}

	return f.Name(), nil

}

var prereqRE = regexp.MustCompile("(?m:^error: ([0-9a-f]{40}) )")

func parsePrereqCommits(verifyStderr string, err error) ([]string, error) {
	// Try to match output like:
	// error: Repository lacks these prerequisite commits:
	// error: 6bddae500d9d103383240f023009ccd4fd627d99

	// TODO(dbentley): this impl match other error lines.
	// We could be more precise by looking for exactly the first line and then
	// only lines that match the line after.

	matches := prereqRE.FindAllStringSubmatch(verifyStderr, -1)
	if len(matches) == 0 {
		// We couldn't find any prereqs
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitErr.Stderr = []byte(verifyStderr)
		}
		return nil, err
	}
	prereqSHAs := []string{}
	for _, match := range matches {
		prereqSHAs = append(prereqSHAs, match[1])
	}
	return prereqSHAs, nil
}

func (db *DB) downloadBundlePrereqs(bundleFilename string) error {
	_, stderr, err := db.dataRepo.RunErr("bundle", "verify", bundleFilename)
	if err == nil {
		return nil
	}

	prereqSHAs, err := parsePrereqCommits(stderr, err)
	if err != nil {
		return nil
	}

	log.Printf("snapshots/git/gitdb/download.go: bundle requires prereqs %v", prereqSHAs)

	if len(db.streams) == 0 {
		return fmt.Errorf("cannot fetch prereqs %v: no streams configured", prereqSHAs)
	}

	// TODO(dbentley): handle multiple streams.
	// How can a bundle refer to a stream?
	stream := db.streams[0]

	for _, prereqSHA := range prereqSHAs {
		commitID := &streamCommitSnapID{streamID: stream.ID, sha: prereqSHA}
		if _, err := db.downloadStreamCommit(commitID); err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) snapPresent(id SnapID) error {
	_, err := db.dataRepo.Run("rev-parse", "--verify", id.localSHA()+"^{object}")
	return err
}
