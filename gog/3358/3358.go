package main

/*
#include <stdlib.h>
#cgo LDFLAGS: -L. -lmyLib
*/
import "C"

func main() {
	s1 := C.myStruct{field: 1} // highlighted as invalid
	s1Ret := C.fun(s1)         // correct
	_ = s1Ret
}
