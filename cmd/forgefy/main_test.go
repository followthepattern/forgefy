package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testing CLI", func() {
	Context("Excluding files from deleting", func() {
		var (
			excludePathes = map[string]struct{}{}
		)

		BeforeEach(func() {
			excludePathes = map[string]struct{}{
				"apps/to/exclude": {},
			}
		})

		It("checks if it excludes the files that match the pattern", func() {
			var result bool
			result = pathToExclude(excludePathes, "apps/to/exclude/apple.txt")
			Expect(result).To(BeTrue())

			result = pathToExclude(excludePathes, "apps/to/exclude/banana.txt")
			Expect(result).To(BeTrue())

			result = pathToExclude(excludePathes, "apps/to/exclude")
			Expect(result).To(BeTrue())

			result = pathToExclude(excludePathes, "apps/to")
			Expect(result).To(BeTrue())

			result = pathToExclude(excludePathes, "apps")
			Expect(result).To(BeTrue())
		})

		It("does not exclude files that don't match with the pattern", func() {
			var result bool
			result = pathToExclude(excludePathes, "apps/not/to/exclude/apple.txt")
			Expect(result).To(BeFalse())

			result = pathToExclude(excludePathes, "apps/not/to/copy/apple.txt")
			Expect(result).To(BeFalse())

			result = pathToExclude(excludePathes, "apps/not/to/copy")
			Expect(result).To(BeFalse())

			result = pathToExclude(excludePathes, "apps/just/exclude/apple.txt")
			Expect(result).To(BeFalse())

			result = pathToExclude(excludePathes, "exclude.txt")
			Expect(result).To(BeFalse())

			result = pathToExclude(excludePathes, "apps/exclude.txt")
			Expect(result).To(BeFalse())
		})

		It("excludes folder in the root", func() {
			excludePathes["apps"] = struct{}{}

			var result bool
			result = pathToExclude(excludePathes, "apps/to/exclude/apple.txt")
			Expect(result).To(BeTrue())

			result = pathToExclude(excludePathes, "apps/to/copy/apple.txt")
			Expect(result).To(BeTrue())

			result = pathToExclude(excludePathes, "apps/to/copy")
			Expect(result).To(BeTrue())

			result = pathToExclude(excludePathes, "apps/just/exclude/apple.txt")
			Expect(result).To(BeTrue())

			result = pathToExclude(excludePathes, "no.txt")
			Expect(result).To(BeFalse())

			result = pathToExclude(excludePathes, "other/not/to/exclude.txt")
			Expect(result).To(BeFalse())

			result = pathToExclude(excludePathes, "apple.txt")
			Expect(result).To(BeFalse())

			result = pathToExclude(excludePathes, "")
			Expect(result).To(BeTrue())

			result = pathToExclude(excludePathes, ".")
			Expect(result).To(BeTrue())
		})
	})
})
