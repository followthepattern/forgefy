package productmap_test

import (
	"github.com/followthepattern/forgefy/productmap"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Product Map", func() {
	Context("Exists", func() {
		It("finds files in product map by filepath", func() {
			file1 := productmap.NewFile("dir1/file1", "content1")
			file2 := productmap.NewFile("dir1/dir2/file2", "content2")

			pm := productmap.NewProductMap()
			pm.Insert(file1)
			pm.Insert(file2)

			exists := pm.Exists("dir1/file1")
			Expect(exists).Should(BeTrue())

			exists = pm.Exists("dir1/dir2/file2")
			Expect(exists).Should(BeTrue())
		})
	})

	Context("Insert", func() {
		It("inserts file successfully", func() {
			pm := productmap.NewProductMap()

			appleFile := productmap.NewFile("test/apps/apple.txt", "apple content")

			err := pm.Insert(appleFile)
			Expect(err).Should(Succeed())

			exists := pm.Exists("test/apps/apple.txt")
			Expect(exists).Should(BeTrue())
		})

		It("returns error inserting the same file twice", func() {
			pm := productmap.NewProductMap()

			appleFile := productmap.NewFile("test/apps/apple.txt", "apple content")

			err := pm.Insert(appleFile)
			Expect(err).Should(Succeed())

			err = pm.Insert(appleFile)
			Expect(err).Should(HaveOccurred())
			Expect(err).Should(MatchError("file test/apps/apple.txt already exists"))
		})
	})
})
