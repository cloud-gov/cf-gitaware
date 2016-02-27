package metadata

import (
	"github.com/codeskyblue/go-sh"
)

// GetRevision returns the current revision of the provided repository path.
func GetRevision(repoPath string) ([]byte, error) {
	return sh.Command(
		"git",
		"rev-parse",
		"HEAD",
		sh.Dir(repoPath),
	).Output()
}
