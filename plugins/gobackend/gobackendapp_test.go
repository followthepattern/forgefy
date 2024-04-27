package gobackend_test

import (
	"github.com/followthepattern/forgefy"
	"github.com/followthepattern/forgefy/forgeio"
	"github.com/followthepattern/forgefy/plugins/gobackend"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Go Back-end Plugin", func() {
	var (
		dummyWriter = forgeio.NewDummyWriter()
	)

	Context("Go back-end plugin Forge", func() {
		var (
			yaml = `
version: 0
product_name: "Test Product Name"
apps:
  - name: backend1
    type: go-backend
features:
  - name: user
    fields:
      - name: ID
        type: string
      - name: email
        type: string
      - name: FirstName
        type: string
      - name: LastName
        type: string
      - name: Age
        type: int
`
		)

		It("forges files", func() {
			f := forgefy.New()
			f.InstallPlugins(gobackend.Plugin{})

			productName, err := f.Forge(yaml, &dummyWriter)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(productName).Should(Equal("Test Product Name"))

			Expect(dummyWriter).Should(forgeio.ContainFiles(
				"docker-compose.yaml",
				"apps/backend1/features/user/controller.go",
				"apps/backend1/features/user/model.go",
				"apps/backend1/features/user/database.go",
				"apps/backend1/features/user/graphql.go",
				"apps/backend1/features/user/rest.go",
				"apps/backend1/features/user/service.go",
			))
		})
	})
})
