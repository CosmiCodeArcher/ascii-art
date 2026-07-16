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

func ProcessInput(text, bannerName string) {
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

	render.Render(os.Stdout, LoadedBanner, ParsedInput)
}
