package main

import (
	"fmt"

	"github.com/dlsniper/u/demo/pk/pk2"
)

func (s *myType) Hello() {

}

func main() {
	e := pk2.Emo()
	fmt.Printf("%t\n", e[0] == pk2.Er)
}
