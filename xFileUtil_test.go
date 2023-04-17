//(c) softland 2019-2023
//softlandia@gmail.com
package xlib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//
func TestReadFileStop(t *testing.T) {
	s, err := ReadFileStop("", "")
	assert.Nil(t, err)
	assert.Empty(t, s)
	_, err = ReadFileStop("utils.go", "")
	assert.Nil(t, err)
	_, err = ReadFileStop("", "*")
	assert.NotNil(t, err)
	s, err = ReadFileStop("test_files/empty_file.txt", "*")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(s))
	s, err = ReadFileStop("test_files/rune_encode_error.txt", "4")
	assert.Nil(t, err)
	assert.Equal(t, "123", s)
	s, err = ReadFileStop("test_files/2line.txt", " ")
	assert.Nil(t, err)
	assert.Equal(t, "1234", s)
}

//SeekFileToString
func TestSeekFileStop(t *testing.T) {
	_, _, err := SeekFileStop("-.-", "-")
	assert.NotNil(t, err)

	index, scanner, _ := SeekFileStop("test_files/866&1251.txt", "~A")
	assert.NotNil(t, scanner)
	scanner.Scan()
	assert.Equal(t, "<OK>", scanner.Text())

	index, _, _ = SeekFileStop("test_files/866&1251.txt", "")
	assert.Less(t, index, 0, fmt.Sprintf("1: %d", index))

	index, _, _ = SeekFileStop("test_files/empty_file.txt", "~")
	assert.Less(t, index, 0, fmt.Sprintf("2: %d", index))

	index, _, err = SeekFileStop("test_files/rune_error_1251.txt", "#")
	assert.Nil(t, err)
	assert.Equal(t, 1, index)
}

//FileExists
func TestFileExists(t *testing.T) {
	if FileExists("test_files") {
		t.Errorf("<FileExist> return true on folder")
	}
	if FileExists("-.-") {
		t.Errorf("<FileExists> return true on non exist file '-.-'")
	}
	if !FileExists("test_files/866&1251.txt") || !FileExists("test_files/2line.txt") || !FileExists("test_files/empty_file.txt") {
		t.Error("<FileExists> return false on exist files: 866&1251.txt, 866to1251.txt, empty_file.txt")
	}
}

//FindFilesExt
func TestFindFilesExt(t *testing.T) {
	_, err := FindFilesExt(nil, ".", "txt")
	assert.NotNil(t, err)

	fl := make([]string, 0, 10)
	n, _ := FindFilesExt(&fl, ".", ".txt")
	assert.Equal(t, n, 5)
}
