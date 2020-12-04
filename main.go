package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	filePath := flag.String("f", "", "File")
	flag.Parse()

	var reader io.Reader = os.Stdin

	if *filePath != "" {
		file, err := os.Open(*filePath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		reader = bufio.NewReader(file)
	}

	fmt.Println(count(reader, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)

	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	}

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}
