// +build darwin linux

//THIS PACKAGE WAS TAKEN DIRECTLY FROM "https://github.com/antoine-richard/gomobile-text" AND WAS NOT MODIFIED IN ANY WAY

package main

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	"golang.org/x/mobile/asset"
	mfont "golang.org/x/mobile/exp/font"
)

func LoadCustomFont() (font *truetype.Font, err error) {
	file, err := asset.Open("System San Francisco Display Regular.ttf")
	if err != nil {
		fmt.Printf("error opening font asset: %v\n", err)
		return loadFallbackFont()
	}
	defer file.Close()
	raw, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error reading font: %v\n", err)
		return loadFallbackFont()
	}
	font, err = truetype.Parse(raw)
	if err != nil {
		fmt.Printf("error parsing font: %v\n", err)
		return loadFallbackFont()
	}
	return font, nil
}

func loadFallbackFont() (font *truetype.Font, err error) {
	// Default font doesn't work on Darwin
	fmt.Println("using Monospace font")
	return truetype.Parse(mfont.Monospace())
}
