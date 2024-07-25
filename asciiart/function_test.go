package asciiart

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestReadBannerFile(t *testing.T) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "test_banner.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	// Write some content to the temporary file
	content := "Test banner content"
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	// Close the temporary file
	if err := tmpFile.Close(); err != nil {
		t.Fatalf("Failed to close temporary file: %v", err)
	}
	// Test ReadBannerFile function
	bannerContent, err := ReadBannerFile(tmpFile.Name())
	if err != nil {
		t.Errorf("ReadBannerFile returned error: %v", err)
	}
	if bannerContent != content {
		t.Errorf("ReadBannerFile did not read expected content. Got: %s, Expected: %s", bannerContent, content)
	}
}

func TestMapCreator(t *testing.T) {
	// Test input string
	input := `
         
         
         
         
         
         
         
         
`
	// Expected output map
	expectedMap := map[rune][]string{
		' ': {
			"         ",
			"         ",
			"         ",
			"         ",
			"         ",
			"         ",
			"         ",
			"         ",
		},
	}
	// Call MapCreator function
	result := MapCreator(input)
	// Compare result with expected output
	if !reflect.DeepEqual(result, expectedMap) {
		t.Errorf("MapCreator returned unexpected result.\nExpected: %v\nGot: %v", expectedMap, result)
	}
}

func TestArtRetriever(t *testing.T) {
	// Define test cases
	tests := []struct {
		input    string
		artMap   map[rune][]string
		expected string
		err      error
	}{
		{
			input: "HE",
			artMap: map[rune][]string{
				'H': {
					"  _    _  ",
					" | |  | | ",
					" | |__| | ",
					" |  __  | ",
					" | |  | | ",
					" |_|  |_| ",
					"          ",
					"          ",
				},
				'E': {
					" ______  ",
					"|  ____| ",
					"| |__    ",
					"|  __|   ",
					"| |____  ",
					"|______| ",
					"         ",
					"         ",
				},
			},
			expected: "  _    _   ______  \n | |  | | |  ____| \n | |__| | | |__    \n |  __  | |  __|   \n | |  | | | |____  \n |_|  |_| |______| \n                   \n                   \n",
			err:      nil,
		},
		{
			input:    "£",
			artMap:   map[rune][]string{},
			expected: "400",
			err:      fmt.Errorf("error! invalid input: %s", string('£')),
		},
	}
	// Iterate over test cases
	for _, test := range tests {
		output, err := ArtRetriever(test.input, test.artMap)
		// Check for error
		if err != nil && err == test.err {
			t.Errorf("Unexpected error for input: %s - %v", test.input, test.err)
			os.Exit(1)
		}
		// Check for expected output
		if output != test.expected {
			t.Errorf("For input: %s, expected:\n%s\nbut got:\n%s\n", test.input, test.expected, output)
		}
	}
}
