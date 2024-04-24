package testdata

import (
	_ "embed"
	"path"

	"github.com/followthepattern/forgefy/apptemplates/apps/backend"
	"github.com/followthepattern/forgefy/productmap"
)

func Directory(appName string) string {
	return path.Join(backend.Directory(appName), "testdata")
}

var (
	//go:embed database.sql.tmpl
	databaseSql string
)

var Files = []productmap.File{
	productmap.NewFile("database.sql", databaseSql),
}
