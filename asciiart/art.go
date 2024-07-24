package asciiart

import (
	"fmt"
	"os"
)

// reads a banner file, creates a map of ASCII art, validates user input,
// and prints the corresponding ASCII art to the console.
func Art(inputText, banner string) (string, error) {
	bannerFile, err := ReadBannerFile("banners/" + banner + ".txt")
	check(err)
	ASCIIArtMap := MapCreator(string(bannerFile))
	asciiArt, err := ArtRetriever(inputText, ASCIIArtMap)
	check(err)
	return asciiArt, nil
}

func check(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", e)
		os.Exit(1)
	}
}
