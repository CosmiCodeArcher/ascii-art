package processor

import (
	"os"
	"fmt"
	"strings"
	"ascii-art/parser"
	"ascii-art/banner"
	"ascii-art/render"
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

	render.Render(LoadedBanner, ParsedInput)
}