package main

type A struct {
}

func (self *A) chain() *A {
	return self
}

func getAnA() (result *A) {
	result = &A{}
	return
}

func getAnotherA() *A {
	return &A{}
}

func main() {
	//This works and gets tab completed
	getAnotherA().chain()

	//So does this.
	t := getAnA()
	t.chain()

	//This is not tab completed and chain is marked as not found
	getAnA().chain()
}
