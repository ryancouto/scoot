// Package repo provides utilities for operating on a git repo.
// Scoot often ends up with multiple git repos. E.g., one reference repo
// and then each checkout is in its own repo.
package repo

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// Repository represents a valid Git repository.
type Repository struct {
	dir string
}

// Where r lives on disk
func (r *Repository) Dir() string {
	return r.dir
}

func (r *Repository) RunErr(args ...string) (stdout string, stderr string, err error) {
	cmd := exec.Command("git", args...)
	stderrBuf := bytes.NewBuffer(nil)
	cmd.Stderr = stderrBuf
	stdout, err = r.runCmd(cmd)
	return stdout, stderrBuf.String(), err
}

// Run a git command in r
func (r *Repository) Run(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	return r.runCmd(cmd)
}

func (r *Repository) runCmd(cmd *exec.Cmd) (string, error) {
	cmd.Dir = r.dir
	log.Println("repo.Repository.Run", cmd.Args)
	data, err := cmd.Output()
	log.Println("repo.Repository.Run complete", err)
	if err != nil {
		err, ok := err.(*exec.ExitError)
		if ok && len(err.Stderr) > 0 {
			log.Println("repo.Repository.Run error:", string(err.Stderr))
		}
	}
	return string(data), err
}

// Run a git command that returns a sha.
func (r *Repository) RunSha(args ...string) (string, error) {
	out, err := r.Run(args...)
	if err != nil {
		return out, err
	}
	return validateSha(out)
}

func validateSha(sha string) (string, error) {
	if len(sha) == 40 || len(sha) == 41 && sha[40] == '\n' {
		return sha[0:40], nil
	}
	return "", fmt.Errorf("sha not 40 or 41 (with a \\n) characters: %q", sha)
}

// NewRepo creates a new Repository for path `dir`.
// It checks that `dir` is a valid path.
func NewRepository(dir string) (*Repository, error) {
	// TODO(dbentley): make sure we handle the case that path is in a git repo,
	// but is not the root of a git repo
	repo := &Repository{dir}
	// TODO(dbentley): we'd prefer to use features present in git 2.5+, but are stuck on 2.4 for now
	// E.g., --resolve-git-dir or --git-path
	topLevel, err := repo.Run("rev-parse", "--show-toplevel")
	if err != nil {
		return nil, err
	}
	topLevel = strings.Replace(topLevel, "\n", "", -1)
	log.Println("git.NewRepository:", dir, topLevel)
	repo.dir = topLevel
	return repo, nil
}
