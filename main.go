package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	filePaths := flag.String("f", "", "File(s) seperated by ,")
	flag.Parse()

	countResult := 0

	if *filePaths != "" {
		countResult = handleFiles(*filePaths, *lines, *bytes)
	} else {
		countResult = count(os.Stdin, *lines, *bytes)
	}

	fmt.Println(countResult)
}

func handleFiles(filePaths string, countLines, countBytes bool) int {
	paths := strings.Split(filePaths, ",")
	countResult := 0
	for _, fp := range paths {
		strippedPath := strings.ReplaceAll(fp, " ", "")
		file, err := os.Open(strippedPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		reader := bufio.NewReader(file)
		countResult += count(reader, countLines, countBytes)
	}

	return countResult
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
