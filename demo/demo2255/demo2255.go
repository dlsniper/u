// demo
package commentstart

type (
// help
	Hello struct{}

// help
	Helloi interface{}
)

// help
type Helllo struct{}

// help
type Hellloi interface{}

const (
// help
	Helloc = 3
)

// help
const Hellloc, hellloc2 = 1, 2

var (
// help
	Hello1, Hellow1 int
)

// help
var Hello2 int

// Helllo
func (a Helllo) Hellllo() {
	_ = Hello1
	_ = Hellow1
	_ = Hello2
	_ = Helloc
	_ = Hellloc
	_ = hellloc2
}

// Demo does things  -> correct
func Demo() {
	Demo2()
}

// Demo does other things ->  incorrect
func Demo2() {
	Demo()
	Demo3()
	Demo4()
	Demo5()
	Demo6()
	Demo7()
	Demo8()
	Demo9()
}

// Demo3
func Demo3() {}

// A Demo4 does things  -> correct
func Demo4() {}

// An Demo5 does things  -> correct
func Demo5() {}

// The Demo6 does things  -> correct
func Demo6() {}

// Demo7 does things  -> correct
//
// Deprecated: use other thing
func Demo7() {}

// Demo8 does things  -> incorrect
//
// Deprecated: use other thing

func Demo8() {}

// Demo does things  -> correct
//
// Deprecated: use other thing

// Demo9 demo
func Demo9() {}

var (
	A int
// demo comment
	C int
)

const (
	DemoDemo = 0

	DemoEmo = 1
)

// This should be correct
const (
	NewConst  = 1
	NewConst2 = 1
)

type (
	Deeeeo struct {
	}
)

// digi
func Gigel() {
	// demo
	B := "Hello"
	_, _, _, _, _ = DemoDemo, DemoEmo, A, B, C
}

func Gigel2() {
	_, _ = NewConst, NewConst2
}

func (d Deeeeo) Gigel() {
	Gigel()
	Gigel2()
}
