package main

import "flag"

func main() {
	boolArg := flag.Bool("boolArg", false, "usage")

	flag.Parse()

	if boolArg != false {  // compilation error non-bool (type *bool) used as if condition
		//...
	}
}
