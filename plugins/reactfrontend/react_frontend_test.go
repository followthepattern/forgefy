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
  - name: frontend
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
				"apps/frontend/app/(account)/components/listPage/listPage.tsx",
				"apps/frontend/lib/storage.ts",
				"apps/frontend/app/(account)/profile/page.tsx",
				"apps/frontend/app/(account)/layout.tsx",
				"apps/frontend/app/(account)/components/headerSidebar/staticSidebar.tsx",
				"apps/frontend/app/icons/chevronUp.tsx",
				"apps/frontend/app/(auth)/login/components/loginForm.tsx",
				"apps/frontend/app/(account)/user/hooks/createUser.tsx",
				"apps/frontend/docker-compose.yml",
				"apps/frontend/app/(account)/components/inputFields/checkbox.tsx",
				"apps/frontend/lib/errorParser.ts",
				"apps/frontend/app/(account)/components/skeletons/tableSkeleton.tsx",
				"apps/frontend/app/icons/chevronRight.tsx",
				"apps/frontend/app/(account)/error.tsx",
				"apps/frontend/app/(account)/photos/[id]/edit/page.tsx",
				"apps/frontend/app/(auth)/login/components/loginLoggedOut.tsx",
				"apps/frontend/app/(account)/components/headerSidebar/headerSidebar.tsx",
				"apps/frontend/app/(auth)/graphql/register.ts",
				"apps/frontend/app/(account)/user/graphql/mutation.ts",
				"apps/frontend/app/(account)/user/hooks/singleUser.tsx",
				"apps/frontend/public/vercel.svg",
				"apps/frontend/app/(account)/photos/const.ts",
				"apps/frontend/app/(account)/components/inputFields/timefield.tsx",
				"apps/frontend/app/(account)/photos/new/page.tsx",
				"apps/frontend/lib/store.ts",
				"apps/frontend/app/(account)/components/headerSidebar/navbar/navigation.ts",
				"apps/frontend/app/(account)/components/inputFields/textarea.tsx",
				"apps/frontend/app/(account)/components/listPage/sectionHeader/components/searchInput.tsx",
				"apps/frontend/app/(account)/user/[id]/edit/page.tsx",
				"apps/frontend/postcss.config.js",
				"apps/frontend/webpack.config.js",
				"apps/frontend/app/(account)/components/inputFields/filefield.tsx",
				"apps/frontend/.env.production",
				"apps/frontend/app/(account)/user/page.tsx",
				"apps/frontend/graphql/hooks/useList.tsx",
				"apps/frontend/app/(account)/photos/model.ts",
				"apps/frontend/app/(account)/photos/[id]/page.tsx",
				"apps/frontend/.gitignore",
				"apps/frontend/public/index.html",
				"apps/frontend/app/icons/folder.tsx",
				"apps/frontend/app/(account)/photos/hooks/listPhotos.tsx",
				"apps/frontend/app/(account)/components/listPage/listPageWrapper/listPageWrapper.tsx",
				"apps/frontend/app/(account)/user/hooks/updateUser.tsx",
				"apps/frontend/app/icons/search.tsx",
				"apps/frontend/lib/time.ts",
				"apps/frontend/graphql/hooks/useCreate.tsx",
				"apps/frontend/app/(account)/components/singleView/layout.tsx",
				"apps/frontend/graphql/hooks/useUpdate.tsx",
				"apps/frontend/app/(public)/home.tsx",
				"apps/frontend/app/(account)/photos/hooks/updatePhotos.tsx",
				"apps/frontend/app/components/buttons/primaryButton.tsx",
				"apps/frontend/app/(auth)/graphql/profile.ts",
				"apps/frontend/app/(account)/components/listPage/sectionHeader/components/sortButton.tsx",
				"apps/frontend/graphql/mutation.ts",
				"apps/frontend/app/(account)/components/listPage/pagination/pagination.tsx",
				"apps/frontend/app/(auth)/layout.tsx",
				"apps/frontend/app/(account)/user/const.ts",
				"apps/frontend/app/index.tsx",
				"apps/frontend/lib/config.ts",
				"apps/frontend/graphql/query.ts",
				"apps/frontend/app/icons/spinner.tsx",
				"apps/frontend/app/icons/plus.tsx",
				"apps/frontend/tsconfig.json",
				"apps/frontend/.env.development",
				"apps/frontend/app/icons/bars.tsx",
				"apps/frontend/app/components/buttons/button.tsx",
				"apps/frontend/graphql/hooks/useSingle.tsx",
				"apps/frontend/app/(account)/photos/graphql/query.ts",
				"apps/frontend/app/(account)/user/graphql/query.ts",
				"apps/frontend/app/icons/chevronDown.tsx",
				"apps/frontend/app/icons/chevronLeft.tsx",
				"apps/frontend/app/(account)/photos/graphql/mutation.ts",
				"apps/frontend/components/withGraphQL.tsx",
				"apps/frontend/app/(account)/user/[id]/page.tsx",
				"apps/frontend/app/(auth)/register/page.tsx",
				"apps/frontend/app/(account)/components/singleView/dataListView/dataListView.tsx",
				"apps/frontend/app/favicon.ico",
				"apps/frontend/app/(account)/components/singleView/gridFields/gridFields.tsx",
				"apps/frontend/app/app.tsx",
				"apps/frontend/app/(account)/components/toastContainer/accountToastContainer.tsx",
				"apps/frontend/package.json",
				"apps/frontend/lib/constants.ts",
				"apps/frontend/app/(account)/components/labels/label.tsx",
				"apps/frontend/app/(account)/components/modals/confirmModal.tsx",
				"apps/frontend/app/(account)/user/hooks/deleteUser.tsx",
				"apps/frontend/public/next.svg",
				"apps/frontend/app/(account)/components/headerSidebar/header.tsx",
				"apps/frontend/app/(auth)/graphql/login.ts",
				"apps/frontend/tailwind.config.js",
				"apps/frontend/app/(account)/components/headerSidebar/navbar/navbar.tsx",
				"apps/frontend/app/(account)/photos/hooks/createPhotos.tsx",
				"apps/frontend/app/(account)/user/hooks/listUser.tsx",
				"apps/frontend/app/(account)/photos/hooks/singlePhotos.tsx",
				"apps/frontend/app/(account)/user/new/page.tsx",
				"apps/frontend/app/components/buttons/secondaryButton.tsx",
				"apps/frontend/lib/strings.ts",
				"apps/frontend/app/(account)/components/listPage/table/table.tsx",
				"apps/frontend/app/routes.tsx",
				"apps/frontend/app/(account)/components/headerSidebar/profileMenu.tsx",
				"apps/frontend/app/(account)/components/inputFields/input.tsx",
				"apps/frontend/app/(account)/user/model.ts",
				"apps/frontend/styles/globals.css",
				"apps/frontend/app/(account)/photos/page.tsx",
				"apps/frontend/graphql/hooks/useDelete.tsx",
				"apps/frontend/app/(account)/components/listPage/sectionHeader/sectionHeader.tsx",
				"apps/frontend/app/(account)/photos/hooks/deletePhotos.tsx",
				"apps/frontend/app/(auth)/login/page.tsx",
				"apps/frontend/app/(account)/components/listPage/listPageWrapper/listingFunctions.ts",
				"apps/frontend/app/(account)/components/listPage/sectionHeader/components/newResourceLink.tsx",
				"apps/frontend/graphql/utils/list.ts",
				"apps/frontend/app/(account)/components/headerSidebar/mobileSidebar.tsx",
				"apps/frontend/models/user.ts",
				"apps/frontend/components/withUserContext.tsx",
				"apps/frontend/app/(account)/components/inputFields/datefield.tsx",
				"apps/frontend/app/components/buttons/tertiaryButton.tsx",
				"apps/frontend/app/(account)/components/inputFields/numberfield.tsx",
				"apps/frontend/components/userContext.tsx",
			))
		})
	})
})
