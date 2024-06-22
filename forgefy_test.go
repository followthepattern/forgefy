package forgefy_test

import (
	"github.com/followthepattern/forgefy"
	"github.com/followthepattern/forgefy/forgeio"
	"github.com/followthepattern/forgefy/plugins/monorepo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Forgefy", func() {
	var (
		dummyWriter = forgeio.NewDummyWriter()
	)

	Context("No plugin Forge", func() {
		var (
			yaml = `
version: 0
product_name: "test product"
`
		)

		It("forges files", func() {
			f := forgefy.New()

			f.InstallPlugins(monorepo.MonoRepo{})

			productName, err := f.Forge(yaml, &dummyWriter)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(productName).Should(Equal("test product"))

			Expect(dummyWriter).Should(forgeio.ContainFiles("docker-compose.yaml"))
		})
	})
})
