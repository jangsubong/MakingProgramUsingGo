package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Get files from current directory
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fname := f.Name()
		EXT := filepath.Ext(fname)
		if EXT != ".go" && EXT != "" { // Skip files with .go or "" extenions
			fmt.Println("Filename: " + fname)
			fmt.Printf("Extension found: %q\n", EXT)
			_ = os.Mkdir(EXT, 0755)                // Mkdir and ignore errors if exists
			err := os.Rename(fname, EXT+"/"+fname) // move file to directory
			if err != nil {
				log.Println("Could to move file:")
				log.Println(err)
			}
		}
	}
}
