package main

import (
	"bufio"
	"fmt"
	"godirwalk"
	"hash/fnv"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"
)

type filehash struct {
	path  string
	hash  uint64
	mtime time.Time
	size  int64
	err   error
}

const scanAll = 0
const ScanLength = 4096

//대상 디렉터리 경로를 입력받음
func input_path(dir string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter a target root directory : ")
	dirs, _ := reader.ReadString('\n')
	dircs := strings.TrimSpace(dirs)
	return dircs
}

// Scan input path directorys
func scanDir(root string) (int, int, [][]filehash, []string) {
	var (
		fileCount     = 0
		dupCount      = 0
		sameSizeCount = 0
		fileByeSize   = make(map[int64][]string)
		fileTime      []time.Time
		file_list     []string
	)

	err := godirwalk.Walk(root, &godirwalk.Options{
		Callback: func(path string, ph *godirwalk.Dirent) error {
			if ph.IsDir() {
				return nil
			}
			file_list = append(file_list, path)
			info, err := os.Stat(path)
			if err == nil {
				fileCount += 1
				size := info.Size()
				fileTime = append(fileTime, info.ModTime())
				if size > 0 {
					files, ok := fileByeSize[size]
					if !ok {
						files = make([]string, 0, 2)
					} else {
						sameSizeCount += 1
					}
					fileByeSize[size] = append(files, path)
				}
			} else {
				fmt.Println(err)
			}
			return nil
		},
		ErrorCallback: func(osPathname string, err error) godirwalk.ErrorAction {
			return godirwalk.SkipNode
		},
		Unsorted: true,
	})
	if err != nil {
		panic(err)
	}

	//Find Samme files : Files size
	samesizeFiles := make([][]filehash, 0, sameSizeCount)
	for size, files := range fileByeSize {
		if len(files) > 1 {
			fh := make([]filehash, len(files))
			for i := 0; i < len(files); i++ {
				fh[i] = filehash{path: files[i], size: size, mtime: fileTime[i]}

			}

			samesizeFiles = append(samesizeFiles, fh) //get filelist
		}

	}
	return fileCount, dupCount, samesizeFiles, file_list
}

// Find Dupliecate files
func getDuplicates(potentialDups [][]filehash, scanLength int64) [][]filehash {
	runtime.GOMAXPROCS(runtime.NumCPU())
	maxFds := runtime.NumCPU()
	throttle := make(chan bool, maxFds)
	for _, files := range potentialDups {
		for idx := 0; idx < len(files); idx++ {
			if scanLength != scanAll || files[idx].size > ScanLength {
				throttle <- true
				go func(p *filehash) {
					getFileChecksum(p, scanLength) //Get files Checksum (Calculating sum64 Hashs)
					<-throttle
				}(&files[idx])
			}
		}
	}
	for i := 0; i < maxFds; i++ {
		throttle <- true
	}
	duplicates := make([][]filehash, 0, len(potentialDups))
	for _, files := range potentialDups {
		hashToFiles := make(map[uint64][]filehash)
		for _, file := range files {
			if file.err == nil {
				files, ok := hashToFiles[file.hash]
				if !ok {
					files = make([]filehash, 0, 2)
				}
				hashToFiles[file.hash] = append(files, file)
			}
		}
		for _, files := range hashToFiles {
			if len(files) > 1 {
				duplicates = append(duplicates, files) // Add files hash info
			}
		}
	}
	return duplicates
}

func removeDuplicates(duplicates [][]filehash) (dupCount int) {
	for _, files := range duplicates {

		//If there are too many duplicate files, they cannot be sorted properly. :(
		sort.Slice(files, func(i int, j int) bool {
			return files[i].mtime.Before(files[j].mtime)
		})
		fmt.Println("Original is", files[0].path)
		for _, path := range files[1:] {
			dupCount += 1
			fmt.Println("Deleting duplicate file ", path.path)
			_ = os.Remove(path.path)
		}
	}
	return dupCount
}

func getFileChecksum(file *filehash, scanSize int64) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	list, err := os.Open(file.path)
	if err != nil {
		file.err = err
		return
	}
	defer list.Close()
	file_hash := fnv.New64a()
	if scanSize != scanAll {
		buf := make([]byte, scanSize)
		fmt.Printf("Scanning doubtful file... %s\n", file.path)
		n, err := list.Read(buf)
		if err == nil {
			file_hash.Write(buf[:n])
			file.hash = file_hash.Sum64()
		} else {
			file.err = err
		}
	} else {
		fmt.Printf("Calculating hash values for file %s...\n", file.path)
		_, file.err = io.Copy(file_hash, list)
		if file.err == nil {
			file.hash = file_hash.Sum64()
		}
	}

}

func RemoveByExt(filelist []string, fileCnt int) (int, int) {
	var extCount = 0
	for _, files := range filelist {
		file, err := os.Stat(files)
		if err != nil {
			panic(err)
		}
		if file.Size() == 0 {
			_ = os.Remove(files)
		}
		fileName := file.Name()
		fmt.Println("Scanning file extension... ", fileName)
		if filepath.Ext(fileName) != ".doc" && filepath.Ext(fileName) != ".DOC" && filepath.Ext(fileName) != ".ppt" && filepath.Ext(fileName) != ".PPT" && filepath.Ext(fileName) != ".xls" && filepath.Ext(fileName) != ".XLS" && filepath.Ext(fileName) != ".xlsx" && filepath.Ext(fileName) != ".XLSX" && filepath.Ext(fileName) != ".xlsb" && filepath.Ext(fileName) != ".XLSB" && filepath.Ext(fileName) != ".hwp" && filepath.Ext(fileName) != ".HWP" && filepath.Ext(fileName) != ".rtf" && filepath.Ext(fileName) != ".RTF" && filepath.Ext(fileName) != ".txt" && filepath.Ext(fileName) != ".TXT" && filepath.Ext(fileName) != ".pdf" && filepath.Ext(fileName) != ".PDF" && filepath.Ext(fileName) != ".csv" && filepath.Ext(fileName) != ".CSV" && filepath.Ext(fileName) != ".eml" && filepath.Ext(fileName) != ".EML" && filepath.Ext(fileName) != ".pst" && filepath.Ext(fileName) != ".PST" && filepath.Ext(fileName) != ".xlsm" && filepath.Ext(fileName) != ".XLSM" && filepath.Ext(fileName) != ".mbox" && filepath.Ext(fileName) != ".MBOX" && filepath.Ext(fileName) != ".ost" && filepath.Ext(fileName) != ".OST" && filepath.Ext(fileName) != ".msg" && filepath.Ext(fileName) != ".MSG" && filepath.Ext(fileName) != ".dbx" && filepath.Ext(fileName) != ".DBX" && filepath.Ext(fileName) != ".emlx" && filepath.Ext(fileName) != ".EMLX" && filepath.Ext(fileName) != ".docx" && filepath.Ext(fileName) != ".DOCX" && filepath.Ext(fileName) != ".pptx" && filepath.Ext(fileName) != ".PPTX" {
			os.Remove(files)
			fmt.Println("Deleting file by extension... ", files)
			extCount += 1
		}
	}
	fmt.Println("\n2nd routine loding...........\n")
	return fileCnt, extCount
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var (
		input         string
		DupCnt        int
		samesizeFiles [][]filehash
		test          string
	)
	startTime := time.Now()
	dirPath := input_path(input)
	st, err := os.Stat(dirPath)
	if err != nil {
		panic(err)
	}
	if !st.IsDir() {
		fmt.Println("Invaild path", dirPath)
	}
	fmt.Println("\nTasks start : ", startTime)
	fmt.Println("\n1st routine loding...........\n")
	fileCnt, _, _, filelst := scanDir(dirPath)
	fileCnt_ext, extCount := RemoveByExt(filelst, fileCnt)
	//Reset
	fileCnt = 0
	//Re allocate value
	fileCnt, DupCnt, samesizeFiles, _ = scanDir(dirPath)
	if len(samesizeFiles) > 0 {
		potentialDups := getDuplicates(samesizeFiles, ScanLength)
		if len(potentialDups) > 0 {
			duplicates := getDuplicates(potentialDups, scanAll)
			if len(duplicates) > 0 {
				DupCnt = removeDuplicates(duplicates)
			}
		}
	}
	elapsedTime := time.Since(startTime)
	fmt.Println("\nFinish : ", time.Now())
	fmt.Println("\nResults---------------------------------------------------------")
	fmt.Printf("\n Program execution time : %s \n", elapsedTime)
	fmt.Printf("\n Total %d file deleting files by extension : %d files   \n", fileCnt_ext, extCount)
	fmt.Printf("\n After %v files, %d duplicates\n", fileCnt, DupCnt)
	fmt.Printf("\n Residual : %d files remain (Not include directorys)\n", fileCnt-DupCnt)
	fmt.Println("\n----------------------------------------------------------------\n")
	fmt.Scanln(&test)
}

//Writer : Myeongjin.Goh
