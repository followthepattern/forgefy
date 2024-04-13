package productmap_test

import (
	"github.com/followthepattern/forgefy/productmap"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("File", func() {
	Context("Content", func() {
		It("parses the template with the given data", func() {
			template := "hello {{.AppName}}"
			file := productmap.NewFile("test", template).WithData(struct {
				AppName string
			}{
				AppName: "test",
			})

			Expect(file.Content()).Should(Equal("hello test"))
		})

		It("doesn't try to parse template if data is nil", func() {
			template := "hello {{.AppName}}"
			file := productmap.NewFile("test", template)

			Expect(file.Content()).Should(Equal("hello {{.AppName}}"))
		})
	})
})
