package metadata_test

import (
	. "github.com/18F/cf-gitaware/metadata"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Metadata", func() {
	var repoPath string

	BeforeEach(func() {
		var err error
		repoPath, err = CreateTempDir()
		Expect(err).NotTo(HaveOccurred())
	})

	Describe(".GetRevision()", func() {
		It("returns a 41-character string", func() {
			err := InitializeRepo(repoPath)
			Expect(err).NotTo(HaveOccurred())

			rev, err := GetRevision(repoPath)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(rev)).To(Equal(41))
		})

		It("returns an error when the directory isn't a repository", func() {
			_, err := GetRevision(repoPath)
			Expect(err).To(HaveOccurred())
		})
	})
})
