package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func inputPath(dir string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter a target root directory : ")
	dirs, _ := reader.ReadString('\n')
	dircs := strings.TrimSpace(dirs)
	return dircs
}

func main() {
	// Get files from current directory
	var input string
	dircs := inputPath(input)

	files, err := ioutil.ReadDir(dircs)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fname := f.Name()
		EXT := filepath.Ext(fname)

		if EXT != ".go" && EXT != "" { // Skip files with .go or "" extenions
			fmt.Println("Filename: " + fname)
			fmt.Printf("Extension found: %q\n", EXT)
			_ = os.MkdirAll(dircs+"./"+EXT, 0755)             // Mkdir and ignore errors if exists
			err := os.Rename(dircs+ "./"+fname, dircs+"/"+EXT+"./"+fname) // move file to directory
			if err != nil {
				log.Println("Could to move file:")
				log.Println(err)
			}
		}
	}
}
