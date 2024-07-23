package asciiart

import (
	"fmt"
	"os"
)

// reads a banner file, creates a map of ASCII art, validates user input,
// and prints the corresponding ASCII art to the console.
func Art() string {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		return ""
	}
	banner := "standard"
	inputText := os.Args[1]
	if len(os.Args) == 3 {
		banner = os.Args[2]
	}

	bannerFile, err := ReadBannerFile("banners/" + banner + ".txt")
	check(err)

	ASCIIArtMap := MapCreator(string(bannerFile))

	asciiArt, err := ArtRetriever(inputText, ASCIIArtMap)
	check(err)

	return asciiArt
}
func check(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", e)
		os.Exit(1)
	}
}
