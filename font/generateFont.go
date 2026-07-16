package font

import "fmt"

func GenerateFont() map[rune][]string {
	charMap := make(map[rune][]string)

	for r := rune(32); r <= rune(126); r++ {
		if r >= 'A' && r <= 'Z' {
			charMap[r] = []string{
				"##########",
				"          ",
				"          ",
				fmt.Sprintf("    %c    ", r),
				"          ",
				"          ",
				"          ",
				"##########",
			}
		} else if r >= 'a' && r <= 'z' {
			charMap[r] = []string{
				"##########",
				"          ",
				"          ",
				fmt.Sprintf("    %c    ", r),
				"          ",
				"          ",
				"          ",
				"##########",
			}
		} else if r >= '0' && r <= '9' {
			charMap[r] = []string{
				"##########",
				"          ",
				"          ",
				fmt.Sprintf("    %c    ", r),
				"          ",
				"          ",
				"          ",
				"##########",
			}
		} else {
			charMap[r] = []string{
				"##########",
				"          ",
				"          ",
				fmt.Sprintf("    %c    ", r),
				"          ",
				"          ",
				"          ",
				"##########",
			}
		}
	}
	return charMap
}
