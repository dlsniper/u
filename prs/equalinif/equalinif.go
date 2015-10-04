package gigel

func main() {
	a := 1
	b := 2

	if a = b {
		println("a = b")
	} else {
		println("a != b")
	}

	if a := b {
		println("a = b")
	} else {
		println("a != b")
	}

	if a = b; a != 1 {
		println("a = b")
	} else { 
		println("a != b")
	}

	if a := b; a == 1 {
		println("a = b")
	} else {
		println("a != b")
	}
}
