package xlib

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//ReadFileStop - read text file to string, stop if found strStop
//return
//		"", nil if input strStop empty
//		"", error if can't open file
//		one string with all rune, strStop not include
//! end of line ignored
//TODO change += to string builder
func ReadFileStop(fileName, strStop string) (string, error) {
	res := ""
	if len(strStop) == 0 {
		return res, nil
	}

	iFile, err := os.Open(fileName)
	if err != nil {
		return res, err
	}
	defer iFile.Close()
	iScanner := bufio.NewScanner(iFile)
	var sb strings.Builder
	for i := 0; iScanner.Scan(); i++ {
		s := iScanner.Text()
		pos := strings.Index(s, strStop)
		if pos > 0 {
			sb.WriteString(s[:pos])
			return sb.String(), nil
		}
		sb.WriteString(s)
	}
	return sb.String(), iScanner.Err()
}

// TextScanner -
type TextScanner struct {
	scanner *bufio.Scanner
	file    *os.File
}

//SeekFileStop - search string in text file and return *bufio.Scanner at founded line
//return number of line if string 'strToSearch' founded
//return scanner on line with string 'strToSearch'. call scanner.Text() - return this line
//return (-1, nil, nil) if string 'strToSearch' not founded
//return (-1, nil, nil) if string 'strToSearch' is empty
//return (0, nil, err) if file not open or error occure when file reading
//successfull opened file NOT CLOSED in any case!
//strToSearch must contain only base ASCII rune
func SeekFileStop(fileName, strToSearch string) (int, *bufio.Scanner, error) {
	if len(strToSearch) == 0 {
		return -1, nil, nil
	}

	iFile, err := os.Open(fileName)
	if err != nil {
		return 0, nil, err
	}
	iScanner := bufio.NewScanner(iFile)
	for i := 0; iScanner.Scan(); i++ {
		s := iScanner.Text()
		if strings.Contains(s, strToSearch) {
			return i, iScanner, nil
		}
	}
	return -1, nil, iScanner.Err()
}

//FileExists - return true if file exist
func FileExists(name string) bool {
	inf, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return !inf.IsDir()
}

//FindFilesExt - search all files in path with 'ext' & put to list
//path - "c:\tmp"
//ext  - ".log"
//sample:  n, err := FindFilesExt(&fl, "c:\\tmp", ".log")
func FindFilesExt(fileList *[]string, path, fileNameExt string) (int, error) {
	if fileList == nil {
		return 0, errors.New("first parameter 'fileList' is nil")
	}
	extFile := strings.ToUpper(fileNameExt)
	index := 0 //index founded files
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			//skip folders
			return nil
		}
		if strings.ToUpper(filepath.Ext(path)) != extFile {
			//skip files with wrong extention
			return nil
		}
		index++
		*fileList = append(*fileList, path)
		return nil
	})
	return index, err
}
