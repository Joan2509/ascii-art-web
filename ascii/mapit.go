package ascii

import (
	"bufio"
	"fmt"
	"os"
)

/*The maps function, opens the required file,
scans line by line and maps the line number to the strings found in the various
lines in the opened file. For all these,
each character line contents are concatenated into a single string
the new string is returned as a line.
*/

func Maps(n int) string {
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	var count int
	var line, str string
	for scanner.Scan() {
		line = scanner.Text()
		count++
		ascMap := make(map[int]string)
		if count == n {
			for _, ch := range line {
				str += string(ch)
			}
		}
		ascMap[n] = str
	}
	defer file.Close()
	return str
}
