package metadata

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/codeskyblue/go-sh"
)

// RepoMetadata represents the current state of a repository.
type RepoMetadata struct {
	Vcs    string `json:"vcs"`
	Branch string `json:"branch"`
	Ref    string `json:"ref"`
}

// toString takes a slice and converts it to a string, with leading and trailing whitespace removed.
func toString(in []byte) string {
	// http://stackoverflow.com/a/18615786/358804
	return strings.TrimSpace(string(in[:]))
}

// GetBranch returns the current branch name of the provided repository path.
func GetBranch(repoPath string) (string, error) {
	// http://stackoverflow.com/a/12142066/358804
	branch, err := sh.Command(
		"git",
		"rev-parse",
		"--abbrev-ref",
		"HEAD",
		sh.Dir(repoPath),
	).Output()

	return toString(branch), err
}

// GetRevision returns the current revision of the provided repository path.
func GetRevision(repoPath string) (string, error) {
	ref, err := sh.Command(
		"git",
		"rev-parse",
		"HEAD",
		sh.Dir(repoPath),
	).Output()

	return toString(ref), err
}

// GetMetadata returns the metadata about the provided repository.
func GetMetadata(repoPath string) (RepoMetadata, error) {
	data := RepoMetadata{"git", "", ""}

	branch, err := GetBranch(repoPath)
	if err != nil {
		return data, err
	}
	data.Branch = branch

	ref, err := GetRevision(repoPath)
	if err != nil {
		return data, err
	}
	data.Ref = ref

	return data, err
}

// WriteMetadata writes the provided repository metadata to the given path. The directory containing the file should already exist.
func WriteMetadata(filePath string, data RepoMetadata) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, jsonBytes, 0644)
}
