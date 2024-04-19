package feature

import (
	_ "embed"
	"path"

	"github.com/followthepattern/forgefy/apptemplates/apps/backend/features"
	"github.com/followthepattern/forgefy/productmap"
)

func Directory(productName, appName, featureName string) string {
	return path.Join(features.Directory(productName, appName), featureName)
}

var (
	//go:embed controller.go.tmpl
	controller string
	//go:embed database.go.tmpl
	database string
	//go:embed graphql.go.tmpl
	graphql string
	//go:embed model.go.tmpl
	model string
	//go:embed rest.go.tmpl
	rest string
	//go:embed service.go.tmpl
	service string
)

var Files = []productmap.File{
	productmap.NewFile("controller.go", controller),
	productmap.NewFile("database.go", database),
	productmap.NewFile("graphql.go", graphql),
	productmap.NewFile("model.go", model),
	productmap.NewFile("rest.go", rest),
	productmap.NewFile("service.go", service),
}
