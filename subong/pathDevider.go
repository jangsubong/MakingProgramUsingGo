package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(dir []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range dir {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func main() {
	lines, err := readLines("pathList.txt")

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var dirs []string
	var files []string

	for _, line := range lines {
		dir, file := filepath.Split(line)
		fmt.Printf("dir: %q\nfile: %q\n", dir, file)
		dirs = append(dirs, dir)
		files = append(files, file)

	}
	writeLines(dirs, "dir.txt")
	writeLines(files, "fileName.txt")

}
