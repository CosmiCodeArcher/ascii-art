package main

import (
	"fmt"
	"errors"
)

func ValidateBanner(bannerMap map[rune][]string) error {
	if bannerMap == nil {
		return errors.New("banner is nil")
	}

	if len(bannerMap) != 95 {
		return fmt.Errorf("banner has %d entries, expected 95", len(bannerMap))
	}

	for char := 32; char <= 126; char++ {
		if _, ok := bannerMap[rune(char)]; !ok {
			return fmt.Errorf("missing character: '%c' (ASCII %d)", rune(char), char)
		}

		if len(bannerMap[rune(char)]) != 8 {
			return fmt.Errorf("character '%c' has %d lines, expected 8", rune(char), len(bannerMap[rune(char)]))
		}
	}

	return nil
}