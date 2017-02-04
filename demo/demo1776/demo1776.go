package main

func noop() {}

func (a int) meth() {
	for {
		defer noop()
	}
}

func main() {
	for {
		func (){
			defer noop()
		}()
	}

	switch 1 {
	case 2: defer noop();
	}

	func (){
		defer noop()
	}()

	for {
		switch 1 {
		case 2: defer noop();
		}
	}

	defer noop();
	for {
		defer noop()
	}
}
