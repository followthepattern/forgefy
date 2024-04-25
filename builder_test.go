package forgefy_test

import (
	"github.com/followthepattern/forgefy"
	"github.com/followthepattern/forgefy/featureset"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Builder", func() {
	Context("Creating Product Map from minimum FeatureSet", func() {
		var (
			yaml = `
product_name: "test"
apps:
  - name: backend1
    type: backend`
		)
		It("checks if local dev files are available", func() {
			fs, err := featureset.UnmarshalYaml([]byte(yaml))
			Expect(err).Should(Succeed())

			builder := forgefy.NewBuilder(fs)
			pm, err := builder.Build()
			Expect(err).Should(Succeed())

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
  - name: backend1
    type: backend
  - name: frontend1
    type: frontend
`
			fs featureset.FeatureSet
		)

		BeforeEach(func() {
			var err error
			fs, err = featureset.UnmarshalYaml([]byte(yaml))
			Expect(err).Should(Succeed())
		})

		It("checks if frontend files are added", func() {
			builder := forgefy.NewBuilder(fs)
			pm, err := builder.Build()
			Expect(err).Should(Succeed())

			exists := pm.Exists("apps/frontend1/package.json")
			Expect(exists).Should(BeTrue())
		})

		It("checks if backend files are added", func() {
			builder := forgefy.NewBuilder(fs)
			pm, err := builder.Build()
			Expect(err).Should(Succeed())

			exists := pm.Exists("apps/backend1/go.mod")
			Expect(exists).Should(BeTrue())
		})
	})
})
