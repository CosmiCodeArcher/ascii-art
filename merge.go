package main

func MergeBanners(base map[rune][]string, priority map[rune][]string) map[rune][]string {
	result :=  make(map[rune][]string)

	for key, value := range base {
		result[key] = value
	}

	for key, value := range priority {
		result[key] = value
	}
	
	return result
}