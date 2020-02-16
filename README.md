# golang util library #

(c) softlandia@gmail.com

>download: go get -u github.com/softlandia/xlib  
>install: go install

## functions ##

    func StrContainBackSlash(s string) bool
      return true if string s contain rune back slash '\'

    func StrIsPrintRune(s string) bool  
      return true if input string contain only of printable rune

    func ChangeFileExt(iFileName, newExt string) string  
      return file name with new extention

    func FileExists(name string) bool  
      return true if file exist

    func Max(x, y int) int
      return max from int

    func SortBytes(b []byte) []byte
      return sorted slice of byte

    func SeekFileStop(fileName, strToSearch string) (int, *bufio.Scanner, error)
      search string in text file and return *bufio.Scanner at founded line
      return number of line if string 'strToSearch' founded
      return scanner on line with string 'strToSearch'. first call scanner.Text() - return this line
      return (-1, nil, nil) if string 'strToSearch' not founded
      return (-1, nil, nil) if string 'strToSearch' is empty
      return (0, nil, err) if file not open or error occure when file reading
      successfull opened file NOT CLOSED in any case!
      strToSearch must contain only base ASCII rune

    func FindFilesExt(fileList *[]string, path, fileNameExt string) (int, error)  
      search all files in path with 'ext' & put to slice
      sample:  n, err := FindFilesExt(&fl, "c:\\tmp", ".log")

## tests ##

coverage 94%  
folder "test_files" contain files for testing, no remove/change/add
