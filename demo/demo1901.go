package main

import (
	"bufio"
	"io"
)

func main() {
	a := io.Reader{}
	reader := bufio.NewReader(a)
	_, _ = reader.ReadString('\n')
}
