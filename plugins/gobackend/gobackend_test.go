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
  - name: {{ $app | PackageName }}
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
				"apps/{{ $app | PackageName }}/features/user/controller.go",
				"apps/{{ $app | PackageName }}/policies/.cerbos.yaml",
				"apps/{{ $app | PackageName }}/features/auth/rest.go",
				"apps/{{ $app | PackageName }}/api/graphql/schema/schema.go",
				"apps/{{ $app | PackageName }}/features/auth/authorization.go",
				"apps/{{ $app | PackageName }}/features/auth/model.go",
				"apps/{{ $app | PackageName }}/models/list.go",
				"apps/{{ $app | PackageName }}/features/user/model.go",
				"apps/{{ $app | PackageName }}/cmd/{{ $app | PackageName }}/main.go",
				"apps/{{ $app | PackageName }}/accesscontrol/accesscontrol.go",
				"apps/{{ $app | PackageName }}/api/middlewares/jwt.go",
				"apps/{{ $app | PackageName }}/api/middlewares/heartbeat.go",
				"apps/{{ $app | PackageName }}/features/user/graphql.go",
				"apps/{{ $app | PackageName }}/repositories/database/sqlbuilder/pagination.go",
				"apps/{{ $app | PackageName }}/Dockerfile",
				"apps/{{ $app | PackageName }}/config/cerbos.go",
				"apps/{{ $app | PackageName }}/types/uint.go",
				"apps/{{ $app | PackageName }}/types/type.go",
				"apps/{{ $app | PackageName }}/features/auth/database.go",
				"apps/{{ $app | PackageName }}/api/graphql/handler.go",
				"apps/{{ $app | PackageName }}/tests/integration/testdata/database.sql",
				"apps/{{ $app | PackageName }}/types/bytes.go",
				"apps/{{ $app | PackageName }}/config/mail.go",
				"apps/{{ $app | PackageName }}/models/userlog.go",
				"apps/{{ $app | PackageName }}/types/int64.go",
				"apps/{{ $app | PackageName }}/features/user/rest.go",
				"apps/{{ $app | PackageName }}/config/organization.go",
				"apps/{{ $app | PackageName }}/api/responses.go",
				"apps/{{ $app | PackageName }}/features/auth/context.go",
				"apps/{{ $app | PackageName }}/docker-compose.yml",
				"apps/{{ $app | PackageName }}/config/jwtkeypair.go",
				"apps/{{ $app | PackageName }}/api/graphql/schema/schema.graph",
				"apps/{{ $app | PackageName }}/hostserver/server.go",
				"apps/{{ $app | PackageName }}/go.mod",
				"apps/{{ $app | PackageName }}/config/server.go",
				"apps/{{ $app | PackageName }}/api/middlewares/session_context_id.go",
				"apps/{{ $app | PackageName }}/features/auth/graphql.go",
				"apps/{{ $app | PackageName }}/types/int.go",
				"apps/{{ $app | PackageName }}/models/response_status.go",
				"apps/{{ $app | PackageName }}/api/middlewares/token.go",
				"apps/{{ $app | PackageName }}/api/api.go",
				"apps/{{ $app | PackageName }}/features/mail/model.go",
				"apps/{{ $app | PackageName }}/container/container.go",
				"apps/{{ $app | PackageName }}/.gitignore",
				"apps/{{ $app | PackageName }}/config/db.go",
				"apps/{{ $app | PackageName }}/policies/user.yaml",
				"apps/{{ $app | PackageName }}/features/user/database.go",
				"apps/{{ $app | PackageName }}/features/user/service.go",
				"apps/{{ $app | PackageName }}/types/validation.go",
				"apps/{{ $app | PackageName }}/features/mail/service.go",
				"apps/{{ $app | PackageName }}/configs/config.yaml",
				"apps/{{ $app | PackageName }}/features/auth/authentication.go",
				"apps/{{ $app | PackageName }}/types/bool.go",
				"apps/{{ $app | PackageName }}/types/convert.go",
				"apps/{{ $app | PackageName }}/configs/dev_private.pem",
				"apps/{{ $app | PackageName }}/config/cfg.go",
				"apps/{{ $app | PackageName }}/api/middlewares/api_logger.go",
				"apps/{{ $app | PackageName }}/features/mail/client.go",
				"apps/{{ $app | PackageName }}/api/rest/handler.go",
				"apps/{{ $app | PackageName }}/api/httpresponses/responses.go",
				"apps/{{ $app | PackageName }}/types/string.go",
				"apps/{{ $app | PackageName }}/types/time.go",
				"apps/{{ $app | PackageName }}/controllers/controllers.go",
				"apps/{{ $app | PackageName }}/configs/dev_public.pem",
				"apps/{{ $app | PackageName }}/api/graphql/resolver.go",
				"apps/{{ $app | PackageName }}/features/auth/controller.go",
				"apps/{{ $app | PackageName }}/configs/config.yaml.tmpl",
			))
		})
	})
})
