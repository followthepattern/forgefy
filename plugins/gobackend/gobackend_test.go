package gobackend_test

import (
	"github.com/followthepattern/forgefy"
	"github.com/followthepattern/forgefy/forgeio"
	"github.com/followthepattern/forgefy/plugins/gobackend"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Go Back-end Plugin", func() {
	var dummyWriter forgeio.DummyWriter

	BeforeEach(func() {
		dummyWriter = forgeio.NewDummyWriter()
	})

	Context("Forging vscode config", func() {
		It("should contain vscode config", func() {
			yaml := `
version: 0
product_name: "Test Product Name"
vscode: true
apps:
  - name: backend
    type: go-backend
features:
  - name: photos
    fields:
      - name: ID
        type: string
      - name: Name
        type: string
      - name: Type
        type: string
`

			f := forgefy.New()
			f.InstallPlugins(gobackend.Plugin{})

			productName, err := f.Forge(yaml, &dummyWriter)
			Expect(err).NotTo(HaveOccurred())
			Expect(productName).To(Equal("Test Product Name"))

			Expect(dummyWriter).To(forgeio.ContainFiles(
				".vscode/launch.json",
			))
		})

		It("should not contain vscode config", func() {
			yaml := `
version: 0
product_name: "Test Product Name"
apps:
  - name: backend
    type: go-backend
features:
  - name: photos
    fields:
      - name: ID
        type: string
      - name: Name
        type: string
      - name: Type
        type: string
`

			f := forgefy.New()
			f.InstallPlugins(gobackend.Plugin{})

			productName, err := f.Forge(yaml, &dummyWriter)
			Expect(err).NotTo(HaveOccurred())
			Expect(productName).To(Equal("Test Product Name"))

			Expect(dummyWriter).NotTo(forgeio.ContainFiles(
				".vscode/launch.json",
			))
		})
	})

})
