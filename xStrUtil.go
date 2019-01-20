package xlib

import (
	"path/filepath"
	"strings"
	"unicode"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
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

//ChangeFileExt - change in path string file name extention
func ChangeFileExt(iFileName, newExt string) string {
	return strings.TrimSuffix(iFileName, filepath.Ext(iFileName)) + newExt
}

//ConvertStrCodePage - convert string from one code page to another
func ConvertStrCodePage(s string, fromCP, toCP int64) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	var err error

	switch fromCP {
	case Cp866:
		s, _, err = transform.String(charmap.CodePage866.NewDecoder(), s)
	case CpWindows1251:
		s, _, err = transform.String(charmap.Windows1251.NewDecoder(), s)
	}
	switch toCP {
	case Cp866:
		s, _, err = transform.String(charmap.CodePage866.NewEncoder(), s)
	case CpWindows1251:
		s, _, err = transform.String(charmap.Windows1251.NewEncoder(), s)
	}
	return s, err
}
