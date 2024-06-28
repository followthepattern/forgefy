package naming

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CapitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func LowerFirst(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

func ToLowerCamelCase(input string) string {
	re := regexp.MustCompile(`(?:_| |\b)(\w)`)
	camelCase := re.ReplaceAllStringFunc(input, func(s string) string {
		return strings.ToUpper(s[len(s)-1:])
	})
	camelCase = strings.ReplaceAll(camelCase, "_", "")
	camelCase = strings.ReplaceAll(camelCase, " ", "")
	if len(camelCase) > 0 {
		r := []rune(camelCase)
		r[0] = unicode.ToLower(r[0])
		camelCase = string(r)
	}
	return camelCase
}

func ToSnakeCase(input string) string {
	// Regular expression to match uppercase letters and digits
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	snake := re.ReplaceAllString(input, `${1}_${2}`)
	snake = strings.ToLower(snake)
	snake = strings.ReplaceAll(snake, " ", "_")
	snake = strings.ReplaceAll(snake, "-", "_")
	return snake
}

func ToHumanReadable(input string) string {
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	humanReadable := re.ReplaceAllString(input, `${1} ${2}`)

	humanReadable = strings.ReplaceAll(humanReadable, "_", " ")
	humanReadable = strings.ReplaceAll(humanReadable, "-", " ")
	caser := cases.Title(language.English)
	humanReadable = caser.String(humanReadable)

	runes := []rune(humanReadable)
	if len(runes) > 0 {
		runes[0] = unicode.ToUpper(runes[0])
	}

	return string(runes)
}
