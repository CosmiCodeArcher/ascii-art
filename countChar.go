package main

func CountChars(charMap map[rune][]string, input string) map[rune]int {
	charCount := make(map[rune]int)

	for _, c := range input {
		if _, ok := charMap[c]; !ok {
			continue
		}
		charCount[c]++
	}

	return charCount
}
