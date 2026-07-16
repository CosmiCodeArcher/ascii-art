package render

import (
	"fmt"
	"io"
	"strings"
)

func Render(w io.Writer, charMap map[rune][]string, lines []string) {
	for _, line := range lines {
		if line == "" {
			fmt.Println()
			continue
		}

		for row := 0; row < 8; row++ {
			var rowBuilder strings.Builder
			for _, char := range line {
				if charRowSlice, ok := charMap[char]; ok && len(charRowSlice) > row {
					rowBuilder.WriteString(charRowSlice[row])
				}
			}
			fmt.Fprintln(w, rowBuilder.String())
		}
	}
}
