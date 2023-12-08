package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, lines bool) int {

	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// if the -l flag is not set, split the input into words
	// otherwise split into lines
	if !lines {
		scanner.Split(bufio.ScanWords)
	}

	// Define a counter
	w := 0

	for scanner.Scan() {
		w += 1
	}
	return w
}
func main() {
	// Defining a boolean flag -l to count lines
	lines := flag.Bool("l", false, "Count lines")
	fmt.Println(count(os.Stdin, *lines))
}
