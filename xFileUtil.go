package xLib

import (
	"os"
)

//StrContainBackSlash - return true if input string contain '\'
func FileExists(name string) bool {
    _, err := os.Stat(name)
    return !os.IsNotExist(err)
}