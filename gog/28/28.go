package main

/*
#include <stdlib.h>
#cgo LDFLAGS: -L. -lmyLib

extern void myprint(int i);

void dofoo(void) {
	int i;
	for (i=0;i<10;i++) {
		myprint(i);
	}
}
*/
import "C"

func main() {
	s1 := C.myStruct{field: 1}
	s1Ret := C.fun(s1)
	_ = s1Ret
}
