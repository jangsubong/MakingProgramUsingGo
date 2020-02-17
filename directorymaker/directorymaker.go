package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func inputPath(dir string) string {
	reader := bufio.NewReader(os.Stdin)

	dirs, _ := reader.ReadString('\n')
	dircs := strings.TrimSpace(dirs)
	return dircs
}

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
	fmt.Printf("Hello, This Program is directory maker \n")

	var input string
	fmt.Printf("Please enter a target directory : \n")
	targetDircs := inputPath(input)
	//fmt.Printf(targetDircs)

	fmt.Printf("Please enter a Path fileList : ")
  fmt.Printf("To use korean, you have to input unicode text file.)
	fileList := inputPath(input)
	//fmt.Printf(fileList)

	lines, err := readLines(fileList)

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	var dirs []string
	var files []string
	var redir string

	//var ext []string

	for _, line := range lines {
		dir, file := filepath.Split(line)
		//fmt.Printf("dir: %q\nfile: %q\n", dir, file)
		dirs = append(dirs, dir)
		files = append(files, file)
		//ext = append(ext, path.Ext(file))
		redir = strings.Replace(dir, ":", "_", -1)

		_ = os.MkdirAll(targetDircs+"/"+redir, 0755)

		fmt.Printf(redir + "\n")

	}
	_ = os.Mkdir("log", 0755)
	writeLines(dirs, "log/dir.txt")
	writeLines(files, "log/fileName.txt")

}
