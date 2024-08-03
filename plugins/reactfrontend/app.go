package reactfrontend

import (
	"io/fs"
	"path"
	"strings"
	"text/template"

	"github.com/followthepattern/forgefy/plugins"
	"github.com/followthepattern/forgefy/plugins/monorepo/templates/apps"
	"github.com/followthepattern/forgefy/plugins/reactfrontend/parsing"
	"github.com/followthepattern/forgefy/plugins/reactfrontend/templates"
	"github.com/followthepattern/forgefy/productmap"
	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/naming"
	"github.com/followthepattern/forgefy/specification/types"
)

var _ plugins.App = &ReactFrontend{}

type ReactFrontend struct {
	port             int
	tailwindPort     int
	parsingFunctions template.FuncMap
}

func NewApp() *ReactFrontend {
	return &ReactFrontend{
		port:         3000,
		tailwindPort: 9999,
		parsingFunctions: template.FuncMap{
			"JSType":      parsing.CreateJSType(types.Registered),
			"HTMLType":    parsing.HTMLType,
			"URL":         parsing.URL,
			"GraphQLName": parsing.GraphQLName,
		},
	}
}

func (ReactFrontend) Name() string {
	return "React Front-end application"
}

func (ReactFrontend) Type() string {
	return "react-frontend"
}

func (plugin *ReactFrontend) GetNextPortNumber() int {
	port := plugin.port

	plugin.port++

	return port
}

func (plugin *ReactFrontend) GetNextTailwindPortNumber() int {
	port := plugin.tailwindPort

	plugin.tailwindPort--

	return port
}

type App struct {
	specification.App
	TailwindPort int
}

func (a App) AppNameCapital() string {
	return strings.ToUpper(a.AppName)
}

func (a App) AppNameToPackageName() string {
	return strings.ToLower(a.AppName)
}

func (ReactFrontend) AddDefaultFiles(pm productmap.ProductMap, fs specification.Product) error {
	return nil
}

func (plugin *ReactFrontend) Build(pm productmap.ProductMap, app specification.App) error {
	dir := templates.Files

	reactApp := App{app, 9999}

	reactApp.App.AppPort = plugin.GetNextPortNumber()
	reactApp.TailwindPort = plugin.GetNextTailwindPortNumber()

	return fs.WalkDir(dir, ".", plugin.createWalkFn(pm, reactApp))
}

func ToFileSuffix(f specification.Feature) string {
	return naming.ToUpperCamelCase(f.FeatureName)
}

func (plugin ReactFrontend) createWalkFn(pm productmap.ProductMap, reactApp App) func(filepath string, d fs.DirEntry, err error) error {
	return func(filepath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(filepath, ".tmpl") {
			return nil
		}

		content, err := fs.ReadFile(templates.Files, filepath)
		if err != nil {
			return err
		}

		filepath = strings.TrimSuffix(filepath, ".tmpl")

		filepath = path.Join(apps.Directory(), reactApp.AppName, filepath)

		if !strings.Contains(filepath, "(feature)") {
			file := productmap.NewFile(
				filepath,
				string(content),
			).WithData(reactApp).
				WithFuncMap(plugin.parsingFunctions)

			return pm.Insert(file)
		}

		for _, feature := range reactApp.Features {
			newFilePath := strings.Replace(filepath, "(feature)", feature.FeatureNameDir(), 1)
			newFilePath = strings.Replace(newFilePath, "(feature)", ToFileSuffix(feature), 1)

			file := productmap.NewFile(
				newFilePath,
				string(content),
			).WithData(feature).
				WithFuncMap(plugin.parsingFunctions)

			err = pm.Insert(file)
			if err != nil {
				return err
			}
		}

		return nil
	}
}
