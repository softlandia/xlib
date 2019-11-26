# golang util library #

(c) softlandia@gmail.com

>download: go get -u github.com/softlandia/xLib  
>install: go install

## dependences ##

>"golang.org/x/text/encoding/charmap"  
>"golang.org/x/text/transform"

## functions ##

1. FileExists(name string) bool
2. StrContainBackSlash(s string) bool
3. StrIsPrintRune(s string) bool
4. ChangeFileExt(iFileName, newExt string) string
5. SeekFileToString(fileName, strToSearch string) (*bufio.Scanner, error)
8. FindFilesExt(fileList *[]string, path, fileNameExt string) (int, error)

## description ##

    func FileExists(name string) bool  //return true if file exist

    func StrContainBackSlash(s string) bool  //return true if string s consist rune back slash '\'

    func StrIsPrintRune(s string) bool  //return true if input string consists only of printable rune

    func ChangeFileExt(iFileName, newExt string) string  //return file name with new extention

    func SeekFileToString(fileName, strToSearch string) (*bufio.Scanner, error)  //read text file fileName and return Scanner at line strToSearch

    func FindFilesExt(fileList *[]string, path, fileNameExt string) (int, error)  //search in path files with extention == fileNameExt and put file name to slice fileList


## tests ##

coverage 96.2%  
folder "test_files" contain files for testing, no remove/change/add
