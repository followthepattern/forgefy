package productmap_test

import (
	"github.com/followthepattern/forgefy/featureset"
	"github.com/followthepattern/forgefy/productmap"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Builder", func() {
	Context("Creating Product Map from minimum FeatureSet", func() {
		It("adds local dev files", func() {
			fs := featureset.FeatureSet{
				ProductName: "product",
			}

			builder := productmap.NewBuilder(fs)
			pm := builder.Create()

			exists := pm.Exists("docker-compose.yaml")
			Expect(exists).Should(BeTrue())
		})
	})
})
