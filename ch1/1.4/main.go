package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lineCounts := make(map[string]int)
	lineFiles := make(map[string]map[string]bool)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, lineCounts, lineFiles, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, lineCounts, lineFiles, arg)
			f.Close()
		}
	}
	for line, n := range lineCounts {
		if n > 1 {
			filesMap := lineFiles[line]

			filenames := make([]string, 0, len(filesMap))
			for filename := range filesMap {
				filenames = append(filenames, filename)
			}

			fileList := strings.Join(filenames, ", ")

			fmt.Printf("The line %q appears a total of %d times.\nIt appears in the file(s): %s\n\n", line, n, fileList)
		}
	}
}

func countLines(f *os.File, lineCounts map[string]int, lineFiles map[string]map[string]bool, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()

		lineCounts[line]++

		if lineFiles[line] == nil {
			lineFiles[line] = make(map[string]bool)
		}

		lineFiles[line][filename] = true
	}
}
