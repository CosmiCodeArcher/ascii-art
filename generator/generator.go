package generator

func GeneratePattern(c rune) []string {
	if c < 'A' || c > 'Z' {
		return []string{}
	}

	patterns := map[rune][]string{
		'A': {
			"   ##   ",
			"  #  #  ",
			" #    # ",
			"#      #",
			"########",
			"#      #",
			"#      #",
			"#      #",
		},

		'B': {
			"##### ",
			"#    #",
			"#   # ",
			"#  #  ",
			"######",
			"#    #",
			"#    #",
			"######",
		},

		'C': {
			" #####",
			"#      #",
			"#       ",
			"#       ",
			"#       ",
			"#       ",
			"#      #",
			" #####",
		},
	}

	result, ok := patterns[c]
	if !ok {
		return []string{}
	}

	return result
}