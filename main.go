package main

import (
	"ascii-art/processor"
	"fmt"
	"os"
)

var usage = `
Usage: go run . [STRING] [BANNER]

Example 1: go run . "Hello\nThere"
Example 2: go run . "Hello\nThere" shadow
Example 3: go run . --output=<filename> "Hello\nThere"
Example 4: go run . --output=<filename> "Hello\nThere" shadow
Example 5: go run . --fs=<filename>
Example 6: go run . --fs=<filename> shadow

Avalaible Banners: standard, shadow, thinkertoy
`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		return
	}

	flags, remaining := parseArgs(os.Args[1:])
	var text string
	var banner string

	if flags.Fs != "" {
		content, err := os.ReadFile(flags.Fs)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot read file: %v\n", err)
			return
		}
		text = string(content)
		if len(remaining) > 0 { banner = remaining[0] } else { banner = "standard" }
	} else {
		if len(remaining) == 0 {
			fmt.Fprintln(os.Stderr, usage)
			return
		}
		text = remaining[0]
		if len(remaining) > 1 { banner = remaining[1] } else { banner = "standard" }
	}

	processor.ProcessInput(flags.Output, text, banner)
}
