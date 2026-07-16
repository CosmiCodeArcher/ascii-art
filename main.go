package main

import (
	"os"
	"fmt"
	"strings"
	"ascii-art/parser"
	"ascii-art/banner"
	"ascii-art/render"
)

var usage = `
Usage: go run . [STRING] [BANNER]

Example 1: go run "Hello\nThere"
Example 2: go run "Hello\nThere" shadow

Avalaible Banners: standard, shadow, thinkertoy
`

var bannerNameErr = "Avalaible Banners: standard, shadow, thinkertoy"

func main() {
	switch len(os.Args) {
	case 1:
		fmt.Fprintf(os.Stderr, usage)
		return
	case 2:
		processInput(os.Args[1], "standard")
	case 3:
		processInput(os.Args[1], os.Args[2])
	}
}

func processInput(text, bannerName string) {
	var bannerFilePath string
	bannerName = strings.ToLower(bannerName)

	switch bannerName {
	case "standard", "shadow", "thinkertoy":
		bannerFilePath = "banners/" + bannerName + ".txt"
	default:
		fmt.Fprintf(os.Stderr, bannerNameErr)
		return
	}
	
	LoadedBanner, err := banner.LoadBanner(bannerFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Not Found")
		return
	}

	ParsedInput := parser.ParseInput(text)

	render.Render(LoadedBanner, ParsedInput)
}