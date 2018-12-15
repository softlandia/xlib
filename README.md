### golang util library  ###

	download: go get -u github.com/softlandia/xLib

functions:  
1. FileExists(name string) bool
2. StrContainBackSlash(s string) bool
3. StrIsPrintRune(s string) bool
4. ChangeFileExt(iFileName, newExt string) string

_________________________________________________________________________
	func FileExists(name string) bool

return true if file exist

	func StrContainBackSlash(s string) bool

return true if last char in string s == '\'

	func StrIsPrintRune(s string) bool

return true if input string consists of printable rune

	func ChangeFileExt(iFileName, newExt string) string {
change file name extention