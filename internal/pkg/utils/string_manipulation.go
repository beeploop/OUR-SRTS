package utils

import "strings"

func StripUnderscore(text string) string {
	return strings.ReplaceAll(text, "_", " ")
}

func WhiteSpaceToUnderscore(text string) string {
	return strings.ReplaceAll(text, " ", "_")
}
