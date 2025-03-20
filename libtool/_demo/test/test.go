package main

import (
	"fmt"

	"github.com/goplus/llgo/c"
	"github.com/goplus/llpkg/libtool"
)

func main() {
	fmt.Println("Simple libtool demonstration")
	
	// Initialize libtool
	ret := libtool.LtDlinit()
	if ret != 0 {
		fmt.Println("Failed to initialize libtool:", c.GoString(libtool.LtDlerror()))
		return
	}
	fmt.Println("Successfully initialized libtool")
	
	// Try to load a common library (libc)
	libName := "libc.so.6" // Linux style
	handle := libtool.LtDlopen(c.Str(libName))
	if handle == nil {
		libName = "libc.dylib" // macOS style
		handle = libtool.LtDlopen(c.Str(libName))
	}
	if handle == nil {
		libName = "c" // Generic style
		handle = libtool.LtDlopen(c.Str(libName))
	}
	
	if handle != nil {
		fmt.Printf("Successfully opened %s\n", libName)
		
		// Try to find a common function (printf)
		symPtr := libtool.LtDlsym(handle, c.Str("printf"))
		if symPtr != nil {
			fmt.Println("Found 'printf' function")
		} else {
			fmt.Println("Symbol 'printf' not found:", c.GoString(libtool.LtDlerror()))
		}
		
		// Close the library
		libtool.LtDlclose(handle)
		fmt.Println("Closed library")
	} else {
		fmt.Println("Could not open any standard library:", c.GoString(libtool.LtDlerror()))
	}
	
	// Clean up libtool
	libtool.LtDlexit()
	fmt.Println("Successfully cleaned up libtool")
}