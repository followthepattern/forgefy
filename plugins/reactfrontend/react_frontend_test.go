package reactfrontend_test

import (
	"github.com/followthepattern/forgefy"
	"github.com/followthepattern/forgefy/forgeio"
	"github.com/followthepattern/forgefy/plugins/reactfrontend"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("React Front-end Plugin", func() {
	var (
		dummyWriter = forgeio.NewDummyWriter()
	)

	Context("React front-end plugin Forge", func() {
		var (
			yaml = `
version: 0
product_name: "Test Product Name"
apps:
  - name: frontend1
    type: react-frontend
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
			f.InstallPlugins(reactfrontend.Plugin{})

			productName, err := f.Forge(yaml, &dummyWriter)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(productName).Should(Equal("Test Product Name"))

			Expect(dummyWriter).Should(forgeio.ContainFiles(
				"apps/frontend1/app/(account)/components/modals/confirmModal.tsx",
				"apps/frontend1/.env.development",
				"apps/frontend1/package.json",
				"apps/frontend1/app/(account)/user/new/page.tsx",
				"apps/frontend1/graphql/auth/profile.ts",
				"apps/frontend1/app/(account)/user/hooks/createUser.tsx",
				"apps/frontend1/app/icons/bars.tsx",
				"apps/frontend1/app/routes.tsx",
				"apps/frontend1/app/(account)/components/buttons/primaryButton.tsx",
				"apps/frontend1/app/(account)/components/headerSidebar/profileMenu.tsx",
				"apps/frontend1/.gitignore",
				"apps/frontend1/graphql/query.ts",
				"apps/frontend1/public/index.html",
				"apps/frontend1/app/(auth)/login/page.tsx",
				"apps/frontend1/lib/config.ts",
				"apps/frontend1/lib/constants.ts",
				"apps/frontend1/graphql/utils/list.ts",
				"apps/frontend1/app/(account)/components/headerSidebar/headerSidebar.tsx",
				"apps/frontend1/app/(account)/components/listPage/sectionHeader/components/searchInput.tsx",
				"apps/frontend1/graphql/mutation.ts",
				"apps/frontend1/app/(account)/components/inputFields/input.tsx",
				"apps/frontend1/app/(account)/components/singleView/gridFields/gridFields.tsx",
				"apps/frontend1/graphql/hooks/useUpdate.tsx",
				"apps/frontend1/app/(account)/profile/page.tsx",
				"apps/frontend1/app/(account)/components/headerSidebar/staticSidebar.tsx",
				"apps/frontend1/app/(account)/user/hooks/updateUser.tsx",
				"apps/frontend1/app/(account)/user/graphql/query.ts",
				"apps/frontend1/app/(account)/components/singleView/dataListView/dataListView.tsx",
				"apps/frontend1/app/(public)/home.tsx",
				"apps/frontend1/graphql/hooks/useDelete.tsx",
				"apps/frontend1/app/(account)/components/inputFields/textarea.tsx",
				"apps/frontend1/lib/strings.ts",
				"apps/frontend1/app/(account)/components/headerSidebar/navbar/navigation.ts",
				"apps/frontend1/app/(account)/user/model.ts",
				"apps/frontend1/app/(account)/user/graphql/mutation.ts",
				"apps/frontend1/graphql/auth/login.ts",
				"apps/frontend1/app/(account)/components/labels/label.tsx",
				"apps/frontend1/app/(auth)/login/components/loginLoggedOut.tsx",
				"apps/frontend1/app/index.tsx",
				"apps/frontend1/app/favicon.ico",
				"apps/frontend1/graphql/hooks/useSingle.tsx",
				"apps/frontend1/app/(account)/user/[id]/edit/page.tsx",
				"apps/frontend1/app/(account)/components/listPage/sectionHeader/components/newResourceLink.tsx",
				"apps/frontend1/public/vercel.svg",
				"apps/frontend1/app/(account)/components/listPage/pagination/pagination.tsx",
				"apps/frontend1/lib/store.ts",
				"apps/frontend1/app/(account)/components/toastContainer/accountToastContainer.tsx",
				"apps/frontend1/app/icons/plus.tsx",
				"apps/frontend1/app/icons/folder.tsx",
				"apps/frontend1/app/(account)/user/[id]/page.tsx",
				"apps/frontend1/app/(account)/user/page.tsx",
				"apps/frontend1/models/user.ts",
				"apps/frontend1/.gitlab-ci.yml",
				"apps/frontend1/app/icons/chevronDown.tsx",
				"apps/frontend1/components/withGraphQL.tsx",
				"apps/frontend1/app/(account)/components/buttons/secondaryButton.tsx",
				"apps/frontend1/app/(account)/user/hooks/listUser.tsx",
				"apps/frontend1/graphql/hooks/useList.tsx",
				"apps/frontend1/app/(account)/components/listPage/table/table.tsx",
				"apps/frontend1/tsconfig.json",
				"apps/frontend1/tailwind.config.js",
				"apps/frontend1/webpack.config.js",
				"apps/frontend1/app/(account)/components/skeletons/tableSkeleton.tsx",
				"apps/frontend1/docker-compose.yml",
				"apps/frontend1/app/(account)/components/headerSidebar/header.tsx",
				"apps/frontend1/app/(account)/components/listPage/sectionHeader/sectionHeader.tsx",
				"apps/frontend1/app/(account)/components/headerSidebar/navbar/navbar.tsx",
				"apps/frontend1/components/userContext.tsx",
				"apps/frontend1/app/(account)/components/listPage/listPage.tsx",
				"apps/frontend1/graphql/hooks/useCreate.tsx",
				"apps/frontend1/app/(account)/components/headerSidebar/mobileSidebar.tsx",
				"apps/frontend1/styles/globals.css",
				"apps/frontend1/app/(auth)/layout.tsx",
				"apps/frontend1/public/next.svg",
				"apps/frontend1/app/(account)/error.tsx",
				"apps/frontend1/app/icons/search.tsx",
				"apps/frontend1/app/(account)/components/listPage/listPageWrapper/listingFunctions.ts",
				"apps/frontend1/app/(account)/components/singleView/layout.tsx",
				"apps/frontend1/postcss.config.js",
				"apps/frontend1/app/app.tsx",
				"apps/frontend1/app/(auth)/login/components/loginForm.tsx",
				"apps/frontend1/components/withUserContext.tsx",
				"apps/frontend1/app/(account)/user/hooks/singleUser.tsx",
				"apps/frontend1/.env.production",
				"apps/frontend1/app/(account)/layout.tsx",
				"apps/frontend1/app/(account)/components/listPage/sectionHeader/components/sortButton.tsx",
				"apps/frontend1/app/(account)/components/listPage/listPageWrapper/listPageWrapper.tsx",
				"apps/frontend1/app/(account)/components/buttons/tertiaryButton.tsx",
				"apps/frontend1/lib/errorParser.ts",
				"apps/frontend1/app/(account)/user/hooks/deleteUser.tsx",
			))
		})
	})
})
