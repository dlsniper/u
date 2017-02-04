package demo1374

import "C"

func _() {
	i := C.int(1)
	_ = make([]*C.char, i)
}
