package productmap_test

import (
	"github.com/followthepattern/forgefy/productmap"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Product Map", func() {
	Context("Exists", func() {
		It("finds files in product map by filepath", func() {
			directory1 := productmap.NewDirectory("directory1")
			file1 := productmap.NewFile("file1", "content1", nil)
			err := directory1.AddFile(file1)
			Expect(err).Should(Succeed())

			directory2 := productmap.NewDirectory("directory2")
			file2 := productmap.NewFile("file2", "content2", nil)
			err = directory2.AddFile(file2)
			Expect(err).Should(Succeed())

			err = directory2.AddDirectory(directory1)
			Expect(err).Should(Succeed())

			pm := productmap.NewProductMap("product")

			err = pm.AddDirectory(directory2)
			Expect(err).Should(Succeed())

			exists := pm.Exists("directory2/directory1/file1")
			Expect(exists).Should(BeTrue())

			exists = pm.Exists("directory2/file2")
			Expect(exists).Should(BeTrue())
		})
	})

	Context("Insert", func() {
		It("inserts file successfully", func() {
			pm := productmap.NewProductMap("product")

			appleFile := productmap.NewFile("apple.txt", "apple content", nil)

			err := pm.Insert("/test/apps", appleFile)
			Expect(err).Should(Succeed())

			exists := pm.Exists("/test/apps/apple.txt")
			Expect(exists).Should(BeTrue())
		})
	})
})
