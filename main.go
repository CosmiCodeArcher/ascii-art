package main

import (
	"os"
	"fmt"
	"ascii-art/processor"
)

var usage = `
Usage: go run . [STRING] [BANNER]

Example 1: go run "Hello\nThere"
Example 2: go run "Hello\nThere" shadow

Avalaible Banners: standard, shadow, thinkertoy
`

func main() {
	switch len(os.Args) {
	case 1:
		fmt.Fprintf(os.Stderr, usage)
		return
	case 2:
		processor.ProcessInput(os.Args[1], "standard")
	case 3:
		processor.ProcessInput(os.Args[1], os.Args[2])
	}
}

