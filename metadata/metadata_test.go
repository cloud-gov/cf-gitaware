package metadata_test

import (
	. "github.com/18F/cf-gitaware/metadata"

	"github.com/codeskyblue/go-sh"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
)

func createTempDir() (string, error) {
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

func createTestRepo() (string, error) {
	path, err := createTempDir()
	createEmptyRepo(path)
	createInitialCommit(path)
	return path, err
}

var _ = Describe("Metadata", func() {
	Describe(".GetRevision()", func() {
		It("returns a 41-character string", func() {
			repoPath, err := createTestRepo()
			Expect(err).NotTo(HaveOccurred())

			rev, err := GetRevision(repoPath)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(rev)).To(Equal(41))
		})

		It("returns an error when the directory isn't a repository", func() {
			tempDirName, err := createTempDir()
			Expect(err).NotTo(HaveOccurred())
			_, err = GetRevision(tempDirName)
			Expect(err).To(HaveOccurred())
		})
	})
})
