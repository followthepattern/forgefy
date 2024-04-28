package reactfrontend

import (
	"fmt"

	"github.com/followthepattern/forgefy/plugins"
	"github.com/followthepattern/forgefy/plugins/reactfrontend/apptemplates"
	"github.com/followthepattern/forgefy/productmap"
	"github.com/followthepattern/forgefy/specification"
)

var _ plugins.App = &ReactFrontend{}

type ReactFrontend struct{}

func (ReactFrontend) Name() string {
	return "React Front-end application"
}

func (ReactFrontend) Type() string {
	return "react-frontend"
}

func (ReactFrontend) DockerComposeInfos(appName, appPath string) []plugins.DockerComposeInfo {
	return []plugins.DockerComposeInfo{
		{
			ServiceName:           appName,
			DockerComposeFilePath: fmt.Sprintf("%s/docker-compose.yml", appPath),
		},
	}
}

func (ReactFrontend) AddDefaultFiles(pm productmap.ProductMap, fs specification.Product) error {
	return nil
}

func (plugin ReactFrontend) Builder(pm productmap.ProductMap, fs specification.Product, app specification.App) error {
	dir := apptemplates.Directory(app.AppName)

	return pm.Insert(dir,
		apptemplates.PackageJSON.WithData(app))
}
