xFileUtil_test.go:37:2: ineffectual assignment to `index` (ineffassign)
	index, scanner, err := SeekFileStop("test_files\\866&1251.txt", "~A")
	^
xFileUtil_test.go:42:18: ineffectual assignment to `err` (ineffassign)
	index, scanner, err = SeekFileStop("test_files\\866&1251.txt", "")
	                ^
xFileUtil_test.go:45:18: ineffectual assignment to `err` (ineffassign)
	index, scanner, err = SeekFileStop("test_files\\empty_file.txt", "~")
	                ^
xFileUtil_test.go:69:5: ineffectual assignment to `err` (ineffassign)
	n, err := FindFilesExt(&fl, ".", ".txt")
	   ^
xFileUtil_test.go:48:9: SA4006: this value of `scanner` is never used (staticcheck)
	index, scanner, err = SeekFileStop("test_files\\rune_error_1251.txt", "#")
	       ^
