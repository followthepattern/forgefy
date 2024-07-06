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
  - name: {{ .AppNamePackage }}
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
				"apps/{{ .AppNamePackage }}/features/user/controller.go",
				"apps/{{ .AppNamePackage }}/policies/.cerbos.yaml",
				"apps/{{ .AppNamePackage }}/features/auth/rest.go",
				"apps/{{ .AppNamePackage }}/api/graphql/schema/schema.go",
				"apps/{{ .AppNamePackage }}/features/auth/authorization.go",
				"apps/{{ .AppNamePackage }}/features/auth/model.go",
				"apps/{{ .AppNamePackage }}/models/list.go",
				"apps/{{ .AppNamePackage }}/features/user/model.go",
				"apps/{{ .AppNamePackage }}/cmd/{{ .AppNamePackage }}/main.go",
				"apps/{{ .AppNamePackage }}/accesscontrol/accesscontrol.go",
				"apps/{{ .AppNamePackage }}/api/middlewares/jwt.go",
				"apps/{{ .AppNamePackage }}/api/middlewares/heartbeat.go",
				"apps/{{ .AppNamePackage }}/features/user/graphql.go",
				"apps/{{ .AppNamePackage }}/repositories/database/sqlbuilder/pagination.go",
				"apps/{{ .AppNamePackage }}/Dockerfile",
				"apps/{{ .AppNamePackage }}/config/cerbos.go",
				"apps/{{ .AppNamePackage }}/types/uint.go",
				"apps/{{ .AppNamePackage }}/types/type.go",
				"apps/{{ .AppNamePackage }}/features/auth/database.go",
				"apps/{{ .AppNamePackage }}/api/graphql/handler.go",
				"apps/{{ .AppNamePackage }}/tests/integration/testdata/database.sql",
				"apps/{{ .AppNamePackage }}/types/bytes.go",
				"apps/{{ .AppNamePackage }}/config/mail.go",
				"apps/{{ .AppNamePackage }}/models/userlog.go",
				"apps/{{ .AppNamePackage }}/types/int64.go",
				"apps/{{ .AppNamePackage }}/features/user/rest.go",
				"apps/{{ .AppNamePackage }}/config/organization.go",
				"apps/{{ .AppNamePackage }}/api/responses.go",
				"apps/{{ .AppNamePackage }}/features/auth/context.go",
				"apps/{{ .AppNamePackage }}/docker-compose.yml",
				"apps/{{ .AppNamePackage }}/config/jwtkeypair.go",
				"apps/{{ .AppNamePackage }}/api/graphql/schema/schema.graph",
				"apps/{{ .AppNamePackage }}/hostserver/server.go",
				"apps/{{ .AppNamePackage }}/go.mod",
				"apps/{{ .AppNamePackage }}/config/server.go",
				"apps/{{ .AppNamePackage }}/api/middlewares/session_context_id.go",
				"apps/{{ .AppNamePackage }}/features/auth/graphql.go",
				"apps/{{ .AppNamePackage }}/types/int.go",
				"apps/{{ .AppNamePackage }}/models/response_status.go",
				"apps/{{ .AppNamePackage }}/api/middlewares/token.go",
				"apps/{{ .AppNamePackage }}/api/api.go",
				"apps/{{ .AppNamePackage }}/features/mail/model.go",
				"apps/{{ .AppNamePackage }}/container/container.go",
				"apps/{{ .AppNamePackage }}/.gitignore",
				"apps/{{ .AppNamePackage }}/config/db.go",
				"apps/{{ .AppNamePackage }}/policies/user.yaml",
				"apps/{{ .AppNamePackage }}/features/user/database.go",
				"apps/{{ .AppNamePackage }}/features/user/service.go",
				"apps/{{ .AppNamePackage }}/types/validation.go",
				"apps/{{ .AppNamePackage }}/features/mail/service.go",
				"apps/{{ .AppNamePackage }}/configs/config.yaml",
				"apps/{{ .AppNamePackage }}/features/auth/authentication.go",
				"apps/{{ .AppNamePackage }}/types/bool.go",
				"apps/{{ .AppNamePackage }}/types/convert.go",
				"apps/{{ .AppNamePackage }}/configs/dev_private.pem",
				"apps/{{ .AppNamePackage }}/config/cfg.go",
				"apps/{{ .AppNamePackage }}/api/middlewares/api_logger.go",
				"apps/{{ .AppNamePackage }}/features/mail/client.go",
				"apps/{{ .AppNamePackage }}/api/rest/handler.go",
				"apps/{{ .AppNamePackage }}/api/httpresponses/responses.go",
				"apps/{{ .AppNamePackage }}/types/string.go",
				"apps/{{ .AppNamePackage }}/types/time.go",
				"apps/{{ .AppNamePackage }}/controllers/controllers.go",
				"apps/{{ .AppNamePackage }}/configs/dev_public.pem",
				"apps/{{ .AppNamePackage }}/api/graphql/resolver.go",
				"apps/{{ .AppNamePackage }}/features/auth/controller.go",
				"apps/{{ .AppNamePackage }}/configs/config.yaml.tmpl",
			))
		})
	})
})
