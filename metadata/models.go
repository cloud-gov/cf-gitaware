package metadata

// RepoMetadata represents the current state of a repository.
type RepoMetadata struct {
	Vcs    string `json:"vcs"`
	Branch string `json:"branch"`
	Ref    string `json:"ref"`
}
