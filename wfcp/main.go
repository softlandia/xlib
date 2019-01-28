package main

import (
	"fmt"
	"os"
	"time"

	"github.com/softlandia/xlib"
)

func main() {

	fmt.Println("<full file scan>")
	t0 := time.Now()
	cp, err := xlib.CodePageDetect(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	switch cp {
	case xlib.CpWindows1251:
		fmt.Println("found 1251 code page")
	case xlib.Cp866:
		fmt.Println("found 866 code page")
	default:
		fmt.Println("code page not found")
	}
	fmt.Printf("elapsed: %v\n", time.Since(t0))

	fmt.Println("<file scan with stop criteria>")
	t0 = time.Now()
	cp, err = xlib.CodePageDetect(os.Args[1], "~A")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	switch cp {
	case xlib.CpWindows1251:
		fmt.Println("found 1251 code page")
	case xlib.Cp866:
		fmt.Println("found 866 code page")
	default:
		fmt.Println("code page not found")
	}
	fmt.Printf("elapsed: %v\n", time.Since(t0))

}
