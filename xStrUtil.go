package xLib

import (
	"path/filepath"
	"strings"
	"unicode"
)

//StrContainBackSlash - return true if input string contain '\'
func StrContainBackSlash(s string) bool {
	return strings.ContainsRune(s, 0x005C)
}

//StrIsPrintRune - return true if input string consists of printable rune
func StrIsPrintRune(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

//ChangeFileExt - change file name extention
func ChangeFileExt(iFileName, newExt string) string {
	return strings.TrimSuffix(iFileName, filepath.Ext(iFileName)) + newExt
}
