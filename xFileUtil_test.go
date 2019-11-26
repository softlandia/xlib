//(c) softland 2019
//softlandia@gmail.com
package xlib

import (
	"testing"
)

//
func TestReadFileStop(t *testing.T) {
	s, err := ReadFileStop("", "")
	if (err != nil) || (s != "") {
		t.Errorf("<ReadFileStop> on strStop '' must return error==nil, but return %v and must return '' string, but return '%s'\n", err, s)
	}

	_, err = ReadFileStop("utils.go", "")
	if (err != nil) || (s != "") {
		t.Errorf("<ReadFileStop> on exist file 'utils.go' and empty strStop must return err = nil end empty string, but return %v, '%s'\n", err, s)
	}

	_, err = ReadFileStop("", "*")
	if err == nil {
		t.Errorf("<ReadFileStop> on file '-.-' must return error, but return nil\n")
	}

	s, err = ReadFileStop("test_files\\empty_file.txt", "*")
	if err != nil {
		t.Errorf("<ReadFileStop> on empty file must return error = nil, but return %v\n", err)
	}
	if len(s) > 0 {
		t.Errorf("<ReadFileStop> on empty file must return result = '', but return %s\n", s)
	}

	s, err = ReadFileStop("test_files\\rune_encode_error.txt", "4")
	if err != nil {
		t.Errorf("<ReadFileStop> on file 'rune_encode_error.txt' must return error = nil, but return %v\n", err)
	}
	if s != "123" {
		t.Errorf("<ReadFileStop> on file 'rune_encode_error.txt' must return result = '123', but return %s\n", s)
	}

	s, err = ReadFileStop("test_files\\2line.txt", " ")
	if err != nil {
		t.Errorf("<ReadFileStop> on file 'rune_encode_error.txt' must return error = nil, but return %v\n", err)
	}
	if s != "1234" {
		/*oFile, _ := os.Create("res.dat")
		oFile.WriteString(s)
		oFile.Close()*/
		t.Errorf("<ReadFileStop> on file 'rune_encode_error.txt' must return result = '1234', but return %s\n", s)
	}
}

//SeekFileToString
func TestSeekFileStop(t *testing.T) {
	_, _, err := SeekFileStop("-.-", "-")
	if err == nil {
		t.Errorf("<SeekFileStop> on file '-.-' must return error, but return nil")
	}

	index, scanner, err := SeekFileStop("test_files\\866&1251.txt", "~A")
	if scanner == nil {
		t.Errorf("<SeekFileStop> on file '866&1251.txt' return scanner == nil")
	}
	scanner.Scan()
	if scanner.Text() != "<OK>" {
		t.Errorf("<SeekFileStop> on line: %d must be string '<OK>'\n", index)
	}

	index, scanner, err = SeekFileStop("test_files\\866&1251.txt", "")
	if index >= 0 {
		t.Errorf("<SeekFileStop> on empty seek str == '' must return index < 0 [-1], return: %d", index)
	}

	index, scanner, err = SeekFileStop("test_files\\empty_file.txt", "~")
	if index >= 0 {
		t.Errorf("<SeekFileStop> on file 'empty_file.txt' and seek str == '~' must return index < 0 [-1], return: %d", index)
	}

	index, scanner, err = SeekFileStop("test_files\\rune_error_1251.txt", "#")
	if (index != 1) || (err != nil) {
		t.Errorf("<SeekFileStop> on file 'rune_error_1251.txt' and seek str == '#' must return index == 1, return: %d, %v", index, err)
	}
}

//FileExists
func TestFileExists(t *testing.T) {
	if FileExists("-.-") {
		t.Errorf("<FileExists> return true on non exist file '-.-'")
	}
	if !FileExists("test_files\\866&1251.txt") || !FileExists("test_files\\2line.txt") || !FileExists("test_files\\empty_file.txt") {
		t.Error("<FileExists> return false on exist files: 866&1251.txt, 866to1251.txt, empty_file.txt")
	}
}

//FindFilesExt
func TestFindFilesExt(t *testing.T) {
	_, err := FindFilesExt(nil, ".", "txt")
	if err == nil {
		t.Errorf("<FindFilesExt> on nil input fileList nust return err, return: %v", err)
	}

	fl := make([]string, 0, 10)
	n, err := FindFilesExt(&fl, ".", ".txt")
	if n != 5 {
		t.Errorf("<FindFilesExt> on current folder must found 6 '.txt', return: %d", n)
	}
}
