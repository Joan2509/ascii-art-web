package asciiart

import (
	"strings"
)

// ArtRetriever returns the ASCII art corresponding to the input string using the provided map.
func ArtRetriever(s string, m map[rune][]string) (string, error) {
	var result strings.Builder

	if s == "" {
		result.WriteString("")
		return result.String(), nil
	}
	lines := strings.Split(s, "\n")

	for ind := 0; ind < len(lines); ind++ {
		if lines[ind] == "" {
			result.WriteString("\n")
		} else {
			for j := 0; j < 8; j++ {
				for _, char := range lines[ind] {
					if asciiArt, ok := m[char]; ok {
						// Add the corresponding ASCII art for the character
						result.WriteString(asciiArt[j])
					}
				}
				result.WriteString("\n")
			}
		}
	}
	return result.String(), nil
}
