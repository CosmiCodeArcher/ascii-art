package main

import (
	"ascii-art/processor"
	"strings"
	"fmt"
	"os"
)

var usage = `
Usage: go run . [STRING] [BANNER]

Example 1: go run "Hello\nThere"
Example 2: go run "Hello\nThere" shadow
Example 3: go run --output=<filename> "Hello\nThere"
Example 4: go run --output=<filename> "Hello\nThere" shadow

Avalaible Banners: standard, shadow, thinkertoy
`

// main.go — detect --output= flag, extract filename, pass it through

func main() {
	switch len(os.Args) {
	case 1:
		fmt.Fprintf(os.Stderr, usage)
		return
	case 2:
		processor.ProcessInput("", os.Args[1], "standard")
		return
	case 3:
		if !strings.HasPrefix(os.Args[1], "--output=") {
			processor.ProcessInput("", os.Args[1], os.Args[2])
			return
		}
	}

	if len(os.Args) >= 3 && strings.HasPrefix(os.Args[1], "--output=") {
	    tagSlice := strings.SplitN(os.Args[1], "=", 2) // = [--output, <filename>]
		fileName := tagSlice[1]

		if len(os.Args) == 4 {
			processor.ProcessInput(fileName, os.Args[2], os.Args[3])
		} else if len(os.Args) == 3 {
			processor.ProcessInput(fileName, os.Args[2], "standard")
		} else {
			fmt.Fprintf(os.Stderr, usage)
		}
	}
}
