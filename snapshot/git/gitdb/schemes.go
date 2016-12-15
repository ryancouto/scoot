package gitdb

import (
	"fmt"
	"strings"
)

func (db *DB) ParseID(id string) (SnapID, error) {
	// TODO(dbentley): move to entry.go if it uses any members

	return parseID(id)
}

type SnapID interface {
	ID() string
	localSHA() string
}

func parseID(id string) (SnapID, error) {
	switch {
	case id == "":
		return nil, fmt.Errorf("empty snapshot ID")
	case strings.HasPrefix(id, "local-commit-"):
		sha := strings.TrimPrefix(id, "local-commit-")
		sha, err := validateSha(sha)
		if err != nil {
			return nil, err
		}
		return &localCommitSnapID{sha}, nil
	case strings.HasPrefix(id, "bundlestore-"):
		sha := strings.TrimPrefix(id, "bundlestore-")
		sha, err := validateSha(sha)
		if err != nil {
			return nil, err
		}
		return &bundlestoreSnapID{sha}, nil
	default:
		return nil, fmt.Errorf("unrecognized snapshot ID %q", id)
	}
}

func validateSha(sha string) (string, error) {
	if len(sha) != 40 {
		return "", fmt.Errorf("sha not 40 characters in %s", sha)
	}
	return sha, nil
}

type localCommitSnapID struct {
	sha string
}

func (id *localCommitSnapID) ID() string {
	return fmt.Sprintf("local-commit-%s", id.sha)
}

func (id *localCommitSnapID) localSHA() string {
	return id.sha
}

type streamCommitSnapID struct {
	streamID CommitStreamID
	sha      string
}

func (id *streamCommitSnapID) ID() string {
	return fmt.Sprintf("stream-commit-%s-%s", id.sha, id.streamID.name)
}

func (id *streamCommitSnapID) localSHA() string {
	return id.sha
}

type bundlestoreSnapID struct {
	sha string
}

func (id *bundlestoreSnapID) ID() string {
	return fmt.Sprintf("bundlestore-%s", id.sha)
}

func (id *bundlestoreSnapID) localSHA() string {
	return id.sha
}

func (id *bundlestoreSnapID) baseName() string {
	return fmt.Sprintf("%s.bundle", id.sha)
}
