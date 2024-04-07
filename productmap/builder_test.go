package productmap_test

import (
	"github.com/followthepattern/forgefy/featureset"
	"github.com/followthepattern/forgefy/productmap"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Builder", func() {
	Context("Creating Product Map from minimum FeatureSet", func() {
		var (
			yaml = `
product_name: "test product"
apps:
  - name: backend1
    type: backend`
		)
		It("checks if local dev files are available", func() {
			fs, err := featureset.UnmarshalYaml([]byte(yaml))
			Expect(err).Should(Succeed())

			builder := productmap.NewBuilder(fs)
			pm := builder.Build()

			exists := pm.Exists("docker-compose.yaml")
			Expect(exists).Should(BeTrue())
		})
	})

	Context("Creating Product Map from extended FeatureSet", func() {
		var (
			yaml = `
version: 0
product_name: "test product"
apps:
  - name: frontend1
    type: frontend
  - name: backend1
    type: backend
`
		)
		It("checks if frontend files are added", func() {
			fs, err := featureset.UnmarshalYaml([]byte(yaml))
			Expect(err).Should(Succeed())

			builder := productmap.NewBuilder(fs)
			pm := builder.Build()

			exists := pm.Exists("frontend1/package.json")
			Expect(exists).Should(BeTrue())
		})

		It("checks if backend files are added", func() {
			fs, err := featureset.UnmarshalYaml([]byte(yaml))
			Expect(err).Should(Succeed())

			builder := productmap.NewBuilder(fs)
			pm := builder.Build()

			exists := pm.Exists("backend1/go.mod")
			Expect(exists).Should(BeTrue())
		})
	})
})
