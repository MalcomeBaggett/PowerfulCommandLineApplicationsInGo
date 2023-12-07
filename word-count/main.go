package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader) int {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// Define the scanner split type into words (default is splitting by lines)
	scanner.Split(bufio.ScanWords)

	// Define a counter
	w := 0

	for scanner.Scan() {
		w += 1
	}
	return w
}
func main() {
	fmt.Println(count(os.Stdin))
}
