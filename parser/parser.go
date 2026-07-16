package parser

import "strings"

func ParseInput(text string) []string {
	return strings.Split(strings.ReplaceAll(text, "\\n", "\n"), "\n")
}