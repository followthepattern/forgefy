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
				"apps/backend1/features/user/controller.go",
				"apps/backend1/policies/.cerbos.yaml",
				"apps/backend1/features/auth/rest.go",
				"apps/backend1/api/graphql/schema/schema.go",
				"apps/backend1/features/auth/authorization.go",
				"apps/backend1/features/auth/model.go",
				"apps/backend1/models/list.go",
				"apps/backend1/features/user/model.go",
				"apps/backend1/cmd/backend1/main.go",
				"apps/backend1/accesscontrol/accesscontrol.go",
				"apps/backend1/api/middlewares/jwt.go",
				"apps/backend1/api/middlewares/heartbeat.go",
				"apps/backend1/features/user/graphql.go",
				"apps/backend1/repositories/database/sqlbuilder/pagination.go",
				"apps/backend1/Dockerfile",
				"apps/backend1/config/cerbos.go",
				"apps/backend1/types/uint.go",
				"apps/backend1/types/type.go",
				"apps/backend1/features/auth/database.go",
				"apps/backend1/api/graphql/handler.go",
				"apps/backend1/tests/integration/testdata/database.sql",
				"apps/backend1/types/bytes.go",
				"apps/backend1/config/mail.go",
				"apps/backend1/models/userlog.go",
				"apps/backend1/types/int64.go",
				"apps/backend1/features/user/rest.go",
				"apps/backend1/config/organization.go",
				"apps/backend1/api/responses.go",
				"apps/backend1/features/auth/context.go",
				"apps/backend1/docker-compose.yml",
				"apps/backend1/config/jwtkeypair.go",
				"apps/backend1/api/graphql/schema/schema.graph",
				"apps/backend1/hostserver/server.go",
				"apps/backend1/go.mod",
				"apps/backend1/config/server.go",
				"apps/backend1/api/middlewares/session_context_id.go",
				"apps/backend1/features/auth/graphql.go",
				"apps/backend1/types/int.go",
				"apps/backend1/models/response_status.go",
				"apps/backend1/api/middlewares/token.go",
				"apps/backend1/api/api.go",
				"apps/backend1/features/mail/model.go",
				"apps/backend1/container/container.go",
				"apps/backend1/.gitignore",
				"apps/backend1/config/db.go",
				"apps/backend1/policies/user.yaml",
				"apps/backend1/features/user/database.go",
				"apps/backend1/features/user/service.go",
				"apps/backend1/types/validation.go",
				"apps/backend1/features/mail/service.go",
				"apps/backend1/configs/config.yaml",
				"apps/backend1/features/auth/authentication.go",
				"apps/backend1/types/bool.go",
				"apps/backend1/types/convert.go",
				"apps/backend1/configs/dev_private.pem",
				"apps/backend1/config/cfg.go",
				"apps/backend1/api/middlewares/api_logger.go",
				"apps/backend1/features/mail/client.go",
				"apps/backend1/api/rest/handler.go",
				"apps/backend1/api/httpresponses/responses.go",
				"apps/backend1/types/string.go",
				"apps/backend1/types/time.go",
				"apps/backend1/controllers/controllers.go",
				"apps/backend1/configs/dev_public.pem",
				"apps/backend1/api/graphql/resolver.go",
				"apps/backend1/features/auth/controller.go",
				"apps/backend1/configs/config.yaml.tmpl",
			))
		})
	})
})
