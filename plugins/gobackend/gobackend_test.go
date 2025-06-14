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

	Context("Go back-end plugin Forge", func() {
		var (
			yaml = `
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
		)

		It("forges files", func() {
			f := forgefy.New()
			f.InstallPlugins(gobackend.Plugin{})

			productName, err := f.Forge(yaml, &dummyWriter)
			Expect(err).NotTo(HaveOccurred())
			Expect(productName).To(Equal("Test Product Name"))

			Expect(dummyWriter).To(forgeio.ContainFiles())
		})
	})

	Context("Forging e2e tests using dagger", func() {
		It("includes dagger files", func() {
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

			Expect(err).To(Succeed())

			Expect(dummyWriter).To(forgeio.ContainFiles(
				".dagger/backend.go",
			))
		})

		It("excludes dagger files", func() {
			yaml := `
version: 0
product_name: "Test Product Name"
exclude_dagger: true
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
			Expect(productName).NotTo(Equal("Test Product Name"))

			Expect(err).To(Succeed())

			Expect(dummyWriter).NotTo(forgeio.ContainFiles(
				".dagger/backend.go",
			))
		})
	})
})
