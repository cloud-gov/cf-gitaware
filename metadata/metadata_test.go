package metadata

import (
	"github.com/codeskyblue/go-sh"

	"os"
	"testing"
)

const testRepoPath = "tmp"

// https://gobyexample.com/writing-files
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createTestRepo() {
	// clean up existing repository
	err := os.RemoveAll(testRepoPath)
	check(err)

	// create new repository
	err = sh.Command("git", "init", testRepoPath).Run()
	check(err)

	// create an initial commit
	err = sh.Command(
		"git",
		"commit",
		"--allow-empty",
		"-m", "\"test commit\"",
		sh.Dir(testRepoPath),
	).Run()
	check(err)
}

// current git sha
func TestGetRevision(t *testing.T) {
	createTestRepo()

	// TODO
}
