package converter

import "strings"

func StringToArt(input string) string {
	lines := strings.Split(strings.ReplaceAll(input, "\\n", "\n"), "\n")

	charMap := map[rune][]string {
		'0': {
			" ###### ",
			"#      #",
			"#      #",
			"#      #",
			" ###### ",
		},
		'1': {
			"##",
			" #",
			" #",
			" #",
			" #",
		},
		'2': {
			"#####",
			"   ##",
			"  #",
			"#",
			"#####",
		},
	}

	var finalString string

	for _, line := range lines {
		for row := 0; row < 5; row++ {
			var rowBuilder strings.Builder
			for _, char := range line {
				if charRowSlice, ok := charMap[char]; ok && len(charRowSlice) > row {
					rowBuilder.WriteString(charRowSlice[row])
				}
			}
			finalString += rowBuilder.String() + "\n"
		}
	}
	return finalString
}