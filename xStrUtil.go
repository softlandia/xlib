package xlib

import (
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
	"time"
)

// Secret - возвращает строку с вырезанной серединой, удобно для отображения токенов и паролей
func Secret(s string) string {
	if len(s) < 4 {
		return s
	}
	return s[:2] + "..." + s[len(s)-2:]
}

// AtoI - convert string to int, if error occure, return def value
func AtoI(s string, def int) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return def
	}
	return int(n)
}

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

//ContainsOtherRune - if sting s contains any other rune not in runes then return true and position of first this rune
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
	for strings.Contains(s, "  ") {
		s = strings.ReplaceAll(s, "  ", " ")
	}
	return s
}

//ReplaceSeparators - return string with one separator rune
// ' .' >> '.'
// '. ' >> '.'
// ' :' >> ':'
// ': ' >> ':'
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
		//{":.", ":"},
		//{".:", "."},
	}
	for _, sep := range SeparatorsList {
		s = strings.ReplaceAll(s, sep.old, sep.new)
	}
	return s
}

// ParseBool - при ошибке возвращает значение по умолчанию
func ParseBool(b string, def bool) bool {
	res, err := strconv.ParseBool(b)
	if err != nil {
		res = def
	}
	return res
}

// ParseDate - если не выйдет, вернёт заданную по умолчанию дату
func ParseDate(dt string, def time.Time) time.Time {
	dtProcessed, err := time.Parse("2006-01-02", dt)
	if err != nil {
		return def
	}
	return dtProcessed
}

