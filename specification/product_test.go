package specification_test

import (
	"github.com/followthepattern/forgefy/specification"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Specification", func() {
	Context("Unmarshal yaml file", func() {
		var (
			yaml = `
forge_version: 0
product_name: "test product"
apps:
  - name: "Test frontend"
    type: "frontend"
  - name: "Test backend"
    type: "backend"
`
		)

		It("populates the product specification with the values from the yaml", func() {
			fs, err := specification.UnmarshalYaml([]byte(yaml))
			Expect(err).Should(Succeed())

			Expect(fs.ProductName).Should(Equal("test product"))
			Expect(fs.ForgeVersion).Should(Equal("0"))
		})
	})
})
