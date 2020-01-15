package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	RemoveDuplicateFile()
}

// CRC32 crc32
func CRC32(str string) uint32 {
	return crc32.ChecksumIEEE(getFileContent(str))
}

// getFilelist get all file in the dir and the sub-dir
func getFilelist(path string) (pathString []string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		pathString = append(pathString, path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return
}

//GetSize get size of the file
func GetSize(fileName string) (size int64) {
	info, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("os.Stat err =", err)
		return
	}
	return info.Size()
}
func getFileContent(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("read file failed, err:", err)
	}
	return content
}

// RemoveDuplicateFile remove duplicate file in thd dir and its suddir
func RemoveDuplicateFile() {
	fileMap := make(map[int64][]string)
	var path string
	fmt.Println("请输入需要删除重复文件的目录如（d:/test/a）")
	fmt.Print(">>>")
	fmt.Scanln(&path)
	//获取该目录下的所有文件
	allfile := getFilelist(path)
	for _, file := range allfile {
		fileMap[GetSize(file)] = append(fileMap[GetSize(file)], file)
	}
	for _, v := range fileMap {
		if len(v) > 1 {
			uniqueMap := make(map[uint32]string)
			duplicateFileSlice := []string{}
			for _, value := range v {
				// fmt.Println(CRC32(value))
				crcValue := CRC32(value)
				_, ok := uniqueMap[crcValue]
				//if the map key is not exists, then add it to the unqiue map(finalmap)
				if ok {
					duplicateFileSlice = append(duplicateFileSlice, value)
				} else {
					uniqueMap[crcValue] = value
				}
			}
			for _, dfsV := range duplicateFileSlice {
				os.Remove(dfsV)
				fmt.Printf("删除文件%v\n", dfsV)
			}
		}
	}
}
