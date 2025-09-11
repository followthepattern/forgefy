package reactfrontend

import (
	"io/fs"
	"path"
	"strings"
	"text/template"

	"github.com/followthepattern/forgefy/forgeio"
	"github.com/followthepattern/forgefy/plugins"
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
			"HTMLType":    parsing.CreateHTMLType(types.Registered),
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

		if !forgeio.IsForgeTemplate(filepath) {
			return nil
		}

		if forgeio.ExcludeMonitoring(filepath, !reactApp.Monitoring) {
			return nil
		}

		content, err := fs.ReadFile(templates.Files, filepath)
		if err != nil {
			return err
		}

		newFilepath := filepath

		if strings.Contains(newFilepath, forgeio.APP_FILE_TOKEN) {
			newFilepath = forgeio.ReplaceAppName(newFilepath, reactApp.AppName)
		}

		newFilepath = forgeio.RemoveTemplateExtension(newFilepath)
		newFilepath = path.Join(productmap.ROOT_DIRECTORY, newFilepath)
		newFilepath = forgeio.CleanFilepath(newFilepath, forgeio.MONITORING_FILE_TOKEN)

		if !strings.Contains(newFilepath, forgeio.FEATURE_TOKEN) {
			file := productmap.NewFile(
				newFilepath,
				string(content),
			).WithData(reactApp).
				WithFuncMap(plugin.parsingFunctions)

			return pm.Insert(file)
		}

		for _, feature := range reactApp.Features {
			newDir, newFile := path.Split(newFilepath)

			newDir = forgeio.ReplaceFeatureName(newDir, naming.ToLowerCamelCase(feature.FeatureName))
			newFile = forgeio.ReplaceFeatureName(newFile, naming.ToUpperCamelCase(feature.FeatureName))

			file := productmap.NewFile(
				path.Join(newDir, newFile),
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
