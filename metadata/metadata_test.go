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

	Describe(".GetMetadata()", func() {
		It("returns the relevant information about the repository", func() {
			err := InitializeRepo(repoPath)
			Expect(err).NotTo(HaveOccurred())

			data, err := GetMetadata(repoPath)
			Expect(err).NotTo(HaveOccurred())

			Expect(data.Vcs).To(Equal("git"))
			Expect(len(data.Ref)).To(Equal(41))
			Expect(data.Branch).To(Equal("master"))
		})
	})
})
