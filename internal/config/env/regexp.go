package env

import (
	"regexp"
	"strings"
)

var (
	matchFirstCap *regexp.Regexp = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCaps  *regexp.Regexp = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func toSnakeCase(text string) string {
	snake := matchFirstCap.ReplaceAllString(text, "${1}_${2}")
	snake = matchAllCaps.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func toUpperCase(text string) string {
	return strings.ToUpper(text)
}

func transformEnvText(text string) string {
	text = toSnakeCase(text)
	text = toUpperCase(text)
	return text
}
