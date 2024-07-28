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
  - name: backend
    type: go-backend
features:
  - name: photos
    fields:
      - name: ID
        type: string
      - name: Name
        type: string
      - name: Type
        type: string
`
		)

		It("forges files", func() {
			f := forgefy.New()
			f.InstallPlugins(gobackend.Plugin{})

			productName, err := f.Forge(yaml, &dummyWriter)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(productName).Should(Equal("Test Product Name"))

			Expect(dummyWriter).Should(forgeio.ContainFiles(
				"apps/backend/config/cerbos.go",
				"apps/backend/db/migrations/1_initialize_schema.up.sql",
				"apps/backend/policies/user.yaml",
				"apps/backend/api/graphql/schema/schema.go",
				"apps/backend/features/mail/client.go",
				"apps/backend/features/photos/rest.go",
				"apps/backend/types/type.go",
				"apps/backend/config/jwtkeypair.go",
				"apps/backend/models/response_status.go",
				"apps/backend/config/server.go",
				"apps/backend/api/middlewares/session_context_id.go",
				"apps/backend/features/auth/authentication.go",
				"apps/backend/api/graphql/resolver.go",
				"apps/backend/api/middlewares/api_logger.go",
				"apps/backend/models/userlog.go",
				"apps/backend/types/bytes.go",
				"apps/backend/types/convert.go",
				"apps/backend/features/auth/context.go",
				"apps/backend/db/migrations/1_initialize_schema.down.sql",
				"apps/backend/api/graphql/schema/schema.graph",
				"apps/backend/features/mail/model.go",
				"apps/backend/features/user/graphql.go",
				"apps/backend/api/responses.go",
				"apps/backend/api/middlewares/token.go",
				"apps/backend/features/photos/graphql.go",
				"apps/backend/features/user/database.go",
				"apps/backend/features/auth/database.go",
				"apps/backend/configs/config.yaml.tmpl",
				"apps/backend/features/auth/authorization.go",
				"apps/backend/accesscontrol/accesscontrol.go",
				"apps/backend/features/photos/service.go",
				"apps/backend/policies/photos.yaml",
				"apps/backend/repositories/database/sqlbuilder/pagination.go",
				"apps/backend/Dockerfile",
				"apps/backend/types/bool.go",
				"apps/backend/go.mod",
				"apps/backend/features/photos/controller.go",
				"apps/backend/api/rest/handler.go",
				"apps/backend/api/middlewares/heartbeat.go",
				"apps/backend/api/middlewares/jwt.go",
				"apps/backend/cmd/backend/main.go",
				"apps/backend/features/auth/model.go",
				"apps/backend/types/time.go",
				"apps/backend/docker-compose.yml",
				"apps/backend/configs/dev_public.pem",
				"apps/backend/configs/config.yaml",
				"apps/backend/configs/dev_private.pem",
				"apps/backend/features/photos/database.go",
				"apps/backend/config/mail.go",
				"apps/backend/features/user/rest.go",
				"apps/backend/api/httpresponses/responses.go",
				"apps/backend/config/organization.go",
				"apps/backend/container/container.go",
				"apps/backend/features/user/service.go",
				"apps/backend/controllers/controllers.go",
				"apps/backend/api/graphql/handler.go",
				"apps/backend/features/user/controller.go",
				"apps/backend/config/cfg.go",
				"apps/backend/types/int.go",
				"apps/backend/types/validation.go",
				"apps/backend/policies/.cerbos.yaml",
				"apps/backend/.gitignore",
				"apps/backend/types/int64.go",
				"apps/backend/types/string.go",
				"apps/backend/tests/integration/testdata/database.sql",
				"apps/backend/hostserver/server.go",
				"apps/backend/features/mail/service.go",
				"apps/backend/models/list.go",
				"apps/backend/features/user/models/user_model.go",
				"apps/backend/features/photos/models/photos_model.go",
				"apps/backend/features/auth/graphql.go",
				"apps/backend/config/db.go",
				"apps/backend/features/auth/rest.go",
				"apps/backend/features/auth/controller.go",
				"apps/backend/types/uint.go",
				"apps/backend/api/api.go",
			))
		})
	})
})
