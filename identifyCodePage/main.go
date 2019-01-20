package main

import (
	"fmt"
	"os"

	"github.com/softlandia/xlib"
)

func main() {

	cp, err := xlib.CodePageDetect(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	switch cp {
	case xlib.CpWindows1251:
		fmt.Printf("found 1251 code page")
		os.Exit(xlib.CpWindows1251)
	case xlib.Cp866:
		fmt.Printf("found 866 code page")
		os.Exit(xlib.Cp866)
	}
	fmt.Printf("code page not found")
	os.Exit(0)
}
