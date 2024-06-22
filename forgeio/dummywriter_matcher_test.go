package forgeio_test

import (
	"github.com/followthepattern/forgefy/forgeio"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DummyWriter Matcher", func() {
	var dummyWriter forgeio.DummyWriter

	BeforeEach(func() {
		dummyWriter = forgeio.NewDummyWriter()
	})

	Context("Write", func() {
		It("finds the files in writer", func() {
			dummyWriter.Write("here/is/test/dir/test.file", nil)
			dummyWriter.Write("there/is/test/dir/other.file", nil)

			Expect(dummyWriter).Should(forgeio.ContainFiles(
				"here/is/test/dir/test.file",
				"there/is/test/dir/other.file",
			))
		})

		It("doesn't find a file in writer", func() {
			dummyWriter.Write("here/is/test/dir/test.file", nil)
			dummyWriter.Write("there/is/test/dir/other.file", nil)

			Expect(dummyWriter).ShouldNot(forgeio.ContainFiles(
				"there/is/test/dir/nonexistent.file",
			))
		})
	})
})
