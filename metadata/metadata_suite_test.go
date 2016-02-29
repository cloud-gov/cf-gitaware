package metadata_test

import (
	"io/ioutil"

	"github.com/codeskyblue/go-sh"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMetadata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Metadata Suite")
}

// helpers

func CreateTempDir() (string, error) {
	return ioutil.TempDir("", "metadata_test")
}

func createEmptyRepo(path string) error {
	return sh.Command("git", "init", path).Run()
}

func createInitialCommit(repoPath string) error {
	return sh.Command(
		"git",
		"commit",
		"--allow-empty",
		"-m", "\"test commit\"",
		sh.Dir(repoPath),
	).Run()
}

func CreateTestRepo() (string, error) {
	path, err := CreateTempDir()
	if err != nil {
		return path, err
	}
	err = createEmptyRepo(path)
	if err != nil {
		return path, err
	}
	err = createInitialCommit(path)
	return path, err
}
