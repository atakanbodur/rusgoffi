package main

// NOTE: There should be NO space between the comments and the `import "C"` line.

/*
#cgo LDFLAGS: ./lib/libhello.so -ldl
#include "./lib/hello.h"
*/
import "C"
import "log"

func main() {
	C.hello(C.CString("world"))
	C.whisper(C.CString("this is code from the dynamic library"))
	var str = C.GoString(C.testReturn(C.CString("atakan")))
	log.Println(str)
}
