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

// https://gobyexample.com/writing-files
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func cleanUpExistingRepo() {
	err := os.RemoveAll(testRepoPath)
	check(err)
}

func createEmptyRepo() {
	err := sh.Command("git", "init", testRepoPath).Run()
	check(err)
}

func createInitialCommit() {
	err := sh.Command(
		"git",
		"commit",
		"--allow-empty",
		"-m", "\"test commit\"",
		sh.Dir(testRepoPath),
	).Run()
	check(err)
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
			check(err)
			Expect(len(rev)).To(Equal(41))
		})

		It("returns an error when the directory isn't a repository", func() {
			tempDirName, err := ioutil.TempDir("", "metadata_test")
			check(err)
			_, err = GetRevision(tempDirName)
			Expect(err).To(HaveOccurred())
		})
	})
})
