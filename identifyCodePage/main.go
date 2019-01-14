package main

import (
	"fmt"
	"os"

	"github.com/softlandia/xLib"
)

func main() {

	cp, err := xLib.CodePageDetect(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	switch cp {
	case xLib.Windows1251:
		fmt.Printf("found 1251 code page")
		os.Exit(xLib.Windows1251)
	case xLib.CodePage866:
		fmt.Printf("found 866 code page")
		os.Exit(xLib.CodePage866)
	}
	fmt.Printf("code page not found")
	os.Exit(0)
}
