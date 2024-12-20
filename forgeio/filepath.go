package forgeio

import "strings"

const (
	TMPL_SUFFIX       = ".tmpl"
	DAGGER_FILE_TOKEN = "._dagger_"
)

func IsForgeTemplate(filepath string) bool {
	return strings.HasSuffix(filepath, TMPL_SUFFIX)
}

func ExcludeTemplate(filepath string, excludeDagger bool) bool {
	return excludeDagger && strings.Contains(filepath, DAGGER_FILE_TOKEN)
}

func CleanFilepath(filepath string, fileToken string) string {
	filepath = strings.TrimSuffix(filepath, ".tmpl")

	return strings.ReplaceAll(filepath, fileToken, "")
}
