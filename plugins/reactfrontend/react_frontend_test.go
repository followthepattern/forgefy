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
  - name: backend
    type: react-frontend
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
			f.InstallPlugins(reactfrontend.Plugin{})

			productName, err := f.Forge(yaml, &dummyWriter)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(productName).Should(Equal("Test Product Name"))

			Expect(dummyWriter).Should(forgeio.ContainFiles(
				"apps/backend/app/(account)/profile/page.tsx",
				"apps/backend/app/index.tsx",
				"apps/backend/models/user.ts",
				"apps/backend/app/routes.tsx",
				"apps/backend/app/(account)/photos/hooks/listPhotos.tsx",
				"apps/backend/app/(auth)/login/components/loginForm.tsx",
				"apps/backend/app/(account)/photos/hooks/updatePhotos.tsx",
				"apps/backend/app/(account)/user/new/page.tsx",
				"apps/backend/app/(account)/user/hooks/deleteUser.tsx",
				"apps/backend/app/(account)/user/hooks/listUser.tsx",
				"apps/backend/app/(account)/user/hooks/singleUser.tsx",
				"apps/backend/tsconfig.json",
				"apps/backend/app/(account)/photos/[id]/edit/page.tsx",
				"apps/backend/app/(account)/components/buttons/tertiaryButton.tsx",
				"apps/backend/app/icons/chevronDown.tsx",
				"apps/backend/app/(auth)/login/components/loginLoggedOut.tsx",
				"apps/backend/app/(account)/components/inputFields/textarea.tsx",
				"apps/backend/app/(account)/components/headerSidebar/navbar/navbar.tsx",
				"apps/backend/app/app.tsx",
				"apps/backend/app/(account)/components/modals/confirmModal.tsx",
				"apps/backend/app/(account)/components/singleView/gridFields/gridFields.tsx",
				"apps/backend/app/(account)/user/model.ts",
				"apps/backend/app/(account)/photos/[id]/page.tsx",
				"apps/backend/app/(account)/components/listPage/sectionHeader/sectionHeader.tsx",
				"apps/backend/app/(account)/components/headerSidebar/staticSidebar.tsx",
				"apps/backend/app/(account)/photos/graphql/mutation.ts",
				"apps/backend/docker-compose.yml",
				"apps/backend/app/(account)/components/listPage/listPage.tsx",
				"apps/backend/app/(account)/photos/page.tsx",
				"apps/backend/app/favicon.ico",
				"apps/backend/app/(account)/user/[id]/page.tsx",
				"apps/backend/app/(account)/photos/model.ts",
				"apps/backend/app/(account)/components/listPage/table/table.tsx",
				"apps/backend/app/(account)/components/labels/label.tsx",
				"apps/backend/lib/store.ts",
				"apps/backend/app/(auth)/graphql/register.ts",
				"apps/backend/webpack.config.js",
				"apps/backend/lib/strings.ts",
				"apps/backend/graphql/hooks/useSingle.tsx",
				"apps/backend/app/(account)/components/singleView/layout.tsx",
				"apps/backend/components/withGraphQL.tsx",
				"apps/backend/app/(account)/components/listPage/listPageWrapper/listPageWrapper.tsx",
				"apps/backend/app/(account)/components/listPage/sectionHeader/components/newResourceLink.tsx",
				"apps/backend/app/(account)/user/graphql/mutation.ts",
				"apps/backend/graphql/query.ts",
				"apps/backend/graphql/hooks/useUpdate.tsx",
				"apps/backend/app/(account)/components/toastContainer/accountToastContainer.tsx",
				"apps/backend/tailwind.config.js",
				"apps/backend/app/(auth)/login/page.tsx",
				"apps/backend/public/index.html",
				"apps/backend/app/icons/folder.tsx",
				"apps/backend/postcss.config.js",
				"apps/backend/.env.development",
				"apps/backend/app/(account)/photos/hooks/createPhotos.tsx",
				"apps/backend/app/(account)/components/singleView/dataListView/dataListView.tsx",
				"apps/backend/lib/errorParser.ts",
				"apps/backend/app/(account)/components/headerSidebar/mobileSidebar.tsx",
				"apps/backend/app/(auth)/graphql/profile.ts",
				"apps/backend/lib/constants.ts",
				"apps/backend/app/(account)/components/inputFields/input.tsx",
				"apps/backend/components/withUserContext.tsx",
				"apps/backend/app/(account)/user/hooks/createUser.tsx",
				"apps/backend/app/(account)/components/skeletons/tableSkeleton.tsx",
				"apps/backend/app/(account)/layout.tsx",
				"apps/backend/public/next.svg",
				"apps/backend/styles/globals.css",
				"apps/backend/app/(account)/components/headerSidebar/profileMenu.tsx",
				"apps/backend/app/(account)/photos/hooks/singlePhotos.tsx",
				"apps/backend/graphql/mutation.ts",
				"apps/backend/.gitlab-ci.yml",
				"apps/backend/graphql/hooks/useDelete.tsx",
				"apps/backend/graphql/utils/list.ts",
				"apps/backend/app/(account)/photos/hooks/deletePhotos.tsx",
				"apps/backend/graphql/hooks/useCreate.tsx",
				"apps/backend/app/(auth)/graphql/login.ts",
				"apps/backend/app/icons/plus.tsx",
				"apps/backend/app/icons/search.tsx",
				"apps/backend/app/(account)/components/buttons/primaryButton.tsx",
				"apps/backend/app/(account)/photos/new/page.tsx",
				"apps/backend/app/(account)/user/graphql/query.ts",
				"apps/backend/package.json",
				"apps/backend/app/(account)/components/headerSidebar/headerSidebar.tsx",
				"apps/backend/app/(account)/components/buttons/secondaryButton.tsx",
				"apps/backend/app/(account)/components/headerSidebar/navbar/navigation.ts",
				"apps/backend/app/(account)/components/listPage/listPageWrapper/listingFunctions.ts",
				"apps/backend/app/icons/bars.tsx",
				"apps/backend/graphql/hooks/useList.tsx",
				"apps/backend/.env.production",
				"apps/backend/app/(account)/photos/graphql/query.ts",
				"apps/backend/app/(account)/components/headerSidebar/header.tsx",
				"apps/backend/app/(account)/components/listPage/sectionHeader/components/sortButton.tsx",
				"apps/backend/app/(auth)/register/page.tsx",
				"apps/backend/.gitignore",
				"apps/backend/public/vercel.svg",
				"apps/backend/lib/config.ts",
				"apps/backend/app/(public)/home.tsx",
				"apps/backend/app/(account)/components/listPage/sectionHeader/components/searchInput.tsx",
				"apps/backend/app/(account)/user/[id]/edit/page.tsx",
				"apps/backend/app/(account)/user/page.tsx",
				"apps/backend/app/(account)/error.tsx",
				"apps/backend/app/(auth)/layout.tsx",
				"apps/backend/components/userContext.tsx",
				"apps/backend/app/(account)/components/listPage/pagination/pagination.tsx",
				"apps/backend/app/(account)/user/hooks/updateUser.tsx",
			))
		})
	})
})
