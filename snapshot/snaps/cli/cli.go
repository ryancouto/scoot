package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/scootdev/scoot/snapshot/git/gitdb"
	"github.com/scootdev/scoot/snapshot/git/repo"
)

// DBInjector
type DBInjector interface {

	// Register registeres any necessary flags on the root command
	Register(rootCmd *cobra.Command)

	// Inject injects the DB that is necessary for all commands
	// TODO(dbentley): change this to be a snapshot.DB
	Inject() (*gitdb.DB, error)
}

// MakeSnapsCLI makes the CLI for the Snaps tool.
func MakeSnapsCLI(injector DBInjector) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "snaps",
		Short: "scoot snaps",
	}
	parentCmd := rootCmd
	add := func(subCmd dbCommand) {
		cmd := subCmd.register()
		cmd.RunE = func(innerCmd *cobra.Command, args []string) error {
			db, err := injector.Inject()
			if err != nil {
				return fmt.Errorf("snaps could not create db: %v", err)
			}
			return subCmd.run(db, innerCmd, args)
		}
		parentCmd.AddCommand(cmd)
	}

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "create a snapshot",
	}
	rootCmd.AddCommand(createCmd)
	parentCmd = createCmd

	add(&ingestGitCommitCommand{})
	add(&uploadCommand{})
	add(&ingestGitCommitAndUploadCommand{})
	add(&downloadCommand{})

	exportCmd := &cobra.Command{
		Use:   "export",
		Short: "export a snapshot",
	}
	rootCmd.AddCommand(exportCmd)
	parentCmd = exportCmd

	add(&gitCheckoutDataRepoCommand{})

	return rootCmd
}

type dbCommand interface {
	// register registers any local flags, returning the command
	register() *cobra.Command

	// run runs the command with the provided DB
	run(db *gitdb.DB, cmd *cobra.Command, args []string) error
}

type ingestGitCommitCommand struct {
	ingestRepo string
	commitish  string
}

func (c *ingestGitCommitCommand) register() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ingest_git_commit",
		Short: "ingest a git commit",
	}
	cmd.Flags().StringVar(&c.ingestRepo, "ingest_repo", "", "repo to ingest from")
	cmd.Flags().StringVar(&c.commitish, "commitish", "", "commitish to ingest")
	return cmd
}

func (c *ingestGitCommitCommand) run(db *gitdb.DB, _ *cobra.Command, _ []string) error {
	if c.ingestRepo == "" {
		return fmt.Errorf("must specify a repo to ingest from")
	}
	ingestRepo, err := repo.NewRepository(c.ingestRepo)
	if err != nil {
		return fmt.Errorf("not a valid repo dir: %v, %v", c.ingestRepo, err)
	}

	id, err := db.IngestGitCommit(ingestRepo, c.commitish)
	if err != nil {
		return err
	}

	fmt.Println(id.ID())
	return nil
}

type uploadCommand struct {
	snapID string
}

func (c *uploadCommand) register() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload",
		Short: "upload a snapshot",
	}
	cmd.Flags().StringVar(&c.snapID, "snap_id", "", "snapshot to upload")
	return cmd
}

func (c *uploadCommand) run(db *gitdb.DB, _ *cobra.Command, _ []string) error {
	id, err := db.ParseID(c.snapID)
	if err != nil {
		return err
	}

	uploadedID, err := db.Upload(id)
	if err != nil {
		return err
	}

	fmt.Println(uploadedID.ID())
	return nil
}

type ingestGitCommitAndUploadCommand struct {
	ingestRepo string
	commitish  string
}

func (c *ingestGitCommitAndUploadCommand) register() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ingest_git_commit_and_upload",
		Short: "ingest a git commit and upload it",
	}
	cmd.Flags().StringVar(&c.ingestRepo, "ingest_repo", "", "repo to ingest from")
	cmd.Flags().StringVar(&c.commitish, "commitish", "", "commitish to ingest")
	return cmd
}

func (c *ingestGitCommitAndUploadCommand) run(db *gitdb.DB, _ *cobra.Command, _ []string) error {
	if c.ingestRepo == "" {
		return fmt.Errorf("must specify a repo to ingest from")
	}
	ingestRepo, err := repo.NewRepository(c.ingestRepo)
	if err != nil {
		return fmt.Errorf("not a valid repo dir: %v, %v", c.ingestRepo, err)
	}

	id, err := db.IngestGitCommitAndUpload(ingestRepo, c.commitish)
	if err != nil {
		return err
	}

	fmt.Println(id.ID())
	return nil
}

type downloadCommand struct {
	snapID string
}

func (c *downloadCommand) register() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download",
		Short: "download a snapshot",
	}
	cmd.Flags().StringVar(&c.snapID, "snap_id", "", "snapshot to download")
	return cmd
}

func (c *downloadCommand) run(db *gitdb.DB, _ *cobra.Command, _ []string) error {
	id, err := db.ParseID(c.snapID)
	if err != nil {
		return err
	}

	downloadedID, err := db.Download(id)
	if err != nil {
		return err
	}

	fmt.Println(downloadedID.ID())
	return nil
}

type gitCheckoutDataRepoCommand struct {
	snapID string
}

func (c *gitCheckoutDataRepoCommand) register() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "git_checkout_data_repo",
		Short: "checkout snap_id in the data repo, modifying its state (!!!)",
	}
	cmd.Flags().StringVar(&c.snapID, "snap_id", "", "snapshot to checkout")
	return cmd
}

func (c *gitCheckoutDataRepoCommand) run(db *gitdb.DB, _ *cobra.Command, _ []string) error {
	id, err := db.ParseID(c.snapID)
	if err != nil {
		return err
	}

	path, err := db.GitCheckoutInDataRepo(id)
	if err != nil {
		return err
	}

	fmt.Println(path)
	return nil
}
