package forgeio

import (
	"strings"

	"github.com/followthepattern/forgefy/specification/naming"
)

const (
	TMPL_SUFFIX       = ".tmpl"
	DAGGER_FILE_TOKEN = "($dagger$)"
	APP_FILE_TOKEN    = "[(application)]"
	FEATURE_TOKEN     = "[(feature)]"
)

func IsForgeTemplate(filepath string) bool {
	return strings.HasSuffix(filepath, TMPL_SUFFIX)
}

func ExcludeTemplate(filepath string, excludeDagger bool) bool {
	return excludeDagger && strings.Contains(filepath, DAGGER_FILE_TOKEN)
}

func RemoveTemplateExtension(filepath string) string {
	return strings.TrimSuffix(filepath, TMPL_SUFFIX)
}

func CleanFilepath(filepath string, fileToken string) string {
	return strings.ReplaceAll(filepath, fileToken, "")
}

func ReplaceAppName(filepath string, appName string) string {
	return strings.ReplaceAll(filepath, APP_FILE_TOKEN, appName)
}

func ReplaceFeatureName(filepath, feature string) string {
	return strings.ReplaceAll(filepath, FEATURE_TOKEN, naming.ToLowerSnakeCase(feature))
}
