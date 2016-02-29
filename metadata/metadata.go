package metadata

import (
	"github.com/codeskyblue/go-sh"
)

// RepoMetadata represents the current state of a repository.
type RepoMetadata struct {
	Vcs    string `json:"vcs"`
	Ref    string `json:"ref"`
	Branch string `json:"branch"`
}

// GetRevision returns the current revision of the provided repository path.
func GetRevision(repoPath string) ([]byte, error) {
	return sh.Command(
		"git",
		"rev-parse",
		"HEAD",
		sh.Dir(repoPath),
	).Output()
}

// GetMetadata returns the metadata about the provided repository.
func GetMetadata(repoPath string) (RepoMetadata, error) {
	ref, err := GetRevision(repoPath)

	// TODO fetch branch
	// http://stackoverflow.com/a/18615786/358804
	data := RepoMetadata{"git", string(ref[:]), "master"}

	return data, err
}
