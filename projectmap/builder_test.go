package projectmap_test

import (
	"github.com/followthepattern/forgefy/featureset"
	"github.com/followthepattern/forgefy/projectmap"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Builder", func() {
	Context("Creating Project Map from minimum FeatureSet", func() {
		It("adds local dev files", func() {
			fs := featureset.FeatureSet{
				ProjectName: "product",
			}

			builder := projectmap.NewBuilder(fs)
			pm := builder.Create()

			exists := pm.Exists("docker-compose.yaml")
			Expect(exists).Should(BeTrue())
		})
	})
})
