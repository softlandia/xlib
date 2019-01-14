package xLib

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	//Windows1251 is the Windows 1251 code page
	Windows1251 = 1
	//CodePage866 is the IBM Code Page 866
	CodePage866 = 2
)

const (
	cp866r1Min  = 0x80
	cp866r1Max  = 0xAF
	cp866r2Min  = 0xE0
	cp866r2Max  = 0xF1
	cp1251s1    = 0xA8
	cp1251s2    = 0xB8
	cp1251r1Min = 0xC0
	cp1251r1Max = 0xFF
)

//CodePageDetect - detect code page of file
//return 0 if code page can not be detected
//return const Windows1251 for cp_1251
//return const CodePage866 for cp_866
func CodePageDetect(fn string) (int, error) {
	var (
		r      rune
		cp1251 int
		cp866  int
	)

	iFile, err := os.Open(fn)
	defer iFile.Close()
	if err != nil {
		return 0, err
	}

	iScanner := bufio.NewScanner(iFile)
	for i := 0; iScanner.Scan(); i++ {
		s := iScanner.Text()
		for j := range s {
			r = rune(s[j])
			switch {
			case r == cp1251s1:
				cp1251++
			case r == cp1251s2:
				cp1251++
			case (r >= cp1251r1Min) && (r <= cp1251r1Max):
				cp1251++
			}
			switch {
			case (r >= cp866r1Min) && (r <= cp866r1Max):
				cp866++
			case (r >= cp866r2Min) && (r <= cp866r2Max):
				cp866++
			}
		}
	}
	switch {
	case cp1251 > cp866:
		return Windows1251, nil
	case cp1251 < cp866:
		return CodePage866, nil
	}
	return 0, nil
}

//SeekFileToString - search string in text file and return *bufio.Scanner at founded line
//return scanner on line with string 'strToSearch'. first call scanner.Text() - return this line
//return (nil, nil) if string 'strToSearch' not found
//return (nil, err) if file not open or error occure when file reading
//opened file not close in any case!!!
func SeekFileToString(fileName, strToSearch string) (*bufio.Scanner, error) {
	iFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	p := -1
	iScanner := bufio.NewScanner(iFile)
	for i := 0; iScanner.Scan(); i++ {
		s := iScanner.Text()
		p = strings.Index(s, strToSearch)
		if p >= 0 {
			return iScanner, nil
		}
	}
	return nil, iScanner.Err()
}

//FileExists - return true if file exist
func FileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

//FindFilesExt - search all files in path with 'ext' & put to list
func FindFilesExt(fileList *[]string, path, fileNameExt string) (int, error) {
	if fileList == nil {
		return 0, errors.New("first parameter 'fileList' is nil")
	}
	i := 0 //index founded files
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if (info.IsDir()) && (filepath.Ext(path) != fileNameExt) {
			//skip folders and files with extention not fileNameExt
			return nil
		}
		//file found
		i++
		*fileList = append(*fileList, path)
		return nil
	})
	return i, err
}
