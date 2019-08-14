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
//newExt must start from '.' sample '.xyz'
func ChangeFileExt(iFileName, newExt string) string {
	return strings.TrimSuffix(iFileName, filepath.Ext(iFileName)) + newExt
}

//StrConvertCodePage - convert string from one code page to another
func StrConvertCodePage(s string, fromCP, toCP int64) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	if fromCP == toCP {
		return s, nil
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

//CodePageAsString - return string is name of char set with id cp
func CodePageAsString(cp int) string {
	switch cp {
	case Cp866:
		return Cp866AsStr
	case CpWindows1251:
		return CpWindows1251AsStr
	case CpUtf8:
		return CpUtf8AsStr
	default:
		return CpEmptyAsStr
	}
}

//ContainsOtherRune - return true if sting s contains any other rune not in runes
//on empty parameters - return false, -1
func ContainsOtherRune(s string, runes ...rune) (bool, int) {
	var (
		i int
		r rune
	)
	if (len(s) == 0) || (len(runes) == 0) {
		return false, -1
	}
	for i, r = range s {
		res := true
		for _, sr := range runes {
			res = (res && (r != sr))
		}
		if res {
			return res, i
		}
	}
	return false, 0
}

//StrCopyStop - return s, stop on rune in stopRune
func StrCopyStop(s string, stopRune ...rune) (string, int) {
	var (
		i int
		r rune
	)
	if len(stopRune) > 0 {
		for i, r = range s {
			for _, sr := range stopRune {
				if r == sr {
					return s[:i], i
				}
			}
		}
	}
	return s, len(s)
}

//ReplaceAllSpace - return string with one space
func ReplaceAllSpace(s string) string {
	for strings.Index(s, "  ") >= 0 {
		s = strings.ReplaceAll(s, "  ", " ")
	}
	return s
}

//ReplaceSeparators - return string with one separator rune
// ' .' >> '.' // '. ' >> '.' // ' :' >> ':' // ': ' >> ':'
func ReplaceSeparators(s string) string {
	type TSeparatorsReplacement struct {
		old string
		new string
	}
	var SeparatorsList = []TSeparatorsReplacement{
		{" .", "."},
		{". ", "."},
		{" :", ":"},
		{": ", ":"},
	}
	for _, sep := range SeparatorsList {
		s = strings.ReplaceAll(s, sep.old, sep.new)
	}
	return s
}
