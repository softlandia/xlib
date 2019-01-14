### golang util library  ###

	download: go get -u github.com/softlandia/xLib
	install: go install

functions:  
1. FileExists(name string) bool
2. StrContainBackSlash(s string) bool
3. StrIsPrintRune(s string) bool
4. ChangeFileExt(iFileName, newExt string) string
5. SeekFileToString(fileName, strToSearch string) (*bufio.Scanner, error)
6. func CodePageDetect(fn string) (int, error)

identifyCodePage - sample of using function CodePageDetect


_________________________________________________________________________
	func FileExists(name string) bool  
return true if file exist

	func StrContainBackSlash(s string) bool
return true if last char in string s == '\'

	func StrIsPrintRune(s string) bool  
return true if input string consists only of printable rune

	func ChangeFileExt(iFileName, newExt string) string   
change file name extention

	func SeekFileToString(fileName, strToSearch string) (*bufio.Scanner, error)  
read text file fileName and return Scanner at line strToSearch

	func CodePageDetect(fn string) (int, error)  
read text file fn and return code page, detect only IBM CodePage866 and Windows1251