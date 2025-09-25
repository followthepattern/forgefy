package flutter

import (
	"io/fs"
	"path"
	"strings"
	"text/template"

	"github.com/followthepattern/forgefy/forgeio"
	"github.com/followthepattern/forgefy/plugins"
	"github.com/followthepattern/forgefy/plugins/flutter/templates"
	"github.com/followthepattern/forgefy/productmap"
	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/naming"
)

var _ plugins.App = &FlutterPluginApp{}

type FlutterPluginApp struct {
	parsingFunctions template.FuncMap
}

func NewApp() *FlutterPluginApp {
	return &FlutterPluginApp{
		parsingFunctions: template.FuncMap{},
	}
}

func (FlutterPluginApp) Name() string {
	return "Flutter application"
}

func (FlutterPluginApp) Type() string {
	return "flutter-mobile"
}

func (plugin *FlutterPluginApp) Build(pm productmap.ProductMap, app specification.App) error {
	dir := templates.EntireDir

	pm.WithFuncMap(plugin.parsingFunctions)

	return fs.WalkDir(dir, ".", plugin.createWalkFn(pm, app))
}

func (b FlutterPluginApp) createWalkFn(pm productmap.ProductMap, app specification.App) func(filepath string, d fs.DirEntry, err error) error {
	dir := templates.EntireDir

	return func(filepath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !forgeio.IsForgeTemplate(filepath) {
			return nil
		}

		if forgeio.ExcludeMonitoring(filepath, !app.Monitoring) {
			return nil
		}

		if forgeio.ExcludeVSCode(filepath, !app.VSCode) {
			return nil
		}

		content, err := fs.ReadFile(dir, filepath)
		if err != nil {
			return err
		}

		newFilepath := filepath

		if strings.Contains(newFilepath, forgeio.APP_FILE_TOKEN) {
			newFilepath = forgeio.ReplaceAppName(newFilepath, naming.ToLowerCamelCase(app.AppName))
		}

		newFilepath = forgeio.RemoveTemplateExtension(newFilepath)
		newFilepath = path.Join(productmap.ROOT_DIRECTORY, newFilepath)
		newFilepath = forgeio.CleanFilepath(newFilepath, forgeio.MONITORING_FILE_TOKEN)
		newFilepath = forgeio.CleanFilepath(newFilepath, forgeio.VSCODE_FILE_TOKEN)

		if !strings.Contains(newFilepath, forgeio.FEATURE_TOKEN) {
			file := productmap.NewFile(
				newFilepath,
				string(content)).
				WithData(app)

			err := pm.Insert(file)
			if err != nil {
				return err
			}
			return nil
		}

		for _, feature := range app.Features {
			file := productmap.NewFile(
				forgeio.ReplaceFeatureName(newFilepath, naming.ToLowerSnakeCase(feature.FeatureName)),
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
