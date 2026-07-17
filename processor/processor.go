package processor

import (
	"ascii-art/banner"
	"ascii-art/parser"
	"ascii-art/render"
	"fmt"
	"os"
	"strings"
)

var bannerNameErr = "Avalaible Banners: standard, shadow, thinkertoy"

func ProcessInput(outputFile, text, bannerName string) {
	var bannerFilePath string
	bannerName = strings.ToLower(bannerName)

	switch bannerName {
	case "standard", "shadow", "thinkertoy":
		bannerFilePath = "banners/" + bannerName + ".txt"
	default:
		fmt.Fprintln(os.Stderr, bannerNameErr)
		return
	}

	LoadedBanner, err := banner.LoadBanner(bannerFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Not Found")
		return
	}

	ParsedInput := parser.ParseInput(text)

	if outputFile != "" {

		openOutputFile, err := os.Create(outputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		defer openOutputFile.Close()

		render.Render(openOutputFile, LoadedBanner, ParsedInput)

	} else {
		render.Render(os.Stdout, LoadedBanner, ParsedInput)
	}
}