package productmap_test

import (
	"github.com/followthepattern/forgefy/productmap"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Product Map", func() {
	Context("Exists", func() {
		It("finds files in product map by filepath", func() {
			file1 := productmap.NewFile("file1", "content1")
			file2 := productmap.NewFile("file2", "content2")

			pm := productmap.NewProductMap()
			dir1, err := pm.AddChild("dir1")
			Expect(err).ShouldNot(HaveOccurred())
			err = dir1.AddFile(file1)
			Expect(err).ShouldNot(HaveOccurred())

			dir2, err := dir1.AddChild("dir2")
			Expect(err).ShouldNot(HaveOccurred())
			err = dir2.AddFile(file2)
			Expect(err).ShouldNot(HaveOccurred())

			exists := pm.Exists("dir1/file1")
			Expect(exists).Should(BeTrue())

			exists = pm.Exists("dir1/dir2/file2")
			Expect(exists).Should(BeTrue())
		})
	})

	Context("Insert", func() {
		It("inserts file successfully", func() {
			pm := productmap.NewProductMap()

			appleFile := productmap.NewFile("apple.txt", "apple content")

			err := pm.Insert("test/apps", appleFile)
			Expect(err).Should(Succeed())

			exists := pm.Exists("test/apps/apple.txt")
			Expect(exists).Should(BeTrue())
		})
	})

	Context("Dirname", func() {
		It("successfully returns with a directory name", func() {
			pm := productmap.NewProductMap()
			dirName := "apps/backend/features/feature"

			appleFile := productmap.NewFile("apple.txt", "apple content")

			err := pm.Insert(dirName, appleFile)
			Expect(err).Should(Succeed())

			dir := pm.FindDirectory(dirName)
			Expect(dir).ShouldNot(BeNil())

			Expect(dir.DirName()).Should(Equal(dirName))

		})
	})
})
