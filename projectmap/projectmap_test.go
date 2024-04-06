package projectmap_test

import (
	"github.com/followthepattern/forgefy/projectmap"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Project Map", func() {
	Context("Exists", func() {
		It("finds files in project map by filepath", func() {
			directory1 := projectmap.NewDirectory("directory1")
			file1 := projectmap.NewFile("file1", "content1", nil)
			err := directory1.AddFile(file1)
			Expect(err).Should(Succeed())

			directory2 := projectmap.NewDirectory("directory2")
			file2 := projectmap.NewFile("file2", "content2", nil)
			err = directory2.AddFile(file2)
			Expect(err).Should(Succeed())

			err = directory2.AddDirectory(directory1)
			Expect(err).Should(Succeed())

			pm := projectmap.NewProjectMap("project")

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
			pm := projectmap.NewProjectMap("project")

			appleFile := projectmap.NewFile("apple.txt", "apple content", nil)

			err := pm.Insert("/test/apps", appleFile)
			Expect(err).Should(Succeed())

			exists := pm.Exists("/test/apps/apple.txt")
			Expect(exists).Should(BeTrue())
		})
	})
})
