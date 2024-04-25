package testdata

import (
	_ "embed"
	"path"

	"github.com/followthepattern/forgefy/plugins/gobackend/apptemplates/tests/integration"
	"github.com/followthepattern/forgefy/productmap"
)

func Directory(appName string) string {
	return path.Join(integration.Directory(appName), "testdata")
}

var (
	//go:embed database.sql.tmpl
	databaseSql string
)

var Files = []productmap.File{
	productmap.NewFile("database.sql", databaseSql),
}
