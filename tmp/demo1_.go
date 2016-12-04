package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := errors.New("Error generated here")
	if err != nil {
		fmt.Printf("ERROR: %+v\r\n", errors.Wrap(err, ""))
	}
}