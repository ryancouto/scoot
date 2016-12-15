package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/scootdev/scoot/os/temp"
	"github.com/scootdev/scoot/snapshot/bundlestore"
	"github.com/scootdev/scoot/snapshot/git/gitdb"
	"github.com/scootdev/scoot/snapshot/git/repo"
	"github.com/scootdev/scoot/snapshot/snaps/cli"
)

func main() {
	inj := &injector{}
	cmd := cli.MakeSnapsCLI(inj)
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

type injector struct {
	repo string
}

func (i *injector) Register(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().StringVar(&i.repo, "data_repo", "", "repo for scoot to use to store its data")
}

func (i *injector) Inject() (*gitdb.DB, error) {
	tempDir, err := temp.TempDirDefault()
	if err != nil {
		return nil, err
	}
	if i.repo == "" {
		wd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		log.Println("--scoot_repo unset; defaulting to current directory", wd)
		i.repo = wd
	}
	dataRepo, err := repo.NewRepository(i.repo)
	if err != nil {
		return nil, fmt.Errorf("could not create a repo at %v: %v", i.repo, err)
	}

	sourceMaster := gitdb.NewCommitStream("source_master", "remotes/origin/master", nil)
	bundles := bundlestore.MakeHTTPStore("http://bundles--0--scoot-bundleserver--devel--dbentley.service.smf1.twitter.biz/bundle/")

	return gitdb.MakeDB(dataRepo, []*gitdb.CommitStream{sourceMaster}, tempDir, bundles), nil
}
