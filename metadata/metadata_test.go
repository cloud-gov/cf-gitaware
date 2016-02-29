package metadata_test

import (
	. "github.com/18F/cf-gitaware/metadata"

	"github.com/codeskyblue/go-sh"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"os"
)

const testRepoPath = "tmp"

func cleanUpExistingRepo() error {
	return os.RemoveAll(testRepoPath)
}

func createEmptyRepo() error {
	return sh.Command("git", "init", testRepoPath).Run()
}

func createInitialCommit() error {
	return sh.Command(
		"git",
		"commit",
		"--allow-empty",
		"-m", "\"test commit\"",
		sh.Dir(testRepoPath),
	).Run()
}

func createTestRepo() {
	cleanUpExistingRepo()
	createEmptyRepo()
	createInitialCommit()
}

var _ = Describe("Metadata", func() {
	Describe(".GetRevision()", func() {
		It("returns a 41-character string", func() {
			createTestRepo()

			rev, err := GetRevision(testRepoPath)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(rev)).To(Equal(41))
		})

		It("returns an error when the directory isn't a repository", func() {
			tempDirName, err := ioutil.TempDir("", "metadata_test")
			Expect(err).NotTo(HaveOccurred())
			_, err = GetRevision(tempDirName)
			Expect(err).To(HaveOccurred())
		})
	})
})
