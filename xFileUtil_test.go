//(c) softland 2019
//softlandia@gmail.com
package xlib

import (
	"os"
	"testing"
)

//CodePageDetect
func TestCodePageDetect(t *testing.T) {
	res, err := CodePageDetect("test_files\\test.txt", "~X~")
	if err != nil {
		t.Errorf("<CodePageDetect> on file '%s' return error: %v", "test.txt", err)
	}
	if res != Cp866 {
		t.Errorf("<CodePageDetect> on file '%s' expected 866 got: %s", "test.txt", CodePageAsString(res))
	}

	res, err = CodePageDetect("test_files\\test.txt")
	if res != CpWindows1251 {
		t.Errorf("<CodePageDetect> on file '%s' expected 1251 got: %s", "test.txt", CodePageAsString(res))
	}

	_, err = CodePageDetect("-.-")
	if err == nil {
		t.Errorf("<CodePageDetect> on file '-.-' must return error, but return nil")
	}

	res, _ = CodePageDetect("test_files\\test2.txt")
	if res != CpEmpty {
		t.Errorf("<CodePageDetect> on file 'test2.txt' expect CpEmpty got: %s", CodePageAsString(res))
	}

	res, err = CodePageDetect("test_files\\test3.txt")
	if (res != CpEmpty) || (err != nil) {
		t.Errorf("<CodePageDetect> on file 'test3.txt' expect CpEmpty and no error got: %s and %v", CodePageAsString(res), err)
	}
}

//SeekFileToString
func TestSeekFileToString(t *testing.T) {
	_, _, err := SeekFileToString("-.-", "-")
	if err == nil {
		t.Errorf("<CodePageDetect> on file '-.-' must return error, but return nil")
	}

	index, scanner, err := SeekFileToString("test_files\\test.txt", "~A")
	if scanner == nil {
		t.Errorf("<CodePageDetect> on file 'test.txt' return scanner == nil")
	}
	scanner.Scan()
	if scanner.Text() != "<OK>" {
		t.Errorf("<SeekFileToString> on line: %d must be string '<OK>'\n", index)
	}

	index, scanner, err = SeekFileToString("test_files\\test.txt", "")
	if index >= 0 {
		t.Errorf("<SeekFileToString> on empty seek str == '' must return index < 0 [-1], return: %d", index)
	}

	index, scanner, err = SeekFileToString("test_files\\test3.txt", "~")
	if index >= 0 {
		t.Errorf("<SeekFileToString> on file 'test3.txt' and seek str == '~' must return index < 0 [-1], return: %d", index)
	}
}

//FileConvertCodePage
func TestFileConvertCodePage(t *testing.T) {
	err := FileConvertCodePage("", 0, 1)
	if err == nil {
		t.Errorf("<FileConvertCodePage> on empty file name expected error, got: %v", err)
	}

	err = FileConvertCodePage("", 0, 0)
	if err != nil {
		t.Errorf("<FileConvertCodePage> on fromCp == toCp expected error==nil, got: %v", err)
	}

	err = FileConvertCodePage("test_files\\test4.txt", Cp866, CpWindows1251)
	if err == nil {
		t.Errorf("<FileConvertCodePage> expected error, got: %v", err)
	}

	os.Link("test_files\\test5.txt", "test_files\\test5.tmp.txt")
	err = FileConvertCodePage("test_files\\test5.tmp.txt", Cp866, CpWindows1251)
	if err != nil {
		t.Errorf("<FileConvertCodePage> expect no err, got: %v", err)
	}
	os.Remove("test_files\\test5.tmp.txt")
}

//FileExists
func TestFileExists(t *testing.T) {
	if FileExists("-.-") {
		t.Errorf("<FileExists> return true on non exist file '-.-'")
	}
	if !FileExists("test_files\\test.txt") || !FileExists("test_files\\test2.txt") || !FileExists("test_files\\test3.txt") {
		t.Error("<FileExists> return false on exist files: test.txt, test2.txt, test3.txt")
	}
}

//FindFilesExt
func TestFindFilesExt(t *testing.T) {
	//fl := make([]string, 0, 10)
	//fl := nil
	_, err := FindFilesExt(nil, ".", "txt")
	if err == nil {
		t.Errorf("<FindFilesExt> on nil input fileList nust return err, return: %v", err)
	}

	fl := make([]string, 0, 10)
	n, err := FindFilesExt(&fl, ".", ".txt")
	if n != 5 {
		t.Errorf("<FindFilesExt> on current folder must found 3 '.txt', return: %d", n)
	}
}
