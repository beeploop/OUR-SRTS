package utils

import (
	"slices"
	"strings"
)

func StripUnderscore(text string) string {
	return strings.ReplaceAll(text, "_", " ")
}

func WhiteSpaceToUnderscore(text string) string {
	return strings.ReplaceAll(text, " ", "_")
}

func Capitalize(text string) string {
	words := strings.Split(text, " ")

	capitalized := slices.AppendSeq(
		make([]string, 0),
		Map(words, func(word string) string {
			first := word[0:1]
			return strings.Join([]string{strings.ToUpper(first), word[1:]}, "")
		}),
	)

	return strings.Join(capitalized, " ")
}

func StripQueryParams(url string) string {
	return strings.Split(url, "?")[0]
}
