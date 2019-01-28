package main

import (
	"fmt"

	"github.com/softlandia/xlib"
)

func main() {
	scanner, err := xlib.SeekFileToString("test.txt", "~A")
	if err != nil {
		panic(err)
	}
	if scanner == nil {
		fmt.Println("text not found")
	} else {
		scanner.Scan()
		if scanner.Text() == "<OK>" {
			fmt.Printf("xlib.SeekFileToString... \tOK\n")
		} else {
			fmt.Printf("xlib.SeekFileToString... \tERROR\n")
		}

	}

	cp, err := xlib.CodePageDetect("test.txt", "~A")
	if err != nil {
		panic(err)
	}
	if cp != xlib.CpWindows1251 {
		fmt.Printf("xlib.CodePageDetect... \tOK\n")
	} else {
		fmt.Printf("xlib.CodePageDetect... \tERROR\n")
	}
}
