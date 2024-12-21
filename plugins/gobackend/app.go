package gobackend

import (
	"io/fs"
	"path"
	"strings"
	"text/template"

	"github.com/followthepattern/forgefy/forgeio"
	"github.com/followthepattern/forgefy/plugins"
	"github.com/followthepattern/forgefy/plugins/gobackend/parsing"
	"github.com/followthepattern/forgefy/plugins/gobackend/templates"
	"github.com/followthepattern/forgefy/productmap"
	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/types"
)

type App struct {
	specification.App
	DbPort     int
	CerbosPort int
}

func newApp(app specification.App) App {
	return App{app, 0, 0}
}

func (a App) AppNameCapital() string {
	return strings.ToUpper(a.AppName)
}

var _ plugins.App = &GoBackendPluginApp{}

type GoBackendPluginApp struct {
	port             int
	dbPort           int
	cerbosPort       int
	parsingFunctions template.FuncMap
}

func NewApp() *GoBackendPluginApp {
	return &GoBackendPluginApp{
		port:       8080,
		cerbosPort: 3592,
		dbPort:     5433,
		parsingFunctions: template.FuncMap{
			"DBName":      parsing.DBName,
			"DBType":      parsing.CreateDBType(types.Registered),
			"ValueDB":     parsing.ValueDB,
			"DBNullable":  parsing.CreateDBNullable(types.Registered),
			"NameGraphQL": parsing.NameGraphQL,
			"TypeGraphQL": parsing.CreateTypeGraphQL(types.Registered),
			"GoType":      parsing.CreateGoType(types.Registered),
			"AsTag":       parsing.AsTag,
			"AsURL":       parsing.AsURL,
			"PackageName": parsing.PackageName,
		},
	}
}

func (GoBackendPluginApp) Name() string {
	return "Go Back-end application"
}

func (GoBackendPluginApp) Type() string {
	return "go-backend"
}

func (plugin *GoBackendPluginApp) GetNextPortNumber() int {
	port := plugin.port

	plugin.port++

	return port
}

func (plugin *GoBackendPluginApp) GetNextDBPort() int {
	port := plugin.dbPort

	plugin.dbPort++

	return port
}

func (plugin *GoBackendPluginApp) GetNextCerbosPort() int {
	port := plugin.cerbosPort

	plugin.cerbosPort++

	return port
}

func (plugin *GoBackendPluginApp) Build(pm productmap.ProductMap, app specification.App) error {
	dir := templates.EntireDir

	goApp := newApp(app)
	goApp = plugin.setDefaults(goApp)

	pm.WithFuncMap(plugin.parsingFunctions)

	return fs.WalkDir(dir, ".", plugin.createWalkFn(pm, goApp))
}

func (plugin *GoBackendPluginApp) setDefaults(goApp App) App {
	goApp.App.AppPort = plugin.GetNextPortNumber()
	goApp.DbPort = plugin.GetNextDBPort()
	goApp.CerbosPort = plugin.GetNextCerbosPort()

	return goApp
}

func (b GoBackendPluginApp) createWalkFn(pm productmap.ProductMap, goApp App) func(filepath string, d fs.DirEntry, err error) error {
	dir := templates.EntireDir

	return func(filepath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !forgeio.IsForgeTemplate(filepath) {
			return nil
		}

		if forgeio.ExcludeTemplate(filepath, goApp.ExcludeDagger) {
			return nil
		}

		content, err := fs.ReadFile(dir, filepath)
		if err != nil {
			return err
		}

		newFilepath := filepath

		if strings.Contains(newFilepath, forgeio.APP_FILE_TOKEN) {
			newFilepath = forgeio.ReplaceAppName(newFilepath, goApp.AppName)
		}

		newFilepath = forgeio.RemoveTemplateExtension(newFilepath)
		newFilepath = path.Join(productmap.ROOT_DIRECTORY, newFilepath)
		newFilepath = forgeio.CleanFilepath(newFilepath, forgeio.DAGGER_FILE_TOKEN)

		if !strings.Contains(newFilepath, "[feature]") {
			file := productmap.NewFile(
				newFilepath,
				string(content)).
				WithData(goApp)

			err := pm.Insert(file)
			if err != nil {
				return err
			}
			return nil
		}

		for _, feature := range goApp.Features {
			newFilepath := strings.ReplaceAll(newFilepath, "[feature]", parsing.PackageName(feature))

			file := productmap.NewFile(
				newFilepath,
				string(content),
			).WithData(feature)

			err = pm.Insert(file)
			if err != nil {
				return err
			}
		}

		return nil
	}
}
