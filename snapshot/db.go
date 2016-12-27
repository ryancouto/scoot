package snapshot

type SnapID interface {
	ID() string
}

type DB interface {
	// Ingest

	// IngestGitCommit ingests the commit identified by commitish from ingestRepo
	IngestGitCommit(ingestRepo *repo.Repository, commitish string) (SnapID, error)

	// Distribute

	// Upload makes sure the Snapshot id is uploaded, returning an ID that can be used
	// anywhere or an error
	Upload(id SnapID) (SnapID, error)

	// Upload makes sure the Snapshot id is downloaded, returning an ID that can be used
	// on this computer or an error
	Download(id SnapID) (SnapID, error)

	// Export

	// GitCheckoutInDataRepo checks out the Snapshot id into the Data Repo.
	// This is ugly and modifies state but is useful for the snaps tool for CI to investigate.
	// Returns the path of the checkout or an error.
	GitCheckoutInDataRepo(id SnapID) (path string, err error)

	// Also implements Checkouter
	Checkouter
}
