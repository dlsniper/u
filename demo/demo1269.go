package main

import (
	"fmt"
)

func someFunc() error {
	return fmt.Errorf("someFunc")
}

func otherFunc() {

}

func someOtherFunc() (string, error) {
	return "demo", nil
}

func demoFunc(demo1 int) {
	demo1 = 2
	switch {
	case 1 == 2:
		demo1 := 1 // variable referenced below the switch statement
		_ = demo1
	default:
		_ = demo1
	}

	_ = demo1 // here the variable is referencing to the one in the case declaration
}

func main() {
	var (
		boolVar bool
		demo    string
		err     error // inspection error is here: Unused variable 'err'
	)

	switch {
	case boolVar:
		err := someFunc()
		if err != nil {
			otherFunc()
		}
	default:
		demo, err = someOtherFunc()
		if err != nil {
			otherFunc()
		}
	}

	_ = demo
}
