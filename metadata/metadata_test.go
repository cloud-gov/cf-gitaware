package metadata

import (
	"github.com/codeskyblue/go-sh"

	"io/ioutil"
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

func TestGetRevision(t *testing.T) {
	createTestRepo()

	rev, err := GetRevision(testRepoPath)
	check(err)
	chars := len(rev)
	if chars != 41 {
		t.Errorf("Expected GetRevision to return a 41-character string, Found %s\n", rev)
	}

	tempDirName, err := ioutil.TempDir("", "metadata_test")
	check(err)
	rev, err = GetRevision(tempDirName)
	if err == nil {
		t.Errorf("Expected GetRevision to give an error when the directory isn't a repository.")
	}
}
