package issue

type dem struct{}

func _() {
	a := Demo{
		Field: 1,  // Refactor here
	}
	_ = a
}
