package xlib

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/softlandia/xlib/internal/cp"
)

const (
	cp866r1Min  = 0x80 //заглавная буква А
	cp866r1Max  = 0xAF //строчная буква п - в этом интервале в 866 раскладке лежит большинство русских букв
	cp866r2Min  = 0xE0 //строчная р
	cp866r2Max  = 0xF1 //строчна ё - в этом интервале лежат остальные русские буквы
	cp1251s1    = 0xA8 //Ё
	cp1251s2    = 0xB8 //ё в этой позиции в 866 лежит псевдографика
	cp1251r1Min = 0xC0 //с этой позиции начинается весь алфавит
	cp1251r1Max = 0xFF //заканчивается
)

//CodePageDetect - detect code page of file
//return 0 if code page can not be detected
//return const xlib.CpWindows1251 for Windows code page 1251
//return const xlib.Cp866 for IBM 866 code page
//EF-BB-BF utf-8 bom
func CodePageDetect(fn string, stopStr ...string) (uint16, error) {
	var (
		r      rune
		cp1251 int
		cp866  int
	)

	iFile, err := os.Open(fn)
	if err != nil {
		return cp.ASCII, err
	}
	defer iFile.Close()

	iScanner := bufio.NewScanner(iFile)
	for i := 0; iScanner.Scan(); i++ {
		s := iScanner.Text()
		if len(stopStr) > 0 {
			if strings.Contains(s, stopStr[0]) { //stopStr[0] - строка, stopStr - слайс строк
				break
			}
		}
		for j := range s {
			r = rune(s[j])
			//проверка принадлежности символа позициям алфавитных символов в кодовой таблице 1251
			switch {
			case r == cp1251s1:
				cp1251++
			case r == cp1251s2:
				cp1251++
			case (r >= cp1251r1Min) && (r <= cp1251r1Max):
				cp1251++
			}
			//проверка принадлежности символа позициям алфавитных символов в кодовой таблице 866
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
		return cp.Windows1251, nil
	case cp1251 < cp866:
		return cp.IBM866, nil
	}
	return cp.ASCII, nil
}

//SeekFileToString - search string in text file and return *bufio.Scanner at founded line
//return number of line if string 'strToSearch' founded
//return scanner on line with string 'strToSearch'. first call scanner.Text() - return this line
//return (-1, nil, nil) if string 'strToSearch' not founded
//return (-1, nil, nil) if string 'strToSearch' is empty
//return (0, nil, err) if file not open or error occure when file reading
//opened file NOT CLOSED in any case!
func SeekFileToString(fileName, strToSearch string) (int, *bufio.Scanner, error) {
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
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
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
	i := 0 //index founded files
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
			//skip folders and files with extention not extFile
			return nil
		}
		//file found
		i++
		*fileList = append(*fileList, path)
		return nil
	})
	return i, err
}

//FileConvertCodePage - replace code page text file from one to another
func FileConvertCodePage(fileName string, fromCP, toCP uint16) error {
	if fromCP == toCP {
		return nil
	}

	iFile, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer iFile.Close()

	//TODO need using sytem tmp folder
	tmpFileName := fileName + "~"
	oFile, err := os.Create(tmpFileName)
	if err != nil {
		return err
	}
	defer oFile.Close()

	s := ""
	iScanner := bufio.NewScanner(iFile)
	for i := 0; iScanner.Scan(); i++ {
		s = iScanner.Text()
		s, err = StrConvertCodePage(s, fromCP, toCP)
		if err != nil {
			oFile.Close()
			os.Remove(tmpFileName)
			return fmt.Errorf("cde page convert error on file '%s': %v", fileName, err)
		}
		fmt.Fprintf(oFile, "%s\n", s)
	}
	oFile.Close()
	iFile.Close()
	return os.Rename(tmpFileName, fileName)
}
