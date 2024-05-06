package gobackend

import (
	"fmt"

	"github.com/followthepattern/forgefy/plugins"
)

func (GoBackendPluginApp) DockerComposeInfos(appName, appPath string) []plugins.DockerComposeInfo {
	return []plugins.DockerComposeInfo{
		{
			ServiceName:           "db",
			DockerComposeFilePath: fmt.Sprintf("%s/docker-compose.yml", appPath),
		},
		{
			ServiceName:           "cerbos",
			DockerComposeFilePath: fmt.Sprintf("%s/docker-compose.yml", appPath),
		},
		{
			ServiceName:           appName,
			DockerComposeFilePath: fmt.Sprintf("%s/docker-compose.yml", appPath),
		},
	}
}
