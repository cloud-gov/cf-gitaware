package metadata_test

import (
	. "github.com/18F/cf-gitaware/metadata"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Metadata", func() {
	Describe(".GetRevision()", func() {
		It("returns a 41-character string", func() {
			repoPath, err := CreateTestRepo()
			Expect(err).NotTo(HaveOccurred())

			rev, err := GetRevision(repoPath)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(rev)).To(Equal(41))
		})

		It("returns an error when the directory isn't a repository", func() {
			tempDirName, err := CreateTempDir()
			Expect(err).NotTo(HaveOccurred())

			_, err = GetRevision(tempDirName)
			Expect(err).To(HaveOccurred())
		})
	})
})
