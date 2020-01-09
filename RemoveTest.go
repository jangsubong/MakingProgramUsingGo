package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

//대상 디렉터리 경로를 입력받음
func inputPath(dir string) string {
	fmt.Println("경로 입력:")
	fmt.Scan(&dir)
	return dir
}

//대상 디렉터리 하위 포함한 정보 출력
func dirReadString(dirpath string) ([]string, []int64) {
	var fileName []string
	var fileInfo []int64
	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			// 파일 명(경로포함) 및 파일 크기 출력
			fileName = append(fileName, path)
			fileInfo = append(fileInfo, info.Size())
		}
		return nil
	})
	if err != nil {
		log.Println(err)
		//	fmt.Println(list)
		//return files, err
		//	return
	}
	return fileName, fileInfo
}

//func export_csv() {}

func removeExtension(f []string) { //Plz, Check this
	//파일 목록이 저장된 슬라이스 입력 및 i에 리스트 저장
	for _, i := range f {
		file, err := os.Stat(i)
		if err != nil {
			panic(err)
		}
		fileName := file.Name()
		fmt.Println("test" + fileName)
		if filepath.Ext(fileName) == ".png" || filepath.Ext(fileName) == ".PNG" {
			// os.Remove(fileName) <-- you are removing the file in the current dir
			os.Remove(i) // I woud not use "i" as the var
			fmt.Println("Deleted " + file.Name())
		}

		/*
			if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".png" {
				os.Remove(file.Name())
				fmt.Println("Deleted", file.Name())
				}
			}*/
	}
}

//func remove_duplicated() {}

func main() {
	var input string
	dirPath := inputPath(input)
	filePath, fileInfo := dirReadString(dirPath) // Do not use _ in go
	removeExtension(filePath)
	//	fmt.Print(file_path[:], file_info[:])
	//	fileinfo, _ := dir_read_string(dirPath)
	fmt.Println(fileInfo)
}
