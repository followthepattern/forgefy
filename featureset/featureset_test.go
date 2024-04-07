package featureset_test

import (
	"github.com/followthepattern/forgefy/featureset"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Feature Set", func() {
	Context("Unmarshal yaml file", func() {
		var (
			yaml = `
version: 0
product_name: "test product"`
		)

		It("populates the feature set object with the values from the yaml", func() {
			fs, err := featureset.UnmarshalYaml([]byte(yaml))
			Expect(err).Should(Succeed())

			Expect(fs.ProductName).Should(Equal("test product"))
			Expect(fs.Version).Should(Equal("0"))
		})
	})
})
